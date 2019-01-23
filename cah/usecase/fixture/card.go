package fixture

import (
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

func getWhiteCardsFixture(amount int) []cah.WhiteCard {
	ret := make([]cah.WhiteCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = cah.WhiteCard{Card: cah.Card{Text: fmt.Sprintf("White card fixture (%d)", i)}}
	}
	return ret
}

func getBlackCardsFixture(amount int) []cah.BlackCard {
	ret := make([]cah.BlackCard, amount)
	for i := 0; i < amount; i++ {
		ret[i] = cah.BlackCard{Card: cah.Card{Text: fmt.Sprintf("Black card fixture (%d)", i)}, BlanksAmount: 1}
	}
	return ret
}
