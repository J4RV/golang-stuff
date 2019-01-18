package data

import (
	"errors"
	"fmt"

	"github.com/j4rv/golang-stuff/cah/model"
)

type CardRepository interface {
	CreateWhite(text, expansion string) error
	CreateBlack(text, expansion string, blanks int) error
	GetWhites() []model.WhiteCard
	GetBlacks() []model.BlackCard
}

type CardMemStore struct {
	whiteCards []model.WhiteCard
	blackCards []model.BlackCard
}

func NewCardStore() *CardMemStore {
	return &CardMemStore{
		whiteCards: []model.WhiteCard{},
		blackCards: []model.BlackCard{},
	}
}

func (s *CardMemStore) CreateWhite(t, e string) error {
	c := model.WhiteCard{}
	c.Text = t
	c.Expansion = e
	err := validateCard(c.Card)
	if err == nil {
		s.whiteCards = append(s.whiteCards, c)
	}
	return err
}

func (s *CardMemStore) CreateBlack(t, e string, blanks int) error {
	if blanks < 1 {
		return errors.New("Black cards need to have at least 1 blank")
	}
	if blanks > 5 {
		return fmt.Errorf("Black cards blanks maximum is five, but got %d", blanks)
	}
	c := model.BlackCard{}
	c.Text = t
	c.Expansion = e
	c.BlanksAmount = blanks
	err := validateCard(c.Card)
	if err == nil {
		s.blackCards = append(s.blackCards, c)
	}
	return err
}

func (s *CardMemStore) GetWhites() []model.WhiteCard {
	return s.whiteCards
}

func (s *CardMemStore) GetBlacks() []model.BlackCard {
	return s.blackCards
}

func validateCard(c model.Card) error {
	if len(c.Text) == 0 {
		return errors.New("Card text cannot be empty")
	}
	if len(c.Text) > 120 {
		return errors.New("Card text cannot be longer than 120")
	}
	return nil
}
