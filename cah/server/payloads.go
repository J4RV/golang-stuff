package main

type chooseWinnerPayload struct {
	Winner int `json:"winner"`
}

type playCardsPayload struct {
	CardIndexes []int `json:"cardIndexes"`
}

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
