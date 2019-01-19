package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah"
)

var games = make(map[string]cah.Game)

func getPlayerIndex(g cah.Game, u cah.User) (int, error) {
	for i, p := range g.Players {
		if p.User.ID == u.ID {
			return i, nil
		}
	}
	return -1, errors.New("You are not playing this game")
}

func getPlayer(g cah.Game, u cah.User) (*cah.Player, error) {
	i, err := getPlayerIndex(g, u)
	if err != nil {
		return &cah.Player{}, errors.New("You are not playing this game")
	}
	return g.Players[i], nil
}

func getWhiteCardsInPlay(g cah.Game) []cah.WhiteCard {
	ret := []cah.WhiteCard{}
	for _, p := range g.Players {
		ret = append(ret, p.WhiteCardsInPlay...)
	}
	return ret
}

func getGame(req *http.Request) (cah.Game, error) {
	id := mux.Vars(req)["gameid"]
	g, ok := games[id]
	if !ok {
		return g, errors.New("Cannot get game from request")
	}
	return g, nil
}

func updateGameState(req *http.Request, s cah.Game) error {
	id := mux.Vars(req)["gameid"]
	_, ok := games[id]
	if !ok {
		return fmt.Errorf("No game found with id '%s'", id)
	}
	games[id] = s
	return nil
}

func createGame(users []cah.User, gameid string) error {
	_, ok := games[gameid]
	if ok {
		return fmt.Errorf("There already exists a game with id '%s'", gameid)
	}
	bd := usecase.Card.AllBlacks()
	wd := usecase.Card.AllWhites()
	var wGameCards = make([]cah.WhiteCard, len(wd))
	for i, c := range wd {
		wGameCards[i] = c
	}
	var bGameCards = make([]cah.BlackCard, len(bd))
	for i, c := range bd {
		bGameCards[i] = c
	}
	p := getPlayersForUsers(users...)
	s := usecase.Game.NewGame(
		usecase.Game.Options().BlackDeck(bGameCards),
		usecase.Game.Options().WhiteDeck(wGameCards),
		usecase.Game.Options().HandSize(15),
	)
	s, err := usecase.Game.Start(p, s, usecase.Game.Options().RandomStartingCzar())
	if err != nil {
		panic(err)
	}
	games[gameid] = s
	return nil
}

// TODO move to data: CreatePlayersFromUsers
func getPlayersForUsers(users ...cah.User) []*cah.Player {
	ret := make([]*cah.Player, len(users))
	for i, u := range users {
		ret[i] = cah.NewPlayer(u)
	}
	return ret
}

// For quick prototyping

func createTestGame() {
	createGame(getTestUsers(), "test")
}

func getTestUsers() []cah.User {
	users := make([]cah.User, 4)
	for i := 0; i < 4; i++ {
		u, _ := usecase.User.ByID(i)
		users[i] = u
	}
	return users
}
