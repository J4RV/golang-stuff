package cah

import (
	"time"
)

type UserStore interface {
	Create(username, password string) (User, error)
	ByName(name string) (u User, ok bool)
	ByID(id int) (u User, ok bool)
}

type UserUsecases interface {
	Register(username, password string) (User, error)
	Login(name, pass string) (u User, ok bool)
	ByID(id int) (u User, ok bool)
}

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Creation time.Time `json:"-"`
	// Games played, games won, cards played...
}
