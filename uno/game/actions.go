package game

import (
	"log"

	"github.com/j4rv/golang-stuff/uno/cards"
)

func Start(s State) State {
	res := clone(s)
	// Fill player hand
	// TODO fill all players hands
	for i := 0; i < 7; i++ {
		for p := range res.Players {
			moveFromTo(&res.Deck, &(res.Players[p]).Hand)
		}
	}
	// Play first card in discard pile
	moveFromTo(&res.Deck, &res.DiscardPile)
	log.Printf("First top card: %v\n", res.TopCard())

	// process first card
	res = ProcessTopCard(res)
	return res
}

func End(s State) State {
	res := clone(s)
	res.Phase = Finished
	return res
}

func PlayCard(i uint64, s State) State {
	if i < 0 || int(i) >= len(s.CurrPlayer().Hand) {
		log.Printf("Card index out of hand: %v\n", i)
		return s
	}
	c := s.CurrPlayer().Hand[i]
	if !canplay(c, s) {
		log.Printf("Not a valid card to play: %v\n", c)
		log.Printf("Current color %s, current value: %s\n", s.TopCard().Color, s.TopCard().Value)
		return s
	}

	res := clone(s)
	log.Printf("playing card: %v", c)
	playcard(i, &res)
	res.Phase = ProcessingPlayedCard
	return res
}

func Draw(s State) State {
	res := clone(s)
	if res.DrawAcum > 0 {
		for i := 0; i < res.DrawAcum; i++ {
			draw(&res)
		}
		// this was a forced draw, clear the draw acum and automatically pass
		res.DrawAcum = 0
		res = Pass(res)
	} else {
		if res.DrawnThisTurn {
			log.Println("Current player already drew this turn")
			return res
		}
		// player wants to draw
		draw(&res)
	}

	return res
}

func ChangeColor(c cards.Color, s State) State {
	res := clone(s)
	res.CurrColor = c
	log.Println("Current color changed to:", c)
	res.Phase = PlayerTurn
	nextplayer(&res)
	return res
}

func Pass(s State) State {
	if !s.DrawnThisTurn {
		log.Println("You need to draw before you can pass")
		return s
	}
	res := clone(s)
	nextplayer(&res)
	return res
}

func ProcessTopCard(s State) State {
	res := clone(s)
	topc := res.TopCard()
	switch topc.Value {
	case cards.Skip:
		res.Skip = true
	case cards.Reverse:
		res.OrderReversed = true
	case cards.DrawTwo:
		res.DrawAcum += 2
	case cards.DrawFour:
		res.DrawAcum += 4
	}
	res.CurrColor = topc.Color
	if res.DrawAcum > 0 {
		res.Phase = ForcingDraw
	} else {
		if res.CurrColor == cards.Wild {
			res.Phase = ChoosingColor
		} else {
			res.Phase = PlayerTurn
			nextplayer(&res)
		}
	}
	return res
}

// internal functions that alter a state

func nextplayer(s *State) {
	// TODO have reverse order into account
	s.CurrPlayerIndex++
	if s.CurrPlayerIndex >= len(s.Players) {
		s.CurrPlayerIndex = 0
	}
	s.DrawnThisTurn = false
}

func playcard(i uint64, s *State) {
	h := &s.CurrPlayer().Hand
	c := s.CurrPlayer().Hand[i]
	*h = append((*h)[:i], (*h)[i+1:]...)
	s.DiscardPile = append(s.DiscardPile, c)
}

func draw(s *State) {
	drawn := moveFromTo(&s.Deck, &s.CurrPlayer().Hand)
	s.DrawnThisTurn = true
	log.Printf("Card drawn: %v\n", drawn)
}

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
		Skip:            s.Skip,
		OrderReversed:   s.OrderReversed,
		DrawAcum:        s.DrawAcum,
		DrawnThisTurn:   s.DrawnThisTurn,
	}
	copy(ret.Players, s.Players)
	copy(ret.DiscardPile, s.DiscardPile)
	copy(ret.Deck, s.Deck)
	return ret
}
