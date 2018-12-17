package main

import (
	"github.com/j4rv/golang-stuff/uno/cards"
)

func moveFromTo(from *[]cards.Card, to *[]cards.Card) cards.Card {
	var c cards.Card
	c = (*from)[0]
	*from = (*from)[1:]
	*to = append(*to, c)
	return c
}

func canplay(c cards.Card, s State) bool {
	topc := s.TopCard()
	if s.DrawAcum == 0 {
		if c.Color == cards.Wild {
			return true
		}
		if topc.Color == cards.Wild {
			return c.Color == s.CurrColor
		}
		if c.Color == topc.Color {
			return true
		}
	}
	// even if Drawacum is not zero, you can answer a draw two with another draw two
	// same for wild draw four
	return c.Value == topc.Value
}

func clone(s State) State {
	ret := State{
		Phase:           s.Phase,
		Players:         make([]Player, len(s.Players)),
		CurrPlayerIndex: s.CurrPlayerIndex,
		DiscardPile:     make([]cards.Card, len(s.DiscardPile)),
		Deck:            make([]cards.Card, len(s.Deck)),
		CurrColor:       s.CurrColor,
		OrderReversed:   s.OrderReversed,
		DrawAcum:        s.DrawAcum,
	}
	copy(ret.Players, s.Players)
	copy(ret.DiscardPile, s.DiscardPile)
	copy(ret.Deck, s.Deck)
	return ret
}
