package data

import (
	"time"
)

const blankChar = "_"

type DBObject struct {
	ID int `json:"id"       db:"id"`
}

type User struct {
	DBObject
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Creation time.Time `json:"creation" db:"creation"`
	// Games played, games won, cards played...
}

type Player struct {
	DBObject
	User   User   `json:"user" db:"user"`
	Hand   []Card `json:"hand" db:"hand"`
	Points []Card `json:"points" db:"points"`
}

type Card struct {
	DBObject
	Text      string `json:"text" db:"text"`
	Expansion string `json:"expansion" db:"expansion"`
}

type BlackCard struct {
	Card
	BlanksAmount int `json:"blanks-amount" db:"blanks-amount"`
}

type WhiteCard struct {
	Card
}

func (c Card) GetText() string {
	return c.Text
}

func (c BlackCard) GetBlanksAmount() int {
	return c.BlanksAmount
}

type Game struct {
	DBObject
	State     State
	Players   []*Player
	BlackDeck []Card
	WhiteDeck []Card
	// Chat?
}

type State uint

const (
	Starting State = iota
	WaitingForAnswers
	CzarChoosingWinner
	/*Additional states from custom gamemodes*/
	Finished
)
