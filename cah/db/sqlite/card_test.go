package sqlite

import (
	"testing"
)

func cardTestSetup(t *testing.T) (*cardStore, func()) {
	InitDB(":memory:")
	CreateTables()
	return NewCardStore(), func() {
		db.Close()
	}
}

func TestWhiteCardCreate(t *testing.T) {
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
	cs, teardown := cardTestSetup(t)
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

func TestBlackCardCreate(t *testing.T) {
	cases := []struct {
		name            string
		text, expansion string
		blanks          int
		errExpected     bool
	}{
		{"first tests", "What's the worst thing about 9/11", "tests", 1, false},
		{"second tests", "_ and _ should never appear together in a single phrase.", "tests", 2, false},
		{"empty text", "", "tests", 1, true},
		{"empty expansion", "Homeless card", "", 1, true},
		{"zero blanks amount", "Lorem", "Ipsum", 0, true},
		{"negative blanks amount", "Lorem", "Ipsum", -1, true},
	}
	cs, teardown := cardTestSetup(t)
	defer teardown()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := cs.CreateBlack(tc.text, tc.expansion, tc.blanks)
			if !tc.errExpected && err != nil {
				t.Fatal(err.Error())
			}
			if tc.errExpected && err == nil {
				t.Fatal("Expected error but found nil")
			}
		})
	}
}

func TestBlackCardGetAll(t *testing.T) {
	cs, teardown := cardTestSetup(t)
	defer teardown()
	db.MustExec(`INSERT INTO black_card (text, expansion, blanks) VALUES
		("Text 1", "tests", 1),
		("Text 2: _ and _", "tests", 2),
		("Text 3: _, _ or _", "tests", 3)
	;`)
	all, err := cs.AllBlacks()
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(all) != 3 {
		t.Fatalf("Expected amount of 3 black cards, but got %d", len(all))
	}
	if all[0].Text != "Text 1" {
		t.Error("First returned card text is not what was expected")
	}
	if all[0].Expansion != "tests" {
		t.Error("First returned card expansion is not what was expected")
	}
	if all[2].Blanks != 3 {
		t.Error("Third returned card blanks is not what was expected")
	}
}
