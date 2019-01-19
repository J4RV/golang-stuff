package data

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/j4rv/golang-stuff/cah"
)

type CardMemStore struct {
	whiteCards []cah.WhiteCard
	blackCards []cah.BlackCard
}

func NewCardStore() *CardMemStore {
	return &CardMemStore{
		whiteCards: []cah.WhiteCard{},
		blackCards: []cah.BlackCard{},
	}
}

func (s *CardMemStore) CreateWhite(t, e string) error {
	c := cah.WhiteCard{}
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
	c := cah.BlackCard{}
	c.Text = t
	c.Expansion = e
	c.BlanksAmount = blanks
	err := validateCard(c.Card)
	if err == nil {
		s.blackCards = append(s.blackCards, c)
	}
	return err
}

func (s *CardMemStore) GetWhites() []cah.WhiteCard {
	return s.whiteCards
}

func (s *CardMemStore) GetBlacks() []cah.BlackCard {
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

// CreateCards creates and stores cards from two readers.
// The reader should provide a card per line. A line can contain "\n"s for card line breaks.
// Lines containing only whitespace are ignored
func (rep *CardMemStore) CreateCards(wdat, bdat io.Reader, expansionName string) error {
	// Create cards from files
	var err error
	err = doEveryLine(wdat, func(t string) {
		if strings.TrimSpace(t) == "" {
			return
		}
		rep.CreateWhite(t, expansionName)
	})
	if err != nil {
		return err
	}
	err = doEveryLine(bdat, func(t string) {
		if strings.TrimSpace(t) == "" {
			return
		}
		blanks := strings.Count(t, "_")
		if blanks == 0 {
			blanks = 1
		}
		rep.CreateBlack(t, expansionName, blanks)
	})
	log.Println("Successfully loaded cards from expansion " + expansionName)
	return err
}

// CreateCardsFromFolder creates and stores cards from an expansion folder
// That folder should contain two files called 'white.md' and 'black.md'
// The files content is treated as explained for the CreateCards function
func (rep *CardMemStore) CreateCardsFromFolder(folderPath, expansionName string) error {
	wdat, err := os.Open(fmt.Sprintf("%s/white.md", folderPath))
	defer wdat.Close()
	if err != nil {
		return err
	}
	bdat, err := os.Open(fmt.Sprintf("%s/black.md", folderPath))
	defer bdat.Close()
	if err != nil {
		return err
	}
	return rep.CreateCards(wdat, bdat, expansionName)
}
