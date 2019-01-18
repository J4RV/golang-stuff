package data

import (
	"fmt"
	"log"

	"github.com/j4rv/golang-stuff/cah/model"
)

var users = make(map[int]*model.User)

var commonPass = "dev"
var commonPassHash, _ = getPassHash(commonPass)

func init() {
	initUsers()
}

func GetUserById(id int) (model.User, error) {
	u, ok := users[id]
	if !ok {
		return model.User{}, fmt.Errorf("Cannot find user with id %d", id)
	}
	return *u, nil
}

func GetUserByLogin(n, p string) (model.User, error) {
	u, err := getUserByName(n)
	if err != nil {
		return u, err
	}
	if !correctPass(p, u.Password) {
		return model.User{}, fmt.Errorf("Incorrect password for user %s", u.Username)
	}
	return u, nil
}

func getUserByName(n string) (model.User, error) {
	for _, u := range users {
		if u.Username == n {
			return *u, nil
		}
	}
	return model.User{}, fmt.Errorf("Cant find user with username '%s'", n)
}

func initUsers() {
	users[0] = &model.User{Username: "Red", ID: 0}
	users[1] = &model.User{Username: "Green", ID: 1}
	users[2] = &model.User{Username: "Blue", ID: 2}
	users[3] = &model.User{Username: "Gold", ID: 3}
	users[4] = &model.User{Username: "Silver", ID: 4}
	for i := range users {
		users[i].Password = commonPassHash
	}
	log.Print("mock users initialized")
}
