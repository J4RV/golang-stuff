package game

import (
	"math/rand"
	"time"
)

const playerHandSize = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame(bd []BlackCard, wd []WhiteCard, p []*Player, opts ...Option) State {
	ret := State{
		BlackDeck:   shuffleB(bd),
		WhiteDeck:   shuffleW(wd),
		Players:     p,
		DiscardPile: []WhiteCard{},
	}
	playersDraw(&ret)
	ret = applyOptions(ret, opts...)
	ret, _ = PutBlackCardInPlay(ret)
	return ret
}

func shuffleB(cards []BlackCard) []BlackCard {
	if cards == nil {
		return cards
	}
	res := make([]BlackCard, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		res[i] = cards[j]
	}
	return res
}

func shuffleW(cards []WhiteCard) []WhiteCard {
	if cards == nil {
		return cards
	}
	res := make([]WhiteCard, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		res[i] = cards[j]
	}
	return res
}

func RandomStartingCzar(s State) State {
	res := s
	res.CurrCzarIndex = rand.Intn(len(s.Players))
	return res
}

type Option func(State) State

func applyOptions(s State, opts ...Option) State {
	res := s
	for _, opt := range opts {
		res = opt(res)
	}
	return res
}
