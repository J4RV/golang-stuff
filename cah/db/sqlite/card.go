package sqlite

import (
	"errors"
	"fmt"
	"strings"

	"github.com/j4rv/golang-stuff/cah"
)

type cardStore struct {
}

func NewCardStore() *cardStore {
	return &cardStore{}
}

func (store *cardStore) CreateWhite(text, exp string) error {
	trimmedText, trimmedExp := strings.TrimSpace(text), strings.TrimSpace(exp)
	if trimmedText == "" || trimmedExp == "" {
		return errors.New("Text and expansion cannot be empty")
	}
	statement, err := db.Prepare(`INSERT INTO white_card (text, expansion) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	statement.Exec(trimmedText, trimmedExp)
	return nil
}

func (store *cardStore) CreateBlack(text, exp string, blanks int) error {
	trimmedText, trimmedExp := strings.TrimSpace(text), strings.TrimSpace(exp)
	if trimmedText == "" || trimmedExp == "" {
		return errors.New("Text and expansion cannot be empty")
	}
	if blanks < 1 {
		return errors.New("Blanks should be at least 1")
	}
	statement, err := db.Prepare(`INSERT INTO black_card (text, expansion, blanks) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	statement.Exec(trimmedText, trimmedExp, blanks)
	return nil
}

func (store *cardStore) AllWhites() ([]cah.WhiteCard, error) {
	rows, err := db.Query(`SELECT * FROM white_card`)
	if err != nil {
		return nil, err
	}
	res := []cah.WhiteCard{}
	for rows.Next() {
		wc := cah.WhiteCard{}
		rows.Scan(&wc.ID, &wc.Text, &wc.Expansion)
		fmt.Println("Found white card", wc)
		res = append(res, wc)
	}
	return res, nil
}

/*
func (store *cardStore) AllBlacks() []cah.BlackCard {
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

func (store *cardStore) ExpansionWhites(exps ...string) []cah.WhiteCard {
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

func (store *cardStore) ExpansionBlacks(exps ...string) []cah.BlackCard {
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

func (store *cardStore) AvailableExpansions() []string {
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

func validateCard(c cah.Card) error {
	if len(c.Text) == 0 {
		return errors.New("Card text cannot be empty")
	}
	if len(c.Text) > 120 {
		return errors.New("Card text cannot be longer than 120")
	}
	return nil
}
*/
