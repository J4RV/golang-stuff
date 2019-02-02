package sqlite

import (
	"github.com/j4rv/golang-stuff/cah"
)

type cardStore struct {
}

func NewCardStore() *cardStore {
	return &cardStore{}
}

func (store *cardStore) CreateWhite(text, exp string) error {
	statement, err := db.Prepare(`INSERT INTO white_card (text, expansion) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	_, err = statement.Exec(text, exp)
	return err
}

func (store *cardStore) CreateBlack(text, exp string, blanks int) error {
	statement, err := db.Prepare(`INSERT INTO black_card (text, expansion, blanks) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = statement.Exec(text, exp, blanks)
	return err
}

func (store *cardStore) AllWhites() ([]cah.WhiteCard, error) {
	res := []cah.WhiteCard{}
	if err := db.Select(&res, "SELECT * FROM white_card"); err != nil {
		return res, err
	}
	return res, nil
}

func (store *cardStore) AllBlacks() ([]cah.BlackCard, error) {
	res := []cah.BlackCard{}
	if err := db.Select(&res, "SELECT * FROM black_card"); err != nil {
		return res, err
	}
	return res, nil
}

/*
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
