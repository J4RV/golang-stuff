package unocards

import "testing"

const unodecksize = 108

func TestNewVanillaSize(t *testing.T) {
	deck := NewVanilla()
	if len(deck) != unodecksize {
		t.Error("Unexpected size of vanilla UNO deck:", len(deck),
			", expected: ", unodecksize)
	}
}

func TestDeck_Filter(t *testing.T) {
	noreds := func(c Card) bool {
		return c.Color != Red
	}
	nofives := func(c Card) bool {
		return c.Value != Five
	}
	filtered := NewVanilla(Filter(noreds), Filter(nofives))
	for _, c := range filtered {
		if c.Color == Red {
			t.Error("Found a 'Red' card and expected none, card:", c)
		}
		if c.Value == Five {
			t.Error("Found a 'Five' card and expected none, card:", c)
		}
	}
}

func TestDeck_Shuffle(t *testing.T) {
	ordered := NewVanilla()
	shuffled := NewVanilla(Shuffle)

	if len(ordered) != len(shuffled) {
		t.Error("Shuffle function altered the deck size!")
		t.FailNow()
	}

	matching := 0.0
	for i := range ordered {
		o := ordered[i]
		s := shuffled[i]
		if o == s {
			matching++
		}
	}

	// If a certain percentage of the deck is the same, something *might* be wrong
	similarityThreshold := float64(len(ordered)) * 0.2
	if matching > similarityThreshold {
		similarity := matching / float64(len(ordered))
		t.Error("Warning, the shuffled deck is too similar to the ordered deck, similarity:", similarity)
	}

	if int(matching) == len(ordered) {
		t.Error("The shuffled deck is EQUAL to the ordered deck, this should happen once in a lifetime")
	}
}
