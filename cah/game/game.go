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
	shuffleB(&ret.BlackDeck)
	shuffleW(&ret.WhiteDeck)
	ret, _ = PutBlackCardInPlay(ret)
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

func shuffleB(cards *[]model.BlackCard) {
	if cards == nil {
		return
	}
	for i, j := range rand.Perm(len(*cards)) {
		(*cards)[i] = (*cards)[j]
	}
}

func shuffleW(cards *[]model.WhiteCard) {
	if cards == nil {
		return
	}
	for i, j := range rand.Perm(len(*cards)) {
		(*cards)[i] = (*cards)[j]
	}
}

type Option func(*model.State)

func applyOptions(s *model.State, opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}
