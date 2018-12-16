package main

import (
	uno "github.com/j4rv/golang-stuff/unocards"
)

type State struct {
	Players       []Player
	Currplayer    Player
	Board         Board
	Skip          bool
	Orderreversed bool
	Drawacum      int
}

type Player struct {
	Name string
	Hand []uno.Card
}

type Board struct {
	Played []uno.Card //Last card is the 'top' card
	Deck   uno.Deck
}
