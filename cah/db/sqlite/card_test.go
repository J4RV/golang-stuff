package sqlite

import (
	"os"
	"testing"
)

func setup(t *testing.T) (*cardStore, func()) {
	initDB("cah_db_card.test.db")
	return NewCardStore(), func() {
		db.Close()
		err := os.Remove("cah_db_card.test.db")
		if err != nil {
			panic(err)
		}
	}
}

func TestCardStore(t *testing.T) {
	cases := []struct {
		name            string
		text, expansion string
		errExpected     bool
	}{
		{"first tests", "A big D", "tests", false},
		{"second tests", "A white guy dabbing", "tests", false},
		{"empty text", "", "tests", true},
		{"empty expansion", "Homeless card", "", true},
	}
	cs, teardown := setup(t)
	defer teardown()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := cs.CreateWhite(tc.text, tc.expansion)
			if !tc.errExpected && err != nil {
				t.Fatal(err.Error())
			}
			if tc.errExpected && err == nil {
				t.Fatal("Expected error but found nil")
			}
		})
	}
}
