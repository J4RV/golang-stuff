package mem

import (
	"time"

	"github.com/j4rv/golang-stuff/cah"
)

type userMemStore struct {
	abstractMemStore
	users map[int]*cah.User
}

func NewUserStore() *userMemStore {
	return &userMemStore{
		users: make(map[int]*cah.User),
	}
}

func (store *userMemStore) Create(username, password string) (cah.User, error) {
	store.lock()
	defer store.release()
	user := cah.User{}
	user.Username = username
	user.Password = password
	user.Creation = time.Now()
	user.ID = store.nextID()
	store.users[user.ID] = &user
	return user, nil
}

func (store *userMemStore) ByID(id int) (cah.User, bool) {
	store.lock()
	defer store.release()
	u, ok := store.users[id]
	if !ok {
		return cah.User{}, false
	}
	return *u, true
}

func (store *userMemStore) ByName(name string) (cah.User, bool) {
	store.lock()
	defer store.release()
	for _, u := range store.users {
		if u.Username == name {
			return *u, true
		}
	}
	return cah.User{}, false
}
