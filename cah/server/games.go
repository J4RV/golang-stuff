package main

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/game"
)

var games = make(map[string]serverGame)

type serverGame struct {
	state         game.State
	userToPlayers map[*data.User]*game.Player
}

func getPlayerIndex(g serverGame, u *data.User) (int, error) {
	player, ok := g.userToPlayers[u]
	if ok {
		for i, p := range g.state.Players {
			if p == player {
				return i, nil
			}
		}
	}
	return -1, errors.New("You are not playing this game")
}

func getPlayer(g serverGame, u *data.User) (game.Player, error) {
	player, ok := g.userToPlayers[u]
	if !ok {
		return *player, errors.New("Could not find player in game")
	}
	return *player, nil
}

func getGame(req *http.Request) serverGame {
	id := mux.Vars(req)["gameid"]
	return games[id]
}

func updateGame(req *http.Request, sg serverGame) {
	id := mux.Vars(req)["gameid"]
	games[id] = sg
}
