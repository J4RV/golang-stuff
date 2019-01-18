package game

import (
	"math/rand"
	"time"

	"github.com/j4rv/golang-stuff/cah/model"
)

const playerHandSize = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame(p []*model.Player, opts ...Option) model.State {
	ret := model.State{
		Players:     p,
		DiscardPile: []model.WhiteCard{},
	}
	applyOptions(&ret, opts...)
	ret, _ = PutBlackCardInPlay(ret)
	ret.BlackDeck = shuffleB(ret.BlackDeck)
	ret.WhiteDeck = shuffleW(ret.WhiteDeck)
	playersDraw(&ret)
	return ret
}

func RandomStartingCzar(s *model.State) {
	s.CurrCzarIndex = rand.Intn(len(s.Players))
}

func BlackDeck(bd []model.BlackCard) Option {
	return func(s *model.State) {
		s.BlackDeck = bd
	}
}

func WhiteDeck(wd []model.WhiteCard) Option {
	return func(s *model.State) {
		s.WhiteDeck = wd
	}
}

func shuffleB(cards []model.BlackCard) []model.BlackCard {
	if cards == nil {
		return cards
	}
	res := make([]model.BlackCard, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		res[i] = cards[j]
	}
	return res
}

func shuffleW(cards []model.WhiteCard) []model.WhiteCard {
	if cards == nil {
		return cards
	}
	res := make([]model.WhiteCard, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		res[i] = cards[j]
	}
	return res
}

type Option func(*model.State)

func applyOptions(s *model.State, opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}
