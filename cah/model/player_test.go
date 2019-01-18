package model

import (
	"fmt"
	"strings"
	"testing"

	"github.com/j4rv/golang-stuff/cah/game"
	"github.com/stretchr/testify/assert"
)

func TestPlayer_extractCardFromHand(t *testing.T) {
	p := Player{}
	c1 := WhiteCard{Card{Text: "A"}}
	c2 := WhiteCard{Card{Text: "B"}}
	c3 := WhiteCard{Card{Text: "C"}}
	p.Hand = []WhiteCard{c1, c2, c3}

	_, err := p.extractCardFromHand(-1)
	if err == nil {
		t.Error("Expected error but did not found it, negative index")
	}

	_, err = p.extractCardFromHand(9)
	if err == nil {
		t.Error("Expected error but did not found it, index over hand size")
	}

	c, err := p.extractCardFromHand(1)
	assert.Equal(t, c.Text, "B", "Unexpected text in extracted hand")
	assert.Equalf(t, err, nil, "Unexpected error %v", err)
	assert.Equalf(t, len(p.Hand), 2, "Hand size did not get reduced, hand: %s, len: ", p.Hand)
	expectedResultHand := []WhiteCard{c1, c3}

	for i := range p.Hand {
		assert.Equalf(t, p.Hand[i], expectedResultHand[i], "Unexpected hand card at position %d", i)
	}
}

func TestPlayer_extractCardsFromHand(t *testing.T) {
	assert := assert.New(t)
	p := Player{}
	p.Hand = game.getWhiteCardsFixture(10)
	indexes := []int{0, 9, 5}
	cards, err := p.extractCardsFromHand(indexes)

	assert.NoError(err)
	assert.Equal(7, len(p.Hand), "Unexpected hand size")
	assert.Equal(3, len(cards), "Unexpected cards size")

	for i, index := range indexes {
		assert.True(strings.Contains(cards[i].Text, fmt.Sprintf("(%d)", index)),
			"Unexpected card order in cards result", cards[i].Text)
	}

	expectedHand := []int{1, 2, 3, 4, 6, 7, 8}
	for i, index := range expectedHand {
		assert.True(strings.Contains(p.Hand[i].Text, fmt.Sprintf("(%d)", index)),
			"Unexpected card order in hand", p.Hand[i].Text)
	}
}
