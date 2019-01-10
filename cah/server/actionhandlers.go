package main

import (
	"encoding/json"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/game"
)

func simpleCAHActionHandler(f func(game.State) (game.State, error)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		s := getState(req)
		newS, err := f(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
		} else {
			updateState(req, newS)
			writeJSONState(w, newS)
		}
	}
}

func giveBlackCardToWinner(w http.ResponseWriter, req *http.Request) {
	var payload chooseWinnerPayload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	newS, err := game.GiveBlackCardToWinner(payload.Winner, getState(req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	updateState(req, newS)
	writeJSONState(w, newS)
}
