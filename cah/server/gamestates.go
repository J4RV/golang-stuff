package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah/game"
)

var states = make(map[string]game.State)

func getState(req *http.Request) game.State {
	id := mux.Vars(req)["stateid"]
	return states[id]
}

func updateState(req *http.Request, s game.State) {
	id := mux.Vars(req)["stateid"]
	states[id] = s
}
