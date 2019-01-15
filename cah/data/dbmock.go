package data

import (
	"fmt"

	"github.com/j4rv/golang-stuff/cah/game"
)

var users = make(map[string]User)
var blackCards []BlackCard
var whiteCards []WhiteCard
var games []Game

var commonPass = "dev"
var commonPassHash, _ = getPassHash(commonPass)

func init() {
	initUsers()
	initCards()
}

func GetBlackCards() []game.BlackCard {
	models := make([]game.BlackCard, len(blackCards))
	for i, c := range blackCards {
		models[i] = game.BlackCard(c)
	}
	return models
}

func GetWhiteCards() []game.WhiteCard {
	models := make([]game.WhiteCard, len(whiteCards))
	for i, c := range whiteCards {
		models[i] = game.WhiteCard(c)
	}
	return models
}

func GetUser(n, p string) (User, error) {
	u := users[n]
	if correctPass(p, u.Password) {
		return u, nil
	}
	return User{}, fmt.Errorf("Incorrect password for user %s", u.Username)
}

func initUsers() {
	users["Red"] = User{Username: "Red"}
	users["Green"] = User{Username: "Green"}
	users["Blue"] = User{Username: "Blue"}
	users["Gold"] = User{Username: "Gold"}
	users["Silver"] = User{Username: "Silver"}
	for _, u := range users {
		u.Password = commonPassHash
	}
}
