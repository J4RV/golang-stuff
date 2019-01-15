package main

import (
	"encoding/json"
	"log"
	"net/http"

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
func loggedIn(ah appHandler) appHandler {
	u, err := userFromSession
	if err != nil {
		return simpleErrorHandler(err, http.StatusForbidden)
	}
	return ah
}
