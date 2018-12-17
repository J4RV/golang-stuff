package main

import (
	"log"

	"github.com/j4rv/golang-stuff/uno/cards"
)

func Start(s State) State {
	res := clone(s)

	// Fill player hands
	for i := 0; i < 7; i++ {
		for p := range res.Players {
			moveFromTo(&res.Deck, &(res.Players[p]).Hand)
		}
	}

	// Play first card in discard pile
	moveFromTo(&res.Deck, &res.DiscardPile)
	for res.TopCard().Value == cards.DrawFour {
		// First card cannot be Draw 4, try again
		moveFromTo(&res.Deck, &res.DiscardPile)
	}

	// process first card
	res = ProcessTopCard(res)
	return res
}

func ProcessTopCard(s State) State {
	res := clone(s)

	log.Println("Pile:", res.DiscardPile)
	log.Println("Processing card:", res.TopCard())
	switch res.TopCard().Value {
	case cards.Skip:
		log.Println("Skipping next player")
		res = NextPlayer(res)
	case cards.Reverse:
		res.OrderReversed = !res.OrderReversed
	case cards.DrawTwo:
		res.DrawAcum += 2
	case cards.DrawFour:
		res.DrawAcum += 4
	}

	res.CurrColor = res.TopCard().Color

	if res.CurrColor == cards.Wild {
		return ChooseColor(res)
	}

	return NextPlayer(res)
}

func ProcessForcedDraw(s State) State {
	res := clone(s)

	for i := 0; i < res.DrawAcum; i++ {
		moveFromTo(&res.Deck, &res.CurrPlayer().Hand)
	}
	// this was a forced draw, clear the draw acum and automatically pass
	res.DrawAcum = 0

	res = NextPlayer(res)
	return res
}

func NextPlayer(s State) State {
	res := clone(s)

	if res.OrderReversed {
		res.CurrPlayerIndex--
		if res.CurrPlayerIndex < 0 {
			res.CurrPlayerIndex = len(res.Players) - 1
		}
	} else {
		res.CurrPlayerIndex++
		if res.CurrPlayerIndex == len(res.Players) {
			res.CurrPlayerIndex = 0
		}
	}

	log.Println("Next player:", res.CurrPlayer().Name)
	res.Phase = PlayerTurnStarting
	return res
}
