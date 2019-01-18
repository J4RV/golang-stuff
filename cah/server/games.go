package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/game"
	"github.com/j4rv/golang-stuff/cah/model"
)

var games = make(map[string]serverGame)
var cardRepo data.CardRepository

func init() {
	cardRepo = data.NewCardStore()
	data.CreateCardsFromFolder(cardRepo, "./expansions/base-uk", "Base-UK")
	data.CreateCardsFromFolder(cardRepo, "./expansions/anime", "Anime")
	data.CreateCardsFromFolder(cardRepo, "./expansions/kikis", "Kikis")
	data.CreateCardsFromFolder(cardRepo, "./expansions/expansion-1", "The First Expansion")
	data.CreateCardsFromFolder(cardRepo, "./expansions/expansion-2", "The Second Expansion")
}

type serverGame struct {
	state         model.State
	userToPlayers map[model.User]*model.Player
}

func getPlayerIndex(g serverGame, u model.User) (int, error) {
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

func getPlayer(g serverGame, u model.User) (*model.Player, error) {
	player, ok := g.userToPlayers[u]
	if !ok {
		return player, errors.New("Could not find player in game")
	}
	return player, nil
}

func getWhiteCardsInPlay(g serverGame) []model.WhiteCard {
	ret := []model.WhiteCard{}
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

func updateGameState(req *http.Request, s model.State) error {
	id := mux.Vars(req)["gameid"]
	g, ok := games[id]
	if !ok {
		return errors.New("Cannot update game state from request")
	}
	g.state = s
	games[id] = g
	return nil
}

func createGame(users []model.User, gameid string) error {
	_, ok := games[gameid]
	if ok {
		return fmt.Errorf("There already exists a game with id '%s'", gameid)
	}
	bd := cardRepo.GetBlacks()
	wd := cardRepo.GetWhites()
	var wGameCards = make([]model.WhiteCard, len(wd))
	for i, c := range wd {
		wGameCards[i] = c
	}
	var bGameCards = make([]model.BlackCard, len(bd))
	for i, c := range bd {
		bGameCards[i] = c
	}
	p := getPlayersForUsers(users...)
	s := game.NewGame(p,
		game.RandomStartingCzar,
		game.BlackDeck(bGameCards),
		game.WhiteDeck(wGameCards),
	)
	sg := serverGame{state: s}
	sg.userToPlayers = make(map[model.User]*model.Player)
	for i, p := range p {
		user, _ := data.GetUserById(i)
		sg.userToPlayers[user] = p
	}
	games[gameid] = sg
	return nil
}

// TODO move to data: CreatePlayersFromUsers
func getPlayersForUsers(users ...model.User) []*model.Player {
	ret := make([]*model.Player, len(users))
	for i, u := range users {
		ret[i] = model.NewPlayer(u)
	}
	return ret
}

// For quick prototyping

func createTestGame() {
	createGame(getTestUsers(), "test")
}

func getTestUsers() []model.User {
	users := make([]model.User, 4)
	for i := 0; i < 4; i++ {
		u, _ := data.GetUserById(i)
		users[i] = u
	}
	return users
}
