package main

type chooseWinnerPayload struct {
	Winner int `json:"winner"`
}

type playCardsPayload struct {
	CardIndexes []int `json:"cardIndexes"`
}

type loginPayload struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}
