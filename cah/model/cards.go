package model

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
