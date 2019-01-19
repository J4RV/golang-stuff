package data

import (
	"fmt"
	"log"

	"github.com/j4rv/golang-stuff/cah"
)

var commonPass = "dev"
var commonPassHash, _ = getPassHash(commonPass)

type userMemStore struct {
	users map[int]*cah.User
}

func NewUserStore() userMemStore {
	return userMemStore{
		users: make(map[int]*cah.User),
	}
}

func (store userMemStore) ByCredentials(name, pass string) (cah.User, error) {
	u, err := store.UserByName(name)
	if err != nil {
		return u, err
	}
	if !correctPass(pass, u.Password) {
		return cah.User{}, fmt.Errorf("Incorrect password for user %s", u.Username)
	}
	return u, nil
}

func (store userMemStore) ByID(id int) (cah.User, error) {
	u, ok := store.users[id]
	if !ok {
		return cah.User{}, fmt.Errorf("Cannot find user with id %d", id)
	}
	return *u, nil
}

func (store userMemStore) UserByName(name string) (cah.User, error) {
	for _, u := range store.users {
		if u.Username == name {
			return *u, nil
		}
	}
	return cah.User{}, fmt.Errorf("Cant find user with username '%s'", name)
}

func PopulateUsers(store *userMemStore) {
	users := store.users
	users[0] = &cah.User{Username: "Red", ID: 0}
	users[1] = &cah.User{Username: "Green", ID: 1}
	users[2] = &cah.User{Username: "Blue", ID: 2}
	users[3] = &cah.User{Username: "Gold", ID: 3}
	users[4] = &cah.User{Username: "Silver", ID: 4}
	for i := range users {
		users[i].Password = commonPassHash
	}
	log.Print("mock users initialized")
}
