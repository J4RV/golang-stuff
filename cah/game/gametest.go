package game

import (
	"fmt"
)

type card struct {
	text string
}

type blackCard struct {
	card
	blanks int
}

func (c card) GetText() string {
	return c.text
}

func (c blackCard) GetBlanksAmount() int {
	return c.blanks
}

func getWhiteCardsFixture(amount int) []WhiteCard {
	ret := make([]WhiteCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = card{text: fmt.Sprintf("White card fixture (%d)", i)}
	}
	return ret
}

func getBlackCardsFixture(amount int) []BlackCard {
	ret := make([]BlackCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = blackCard{card{text: fmt.Sprintf("Black card fixture (%d)", i)}, 1}
	}
	return ret
}

func getPlayerFixture(name string) *Player {
	return &Player{
		Name:             name,
		Hand:             []WhiteCard{},
		WhiteCardsInPlay: []WhiteCard{},
		Points:           []BlackCard{},
	}
}

func getStateFixture() State {
	return State{
		BlackDeck:   getBlackCardsFixture(20),
		WhiteDeck:   getWhiteCardsFixture(40),
		DiscardPile: []WhiteCard{},
		Players: []*Player{
			getPlayerFixture("Player1"),
			getPlayerFixture("Player2"),
			getPlayerFixture("Player3"),
		},
	}
}
