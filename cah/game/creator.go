package game

import (
	"github.com/j4rv/golang-stuff/cah"
)

func (_ GameController) NewGame(opts ...cah.Option) cah.Game {
	ret := cah.Game{
		Players:     []*cah.Player{},
		HandSize:    10,
		DiscardPile: []cah.WhiteCard{},
		WhiteDeck:   []cah.WhiteCard{},
		BlackDeck:   []cah.BlackCard{},
	}
	applyOptions(&ret, opts...)
	return ret
}
