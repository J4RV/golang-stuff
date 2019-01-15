package data

import (
	"fmt"

	"github.com/j4rv/golang-stuff/cah/game"
)

var users = make(map[int]User)
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

func GetUserById(id int) (User, error) {
	u, ok := users[id]
	if !ok {
		return User{}, fmt.Errorf("Cannot find user with id %d", id)
	}
	return u, nil
}

func GetUserByLogin(n, p string) (User, error) {
	u, err := getUserByName(n)
	if err != nil {
		return u, err
	}
	if !correctPass(p, u.Password) {
		return User{}, fmt.Errorf("Incorrect password for user %s", u.Username)
	}
	return u, nil
}

func getUserByName(n string) (User, error) {
	for _, u := range users {
		if u.Username == n {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("Cant find user with username '%s'", n)
}

func initUsers() {
	users[0] = User{Username: "Red"}
	users[1] = User{Username: "Green"}
	users[2] = User{Username: "Blue"}
	users[3] = User{Username: "Gold"}
	users[4] = User{Username: "Silver"}
	for _, u := range users {
		u.Password = commonPassHash
	}
}
