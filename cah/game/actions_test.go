package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextCzar(t *testing.T) {
	assert := assert.New(t)
	s := getStateFixture()
	s.Phase = CzarChoosingWinner
	assert.Equal(s.CurrCzarIndex, 0, "Unexpected first czar")

	s, err := NextCzar(s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 1, "Unexpected second czar")

	s, err = NextCzar(s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 2, "Unexpected third czar")

	s, err = NextCzar(s)
	assert.Equalf(err, nil, "Unexpected error: %v", err)
	assert.Equal(s.CurrCzarIndex, 0, "Unexpected fourth czar")
}

func TestNextCzar_errors(t *testing.T) {
	assert := assert.New(t)
	s := getStateFixture()

	s.BlackCardInPlay = s.BlackDeck[0]
	s.Phase = CzarChoosingWinner
	s, err := NextCzar(s)
	assert.NotEqual(err, nil, "Expected 'black card in play' error but found nil")

	s.Phase = Finished
	s.BlackCardInPlay = nil
	s, err = NextCzar(s)
	assert.NotEqual(err, nil, "Expected 'incorrect phase' error but found nil")
}
