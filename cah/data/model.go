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
	BlanksAmount int `json:"blanksAmount" db:"blanksAmount"`
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
	State     State     `json:"state" db:"state"`
	Players   []*Player `json:"players" db:"players"`
	BlackDeck []Card    `json:"blackDeck" db:"blackDeck"`
	WhiteDeck []Card    `json:"whiteDeck" db:"whiteDeck"`
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
