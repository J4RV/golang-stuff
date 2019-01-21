package cah

type GameStore interface {
	Create(Game) error
	ByStatePhase(Phase) []Game
}

type GameUsecases interface {
	Create(owner User, name, pass string) error
	AllOpen() []Game
	//TODO Join(Game) error
	//Start(gameID int, options ...Option) error
}

type Game struct {
	ID       int    `json:"id" db:"id"`
	OwnerID  int    `json:"ownerID" db:"ownerID"`
	Users    []User `json:"users" gorm:"many2many:game_users;"`
	Name     string `json:"name" db:"name"`
	Password string `json:"-" db:"password"`
	StateID  int    `json:"stateID" db:"stateID"`
}
