package game

import (
	"math/rand"
	"time"

	"github.com/j4rv/golang-stuff/cah"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame(p []*cah.Player, opts ...Option) cah.Game {
	ret := cah.Game{
		Players:     p,
		HandSize:    10,
		DiscardPile: []cah.WhiteCard{},
	}
	applyOptions(&ret, opts...)
	shuffleB(&ret.BlackDeck)
	shuffleW(&ret.WhiteDeck)
	return ret
}

func RandomStartingCzar(s *cah.Game) {
	s.CurrCzarIndex = rand.Intn(len(s.Players))
}

func HandSize(size int) Option {
	return func(s *cah.Game) {
		s.HandSize = size
	}
}

func BlackDeck(bd []cah.BlackCard) Option {
	return func(s *cah.Game) {
		s.BlackDeck = bd
	}
}

func WhiteDeck(wd []cah.WhiteCard) Option {
	return func(s *cah.Game) {
		s.WhiteDeck = wd
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

type Option func(*cah.Game)

func applyOptions(s *cah.Game, opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}
