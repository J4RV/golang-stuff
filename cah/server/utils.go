package server

import (
	"errors"

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
	users := getTestUsers()
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
	usecase.Game.Create(users[1], "A long and descriptive game name", "")
	usecase.Game.Create(users[0], "Amo a juga", "1234")
	/* Leave it open to check the game list
	p := getPlayersForUsers(users...)
	s, err := usecase.GameState.Start(p, s,
		usecase.GameState.Options().RandomStartingCzar(),
		usecase.GameState.Options().BlackDeck(bGameCards),
		usecase.GameState.Options().WhiteDeck(wGameCards),
		usecase.GameState.Options().HandSize(15),
	)
	if err != nil {
		panic(err)
	}*/
}

func getTestUsers() []cah.User {
	users := make([]cah.User, 4)
	for i := 0; i < 4; i++ {
		u, _ := usecase.User.ByID(i + 1)
		users[i] = u
	}
	return users
}
