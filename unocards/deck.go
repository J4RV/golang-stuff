package unocards

import (
	"math/rand"
)

//Deck is just a slice of cards
type Deck []Card

//NewVanilla returns a vanilla set of UNO cards
func NewVanilla(opts ...option) Deck {
	var cards Deck

	for _, value := range allvalues {
		switch value {
		case Zero:
			// Zero: One of each color
			for _, color := range nonwildcolors {
				cards = append(cards, Card{Color: color, Value: value})
			}
		case Choose, DrawFour:
			// Wild cards: Four of each
			for i := 0; i < 4; i++ {
				cards = append(cards, Card{Color: Wild, Value: value})
			}
		default:
			// All others: Two of each color
			for _, color := range nonwildcolors {
				cards = append(cards, Card{Color: color, Value: value})
				cards = append(cards, Card{Color: color, Value: value})
			}
		}
	}

	return applyOptions(cards, opts...)
}

//Shuffle returns a copy of cards in a random order
//for true randomness, remember to set a seed to the rand package
func Shuffle(cards Deck) Deck {
	if cards == nil {
		return cards
	}
	res := make(Deck, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		res[i] = cards[j]
	}
	return res
}

//Filter returns a new slice of cards, containing only the cards that pass the filter function
func Filter(f func(c Card) bool) option {
	return func(cards Deck) Deck {
		var res Deck
		for _, card := range cards {
			if f(card) {
				res = append(res, card)
			}
		}
		return res
	}
}

//Multiply returns an option that copies all cards in a deck X amount of times
func Multiply(n int) option {
	return func(cards Deck) Deck {
		var res Deck
		for _, card := range cards {
			for i := 0; i < n; i++ {
				res = append(res, card)
			}
		}
		return res
	}
}

// internals

type option func(Deck) Deck

func applyOptions(c Deck, opts ...option) Deck {
	var res = c
	for _, opt := range opts {
		res = opt(res)
	}
	return res
}
