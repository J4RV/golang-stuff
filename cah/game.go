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
	Start(Game, ...Option) error
	//Start(gameID int, options ...Option) error
}

type Game struct {
	ID       int
	Owner    User
	UserID   int
	Users    []User `gorm:"many2many:game_users;"`
	Name     string
	Password string
	State    GameState
	StateID  int
}
