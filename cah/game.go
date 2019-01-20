package cah

type GameStore interface{}
type GameUsecases interface{}

type Game struct {
	ID       int       `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Password string    `json:"-" db:"password"`
	State    GameState `json:"state"`
}
