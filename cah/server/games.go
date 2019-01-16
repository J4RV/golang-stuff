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
	userToPlayers map[data.User]*game.Player
}

func getPlayerIndex(g serverGame, u data.User) (int, error) {
	player, ok := g.userToPlayers[u]
	if ok {
		for i, p := range g.state.Players {
			if p.ID == player.ID {
				return i, nil
			}
		}
	}
	return -1, errors.New("You are not playing this game")
}

func getPlayer(g serverGame, u data.User) (*game.Player, error) {
	player, ok := g.userToPlayers[u]
	if !ok {
		return player, errors.New("Could not find player in game")
	}
	return player, nil
}

func getWhiteCardsInPlay(g serverGame) []game.WhiteCard {
	ret := []game.WhiteCard{}
	for _, p := range g.state.Players {
		ret = append(ret, p.WhiteCardsInPlay...)
	}
	return ret
}

func getGame(req *http.Request) (serverGame, error) {
	id := mux.Vars(req)["gameid"]
	g, ok := games[id]
	if !ok {
		return g, errors.New("Cannot get game from request")
	}
	return g, nil
}

func updateGameState(req *http.Request, s game.State) error {
	id := mux.Vars(req)["gameid"]
	g, ok := games[id]
	if !ok {
		return errors.New("Cannot update game state from request")
	}
	g.state = s
	return nil
}
