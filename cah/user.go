package cah

import (
	"time"
)

type UserStore interface {
	ByID(id int) (User, error)
	ByCredentials(name, pass string) (User, error)
}

type UserUsecases interface {
	ByID(id int) (User, error)
	Login(name, pass string) (User, error)
}

type User struct {
	ID       int       `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Creation time.Time `json:"creation" db:"creation"`
	// Games played, games won, cards played...
}
