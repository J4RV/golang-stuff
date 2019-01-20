package cah

import (
	"time"
)

type UserStore interface {
	Create(username, password string) (User, error)
	ByCredentials(name, pass string) (User, error)
	ByName(name string) (u User, ok bool)
	ByID(id int) (u User, ok bool)
}

type UserUsecases interface {
	Register(username, password string) (User, error)
	Login(name, pass string) (u User, ok bool)
	ByID(id int) (u User, ok bool)
}

type User struct {
	ID       int       `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"-" db:"password"`
	Creation time.Time `json:"-" db:"creation"`
	// Games played, games won, cards played...
}
