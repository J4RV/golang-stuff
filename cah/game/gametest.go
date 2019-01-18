package game

import (
	"fmt"

	"github.com/j4rv/golang-stuff/cah/model"
)

func getWhiteCardsFixture(amount int) []model.WhiteCard {
	ret := make([]model.WhiteCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = model.WhiteCard{model.Card{Text: fmt.Sprintf("White card fixture (%d)", i)}}
	}
	return ret
}

func getBlackCardsFixture(amount int) []model.BlackCard {
	ret := make([]model.BlackCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = model.BlackCard{model.Card{Text: fmt.Sprintf("Black card fixture (%d)", i)}, 1}
	}
	return ret
}

func getPlayerFixture(name string) *model.Player {
	return &model.Player{
		User:             model.User{Username: name},
		Hand:             []model.WhiteCard{},
		WhiteCardsInPlay: []model.WhiteCard{},
		Points:           []model.BlackCard{},
	}
}

func getStateFixture() model.State {
	return model.State{
		BlackDeck:   getBlackCardsFixture(20),
		WhiteDeck:   getWhiteCardsFixture(40),
		DiscardPile: []model.WhiteCard{},
		Players: []*model.Player{
			getPlayerFixture("Player1"),
			getPlayerFixture("Player2"),
			getPlayerFixture("Player3"),
		},
	}
}
