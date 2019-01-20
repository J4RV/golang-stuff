package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah"
)

func playerIndex(g cah.Game, u cah.User) (int, error) {
	for i, p := range g.Players {
		if p.User.ID == u.ID {
			return i, nil
		}
	}
	return -1, errors.New("You are not playing this game")
}

func player(g cah.Game, u cah.User) (*cah.Player, error) {
	i, err := playerIndex(g, u)
	if err != nil {
		return &cah.Player{}, errors.New("You are not playing this game")
	}
	return g.Players[i], nil
}

func gameFromRequest(req *http.Request) (cah.Game, error) {
	id := mux.Vars(req)["gameID"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return cah.Game{}, err
	}
	g, err := usecase.Game.ByID(idInt)
	if err != nil {
		return g, errors.New("Could not get game from request")
	}
	return g, nil
}

func createGame(users []cah.User) error {
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
	createGame(getTestUsers())
}

func getTestUsers() []cah.User {
	users := make([]cah.User, 4)
	for i := 0; i < 4; i++ {
		u, _ := usecase.User.ByID(i + 1)
		users[i] = u
	}
	return users
}
