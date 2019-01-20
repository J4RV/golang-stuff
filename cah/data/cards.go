package data

import (
	"errors"
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

type cardMemStore struct {
	abstractMemStore
	whiteCards []cah.WhiteCard
	blackCards []cah.BlackCard
}

func NewCardStore() *cardMemStore {
	return &cardMemStore{
		whiteCards: []cah.WhiteCard{},
		blackCards: []cah.BlackCard{},
	}
}

func (s *cardMemStore) CreateWhite(t, e string) error {
	c := cah.WhiteCard{}
	c.ID = s.nextID()
	c.Text = t
	c.Expansion = e
	err := validateCard(c.Card)
	if err != nil {
		return err
	}
	s.whiteCards = append(s.whiteCards, c)
	return nil
}

func (s *cardMemStore) CreateBlack(t, e string, blanks int) error {
	if blanks < 1 {
		return errors.New("Black cards need to have at least 1 blank")
	}
	if blanks > 5 {
		return fmt.Errorf("Black cards blanks maximum is five, but got %d", blanks)
	}
	c := cah.BlackCard{}
	c.ID = s.nextID()
	c.Text = t
	c.Expansion = e
	c.BlanksAmount = blanks
	err := validateCard(c.Card)
	if err != nil {
		return err
	}
	s.blackCards = append(s.blackCards, c)
	return nil
}

func (s *cardMemStore) AllWhites() []cah.WhiteCard {
	return s.whiteCards
}

func (s *cardMemStore) AllBlacks() []cah.BlackCard {
	return s.blackCards
}

func validateCard(c cah.Card) error {
	if len(c.Text) == 0 {
		return errors.New("Card text cannot be empty")
	}
	if len(c.Text) > 120 {
		return errors.New("Card text cannot be longer than 120")
	}
	return nil
}
