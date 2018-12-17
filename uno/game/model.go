package main

import (
	"github.com/j4rv/golang-stuff/uno/cards"
)

type Phase uint8

const (
	Starting Phase = iota
	PlayerTurnStarting
	PlayerThinkingAfterDraw
	ProcessingCard
	ProcessingForcedDraw
	PlayerChoosingColor
	Finished
)

type Player struct {
	Name string
	Hand []cards.Card
}

type State struct {
	Players         []Player
	Deck            []cards.Card
	Phase           Phase
	CurrPlayerIndex int
	DiscardPile     []cards.Card //Last card is the 'top' card
	CurrColor       cards.Color
	OrderReversed   bool
	DrawAcum        int
}

func (s State) TopCard() cards.Card {
	cards := s.DiscardPile
	if len(cards) == 0 {
		panic("no cards to process, should never happen")
	}
	return cards[len(cards)-1]
}

func (s State) CurrPlayer() *Player {
	return &s.Players[s.CurrPlayerIndex]
}
