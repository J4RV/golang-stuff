package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/j4rv/golang-stuff/cah/data"
)

func processLogin(w http.ResponseWriter, req *http.Request) {
	var payload loginPayload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Misconstructed payload", http.StatusBadRequest)
		return
	}
	u, err := data.GetUserByLogin(payload.Username, payload.Password)
	if err != nil {
		log.Printf("Someone tried to login using user '%s'", u.Username)
		http.Error(w, "Incorrect login", http.StatusForbidden)
		return
	}
	session, err := store.Get(req, sessionid)
	session.Values[userid] = u.ID
	session.Save(req, w)
	log.Printf("User %s just logged in!", u.Username)
	// everything ok, back to index with your brand new session!
	http.Redirect(w, req, "/", http.StatusOK)
}

// wraps an appHandler to make sure that a valid user is logged in
type loggedHandler func(http.ResponseWriter, *http.Request, data.User) error

func (fn loggedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := userFromSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	err = fn(w, r, u)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
	SESSIONS STUFF
*/

var store *sessions.CookieStore

const sessionid = "session_token"
const userid = "user_id"

func init() {
	skey := os.Getenv("SESSION_KEY")
	if skey == "" {
		panic("Please set SESSION_KEY environment variable; it is needed to have secure cookies")
	}
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}

func userFromSession(r *http.Request) (data.User, error) {
	session, err := store.Get(r, sessionid)
	if err != nil {
		return data.User{}, err
	}
	val, ok := session.Values[userid]
	if !ok {
		return data.User{}, fmt.Errorf("Tried to get user from session without an id: '%s'", session)
	}
	id, ok := val.(int)
	if !ok {
		return data.User{}, fmt.Errorf("Session with non int id value: '%s'", session)
	}
	u, err := data.GetUserById(id)
	if err != nil {
		return u, err
	}
	return u, nil
}
