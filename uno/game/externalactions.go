package main

import (
	"log"

	"github.com/j4rv/golang-stuff/uno/cards"
)

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

	c = res.CurrPlayer().Hand[i]
	h := res.CurrPlayer().Hand
	res.CurrPlayer().Hand = append(h[:i], h[i+1:]...)
	res.DiscardPile = append(res.DiscardPile, c)

	if len(res.CurrPlayer().Hand) == 0 {
		res = End(res)
	} else {
		res.Phase = ProcessingCard
	}

	return res
}

func DrawOne(s State) State {
	res := clone(s)

	moveFromTo(&res.Deck, &res.CurrPlayer().Hand)

	res.Phase = PlayerThinkingAfterDraw
	return res
}

func ForcedDraw(s State) State {
	res := clone(s)

	res.Phase = ProcessingForcedDraw
	return res
}

func ChangeColor(c cards.Color, s State) State {
	res := clone(s)

	res.CurrColor = c
	log.Println("Current color changed to:", c)

	res = NextPlayer(res)
	return res
}

func Pass(s State) State {
	if s.Phase == PlayerTurnStarting {
		log.Println("You need to draw before you can pass")
		return s
	}
	if s.Phase != PlayerThinkingAfterDraw {
		panic("called Pass with an invalid state")
	}
	return NextPlayer(s)
}
