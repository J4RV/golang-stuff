package main

type chooseWinnerPayload struct {
	Winner int `json:"winner"`
}

type playCardsPayload struct {
	CardIndexes []int `json:"card-indexes"`
}
