package cah

type GameStore interface {
	Create(Game) error
	ByStatePhase(Phase) []Game
}
type GameUsecases interface {
	Create(owner User, name, pass string, expansions []string, state GameState) error
	AllOpen() []Game
}

type Game struct {
	ID         int      `json:"id" db:"id"`
	OwnerID    int      `json:"ownerID" db:"ownerID"`
	Name       string   `json:"name" db:"name"`
	Password   string   `json:"-" db:"password"`
	Expansions []string `json:"expansions" db:"expansions"`
	StateID    int      `json:"stateID" db:"stateID"`
}
