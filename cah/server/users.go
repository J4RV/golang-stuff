package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	skey := os.Getenv("SESSION_KEY")
	if skey == "" {
		panic("Please set SESSION_KEY environment variable; it is needed to have secure cookies")
	}
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}

/*
func loginPassword(n, p string) (string, error) {
	u, err := data.GetUser(n, p)
	if err != nil {
		return "", err
	}
	newSid, err := generateRandomString(32)
	if err != nil {
		return "", err
	}
	sessionIds[newSid] = u
	return newSid, nil
}

func userFromSession(sid string) (data.User, error) {
	u, ok := sessionIds[sid]
	if !ok {
		return data.User{}, fmt.Errorf("Incorrect session id: '%s'", sid)
	}
	return u, nil
}

func processLogin(w http.ResponseWriter, req *http.Request) error {
	var payload loginPayload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	sid, err := loginPassword(payload.Name, payload.Pass)
	if err != nil {
		return err
	}
	addSessionCookie(&w, sid)

}
*/
const sessionid = "session_token"

/*
func requiresSessionCookie(appHandler) {
	c, err := store.Get(sessionid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return c.Value,
}*/

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, sessionid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Session", session.Values)
	// Set some session values.
	session.Values["nick"] = "Anon"
	session.Values["userid"] = 0
	// Save it before we write to the response/return from the handler.
	log.Println("Session", session.Values)
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// session generation stuff

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
