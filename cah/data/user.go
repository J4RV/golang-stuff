package data

import (
	"fmt"
	"log"
	"time"

	"github.com/j4rv/golang-stuff/cah"
	"golang.org/x/crypto/bcrypt"
)

var commonPass = "dev"

type userMemStore struct {
	lastID int
	users  map[int]*cah.User
}

func NewUserStore() *userMemStore {
	return &userMemStore{
		users: make(map[int]*cah.User),
	}
}

func (store *userMemStore) Create(username, password string) (cah.User, error) {
	user := cah.User{}
	passhash, err := getPassHash(password)
	if err != nil {
		return user, err
	}
	user.Username = username
	user.Password = passhash
	user.Creation = time.Now()
	user.ID = store.lastID
	store.lastID++
	store.users[user.ID] = &user
	return user, nil
}

func (store *userMemStore) ByCredentials(name, pass string) (cah.User, error) {
	u, ok := store.ByName(name)
	if !ok {
		return u, fmt.Errorf("There is no user with username '%s'", u.Username)
	}
	if !correctPass(pass, u.Password) {
		return cah.User{}, fmt.Errorf("Incorrect password for user '%s'", u.Username)
	}
	return u, nil
}

func (store *userMemStore) ByID(id int) (cah.User, bool) {
	u, ok := store.users[id]
	if !ok {
		return cah.User{}, false
	}
	return *u, true
}

func (store *userMemStore) ByName(name string) (cah.User, bool) {
	for _, u := range store.users {
		if u.Username == name {
			return *u, true
		}
	}
	return cah.User{}, false
}

// internal

func getPassHash(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(b), err
}

func correctPass(pass string, storedhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedhash), []byte(pass))
	return err == nil
}

// for testing

func PopulateUsers(store *userMemStore) {
	store.Create("Red", commonPass)
	store.Create("Green", commonPass)
	store.Create("Blue", commonPass)
	store.Create("Gold", commonPass)
	store.Create("Silver", commonPass)
	log.Print("Base users initialized")
}
