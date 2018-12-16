package main

import (
	uno "github.com/j4rv/golang-stuff/unocards"
)

type State struct {
	Players       []*Player
	Currplayer    *Player
	Played        []uno.Card //Last card is the 'top' card
	Deck          []uno.Card
	Currcolor     uno.Color
	Skip          bool
	Orderreversed bool
	Drawacum      int
}

type Player struct {
	Name string
	Hand []uno.Card
}
