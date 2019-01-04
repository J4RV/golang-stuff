package cards

import (
	"fmt"
	"math/rand"
	"strings"
)

func PrettyCardsString(h []Card) string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = fmt.Sprintf("(%d) %s", i, h[i].String())
	}
	return strings.Join(strs, ", ")
}

//NewVanilla returns a vanilla set of UNO cards
func NewVanilla(opts ...Option) []Card {
	var cards []Card

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
func Shuffle(cards []Card) []Card {
	if cards == nil {
		return cards
	}
	res := make([]Card, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		res[i] = cards[j]
	}
	return res
}

//Filter returns a new slice of cards, containing only the cards that pass the filter function
func Filter(f func(c Card) bool) Option {
	return func(cards []Card) []Card {
		var res []Card
		for _, card := range cards {
			if f(card) {
				res = append(res, card)
			}
		}
		return res
	}
}

//Multiply returns an option that copies all cards in a deck X amount of times
func Multiply(n int) Option {
	return func(cards []Card) []Card {
		var res []Card
		for _, card := range cards {
			for i := 0; i < n; i++ {
				res = append(res, card)
			}
		}
		return res
	}
}

// internals

type Option func([]Card) []Card

func applyOptions(c []Card, opts ...Option) []Card {
	var res = c
	for _, opt := range opts {
		res = opt(res)
	}
	return res
}
