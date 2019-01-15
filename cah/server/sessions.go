package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/j4rv/golang-stuff/cah/data"
)

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
