package usecase

import (
	"log"

	"github.com/j4rv/golang-stuff/cah"
)

func (control gameController) NewGame(opts ...cah.Option) cah.GameState {
	ret := cah.GameState{
		Players:     []*cah.Player{},
		HandSize:    10,
		DiscardPile: []cah.WhiteCard{},
		WhiteDeck:   []cah.WhiteCard{},
		BlackDeck:   []cah.BlackCard{},
	}
	applyOptions(&ret, opts...)
	ret, err := control.store.Create(ret)
	if err != nil {
		log.Println("Error while creating a new game:", err)
	}
	return ret
}
