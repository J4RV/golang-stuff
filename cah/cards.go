package cah

import "io"

type CardStore interface {
	CreateWhite(text, expansion string) error
	CreateBlack(text, expansion string, blanks int) error
	AllWhites() []WhiteCard
	AllBlacks() []BlackCard
	ExpansionWhites(...string) []WhiteCard
	ExpansionBlacks(...string) []BlackCard
	AvailableExpansions() []string
}

type CardUsecases interface {
	CreateFromReaders(wdat, bdat io.Reader, expansionName string) error
	CreateFromFolder(folderPath, expansionName string) error
	AllWhites() []WhiteCard
	AllBlacks() []BlackCard
	ExpansionWhites(...string) []WhiteCard
	ExpansionBlacks(...string) []BlackCard
	AvailableExpansions() []string
}

type Card struct {
	ID        int    `json:"-" db:"id"`
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
