package main

import (
	. "github.com/j4rv/golang-stuff/unocards"
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
	Hand []Card
}

type Board struct {
	Topcard Card
	Discard []Card
	Deck    Deck
}
