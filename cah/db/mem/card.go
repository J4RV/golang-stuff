package mem

import (
	"errors"
	"fmt"
	"log"

	"github.com/j4rv/golang-stuff/cah"
)

type cardMemStore struct {
	abstractMemStore
	whiteCards map[string][]cah.WhiteCard
	blackCards map[string][]cah.BlackCard
}

func NewCardStore() *cardMemStore {
	return &cardMemStore{
		whiteCards: make(map[string][]cah.WhiteCard),
		blackCards: make(map[string][]cah.BlackCard),
	}
}

func (store *cardMemStore) CreateWhite(t, e string) error {
	if len(t) == 0 {
		return errors.New("Card text cannot be empty")
	}
	if len(t) > 120 {
		return errors.New("Card text cannot be longer than 120")
	}
	store.Lock()
	defer store.Unlock()
	c := cah.WhiteCard{}
	c.ID = store.nextID()
	c.Text = t
	c.Expansion = e
	store.whiteCards[e] = append(store.whiteCards[e], c)
	return nil
}

func (store *cardMemStore) CreateBlack(t, e string, blanks int) error {
	if len(t) == 0 {
		return errors.New("Card text cannot be empty")
	}
	if len(t) > 120 {
		return errors.New("Card text cannot be longer than 120")
	}
	if blanks < 1 {
		return errors.New("Black cards need to have at least 1 blank")
	}
	if blanks > 5 {
		return fmt.Errorf("Black cards blanks maximum is five, but got %d", blanks)
	}
	store.Lock()
	defer store.Unlock()
	c := cah.BlackCard{}
	c.ID = store.nextID()
	c.Text = t
	c.Expansion = e
	c.Blanks = blanks
	store.blackCards[e] = append(store.blackCards[e], c)
	return nil
}

func (store *cardMemStore) AllWhites() []cah.WhiteCard {
	store.Lock()
	defer store.Unlock()
	ret := []cah.WhiteCard{}
	for _, whiteCards := range store.whiteCards {
		for _, whiteCard := range whiteCards {
			ret = append(ret, whiteCard)
		}
	}
	return ret
}

func (store *cardMemStore) AllBlacks() []cah.BlackCard {
	store.Lock()
	defer store.Unlock()
	ret := []cah.BlackCard{}
	for _, blackCards := range store.blackCards {
		for _, blackCard := range blackCards {
			ret = append(ret, blackCard)
		}
	}
	return ret
}

func (store *cardMemStore) ExpansionWhites(exps ...string) []cah.WhiteCard {
	store.Lock()
	defer store.Unlock()
	ret := []cah.WhiteCard{}
	for _, exp := range exps {
		cards, ok := store.whiteCards[exp]
		if !ok {
			log.Printf("Could not find white cards from expansion %s\n", exp)
			continue
		}
		ret = append(ret, cards...)
	}
	return ret
}

func (store *cardMemStore) ExpansionBlacks(exps ...string) []cah.BlackCard {
	store.Lock()
	defer store.Unlock()
	ret := []cah.BlackCard{}
	for _, exp := range exps {
		cards, ok := store.blackCards[exp]
		if !ok {
			log.Printf("Could not find black cards from expansion %s\n", exp)
			continue
		}
		ret = append(ret, cards...)
	}
	return ret
}

func (store *cardMemStore) AvailableExpansions() []string {
	store.Lock()
	defer store.Unlock()
	keys := make([]string, len(store.whiteCards))
	i := 0
	for expansion := range store.whiteCards {
		keys[i] = expansion
		i++
	}
	return keys
}
