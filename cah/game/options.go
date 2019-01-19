package game

import (
	"log"
	"math/rand"

	"github.com/j4rv/golang-stuff/cah"
)

func (control GameController) Options() cah.GameOptions {
	return control.options
}

type Options struct{}

func applyOptions(s *cah.Game, opts ...cah.Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (_ Options) HandSize(size int) cah.Option {
	return func(s *cah.Game) {
		s.HandSize = size
	}
}

func (_ Options) BlackDeck(bd []cah.BlackCard) cah.Option {
	return func(s *cah.Game) {
		s.BlackDeck = bd
		shuffleB(&s.BlackDeck)
	}
}

func (_ Options) WhiteDeck(wd []cah.WhiteCard) cah.Option {
	return func(s *cah.Game) {
		s.WhiteDeck = wd
		shuffleW(&s.WhiteDeck)
	}
}

func (_ Options) RandomStartingCzar() cah.Option {
	return func(s *cah.Game) {
		if len(s.Players) == 0 {
			log.Println("WARNING Tried to call RandomStartingCzar using a game without players")
			return
		}
		s.CurrCzarIndex = rand.Intn(len(s.Players))
	}
}

func shuffleB(cards *[]cah.BlackCard) {
	if cards == nil {
		return
	}
	for i, j := range rand.Perm(len(*cards)) {
		(*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i]
	}
}

func shuffleW(cards *[]cah.WhiteCard) {
	if cards == nil {
		return
	}
	for i, j := range rand.Perm(len(*cards)) {
		(*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i]
	}
}
