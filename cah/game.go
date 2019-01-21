package cah

type GameStore interface {
	Create(Game) error
	ByID(int) (Game, error)
	ByStatePhase(Phase) []Game
	Update(Game) error
}

type GameUsecases interface {
	Create(owner User, name, pass string) error
	ByID(int) (Game, error)
	AllOpen() []Game
	UserJoins(User, Game) error
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
