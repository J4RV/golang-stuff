package game

import "math/rand"

func NewGame(bd []BlackCard, wd []WhiteCard, p []Player, opts ...Option) State {
	return applyOptions(State{
		BlackDeck: bd,
		WhiteDeck: wd,
		Players:   p,
	}, opts...)
}

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
