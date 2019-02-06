package sqlite

import (
	"github.com/j4rv/golang-stuff/cah"
	"github.com/jmoiron/sqlx"
)

type cardStore struct {
}

func NewCardStore() *cardStore {
	return &cardStore{}
}

func (store *cardStore) CreateWhite(text, exp string) error {
	_, err := db.Exec(`INSERT INTO white_card (text, expansion) VALUES (?, ?)`, text, exp)
	return err
}

func (store *cardStore) CreateBlack(text, exp string, blanks int) error {
	_, err := db.Exec(`INSERT INTO black_card (text, expansion, blanks) VALUES (?, ?, ?)`, text, exp, blanks)
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

func (store *cardStore) ExpansionWhites(exps ...string) ([]cah.WhiteCard, error) {
	res := []cah.WhiteCard{}
	query, args, err := sqlx.In("SELECT * FROM white_card WHERE expansion IN (?)", exps)
	if err != nil {
		return res, err
	}
	if err = db.Select(&res, query, args...); err != nil {
		return res, err
	}
	return res, nil
}

func (store *cardStore) ExpansionBlacks(exps ...string) ([]cah.BlackCard, error) {
	res := []cah.BlackCard{}
	query, args, err := sqlx.In("SELECT * FROM black_card WHERE expansion IN (?)", exps)
	if err != nil {
		return res, err
	}
	if err = db.Select(&res, query, args...); err != nil {
		return res, err
	}
	return res, nil
}

func (store *cardStore) AvailableExpansions() ([]string, error) {
	res := []string{}
	if err := db.Select(&res, "SELECT DISTINCT expansion FROM white_card "); err != nil {
		return res, err
	}
	return res, nil
}
