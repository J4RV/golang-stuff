package cah

import "io"

type CardService interface {
	CreateWhite(text, expansion string) error
	CreateBlack(text, expansion string, blanks int) error
	GetWhites() []WhiteCard
	GetBlacks() []BlackCard
	CreateCards(wdat, bdat io.Reader, expansionName string) error
	CreateCardsFromFolder(folderPath, expansionName string) error
}

type CardController interface {
}

type Card struct {
	ID        int    `json:"id" db:"id"`
	Text      string `json:"text" db:"text"`
	Expansion string `json:"expansion" db:"expansion"`
}

type WhiteCard struct {
	Card
}

type BlackCard struct {
	Card
	BlanksAmount int `json:"blanksAmount" db:"blanksAmount"`
}
