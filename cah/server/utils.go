package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah"
)

// TODO most things here should be moved
func playerIndex(g cah.GameState, u cah.User) (int, error) {
	for i, p := range g.Players {
		if p.User.ID == u.ID {
			return i, nil
		}
	}
	return -1, errors.New("You are not playing this game")
}

func player(g cah.GameState, u cah.User) (*cah.Player, error) {
	i, err := playerIndex(g, u)
	if err != nil {
		return &cah.Player{}, errors.New("You are not playing this game")
	}
	return g.Players[i], nil
}

func gameFromRequest(req *http.Request) (cah.GameState, error) {
	id := mux.Vars(req)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return cah.GameState{}, err
	}
	g, err := usecase.GameState.ByID(idInt)
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
	s := usecase.GameState.NewGame(
		usecase.GameState.Options().BlackDeck(bGameCards),
		usecase.GameState.Options().WhiteDeck(wGameCards),
		usecase.GameState.Options().HandSize(15),
	)
	usecase.Game.Create(users[0], "Test", "", []string{"base-uk", "expansion-1", "expansion-2", "kikis"}, s)
	/* Leave it open to check the game list
	p := getPlayersForUsers(users...)
	s, err := usecase.GameState.Start(p, s, usecase.GameState.Options().RandomStartingCzar())
	if err != nil {
		panic(err)
	}*/
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
