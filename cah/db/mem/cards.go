package mem

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

func (store *cardMemStore) CreateWhite(t, e string) error {
	store.lock()
	defer store.release()
	c := cah.WhiteCard{}
	c.ID = store.nextID()
	c.Text = t
	c.Expansion = e
	err := validateCard(c.Card)
	if err != nil {
		return err
	}
	store.whiteCards = append(store.whiteCards, c)
	return nil
}

func (store *cardMemStore) CreateBlack(t, e string, blanks int) error {
	store.lock()
	defer store.release()
	if blanks < 1 {
		return errors.New("Black cards need to have at least 1 blank")
	}
	if blanks > 5 {
		return fmt.Errorf("Black cards blanks maximum is five, but got %d", blanks)
	}
	c := cah.BlackCard{}
	c.ID = store.nextID()
	c.Text = t
	c.Expansion = e
	c.BlanksAmount = blanks
	err := validateCard(c.Card)
	if err != nil {
		return err
	}
	store.blackCards = append(store.blackCards, c)
	return nil
}

func (store *cardMemStore) AllWhites() []cah.WhiteCard {
	store.lock()
	defer store.release()
	return store.whiteCards
}

func (store *cardMemStore) AllBlacks() []cah.BlackCard {
	store.lock()
	defer store.release()
	return store.blackCards
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
