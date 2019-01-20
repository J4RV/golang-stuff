package cah

type GameStore interface {
	Create(Game) error
	ByStatePhase(Phase) []Game
}
type GameUsecases interface {
	Create(Name, Pass string, state GameState) error
	AllOpen() []Game
}

type Game struct {
	ID       int       `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Password string    `json:"-" db:"password"`
	State    GameState `json:"state"`
}
