package game

import (
	"fmt"
	"testing"
)

func TestPlayerRemoveCard(t *testing.T) {
	p := Player{}
	c1 := cardMock{Text: "A"}
	c2 := cardMock{Text: "B"}
	c3 := cardMock{Text: "C"}
	p.Hand = []WhiteCard{c1, c2, c3}

	err := p.removeCardFromHand(-1)
	if err == nil {
		t.Error("Expected error but did not found it, negative index")
	}

	err = p.removeCardFromHand(9)
	if err == nil {
		t.Error("Expected error but did not found it, index over hand size")
	}

	err = p.removeCardFromHand(1)
	expectedResultHand := []WhiteCard{c1, c3}
	if err != nil {
		t.Error(err)
	}
	if len(p.Hand) != 2 {
		msg := fmt.Sprintf("Hand size did not get reduced, hand: %s, len: ", p.Hand)
		t.Error(msg)
	}

	for i := range p.Hand {
		if p.Hand[i] != expectedResultHand[i] {
			t.Error(fmt.Sprintf("Unexpected hand card at position %d", i))
		}
	}
}

type cardMock struct {
	Text string
}

func (c cardMock) text() string {
	return c.Text
}
