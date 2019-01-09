package data

import (
	"strings"
	"time"
)

const blankChar = "_"

type User struct {
	ID       int       `json:"id"       db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Creation time.Time `json:"creation" db:"creation"`
	// Games played, games won, cards played...
}

type Player struct {
	ID     int    `json:"id" db:"id"`
	User   User   `json:"user" db:"user"`
	Hand   []Card `json:"hand" db:"hand"`
	Points []Card `json:"points" db:"points"`
}

type Card struct {
	ID        int    `json:"id" db:"id"`
	IsBlack   bool   `json:"is_black" db:"is_black"`
	Text      string `json:"text" db:"text"`
	Expansion string `json:"expansion" db:"expansion"`
}

func (c Card) GetText() string {
	return c.Text
}

func (c Card) BlanksAmount() int {
	return strings.Count(c.Text, blankChar)
}

type Game struct {
	ID        int `json:"id" db:"id"`
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
