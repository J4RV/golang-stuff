package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah/game"
)

var states = make(map[string]game.State)

func getPlayerIndex(req *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(req)["playerid"])
	if err != nil {
		return -1, err
	}
	return id, nil
}

func getPlayer(req *http.Request) (game.Player, error) {
	id, err := getPlayerIndex(req)
	if err != nil {
		return game.Player{}, err
	}
	s := getState(req)
	if id < 0 || id >= len(s.Players) {
		return game.Player{}, errors.New("Non valid player index in request")
	}
	return *s.Players[id], nil
}

func getState(req *http.Request) game.State {
	id := mux.Vars(req)["gameid"]
	return states[id]
}

func updateState(req *http.Request, s game.State) {
	id := mux.Vars(req)["gameid"]
	states[id] = s
}
