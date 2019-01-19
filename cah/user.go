package cah

import (
	"time"
)

type UserService interface {
}

type User struct {
	ID       int       `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Creation time.Time `json:"creation" db:"creation"`
	// Games played, games won, cards played...
}
