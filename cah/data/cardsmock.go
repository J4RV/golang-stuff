package data

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func initCards() {
	err := loadCards("base-uk", &whiteCards, &blackCards)
	if err != nil {
		log.Fatal(err)
	}
}

func doEveryLine(f *os.File, fun func(string)) error {
	s := bufio.NewScanner(f)
	for s.Scan() {
		t := strings.Replace(s.Text(), "\\n", "\n", -1)
		fun(t)
	}
	return s.Err()
}

func loadCards(expansion string, wd *[]WhiteCard, bd *[]BlackCard) error {
	wdat, err := os.Open(fmt.Sprintf("./expansions/%s/white.md", expansion))
	if err != nil {
		return err
	}
	defer wdat.Close()

	bdat, err := os.Open(fmt.Sprintf("./expansions/%s/black.md", expansion))
	if err != nil {
		return err
	}
	defer bdat.Close()

	err = doEveryLine(wdat, func(t string) {
		*wd = append(*wd, WhiteCard{Card: Card{Text: t, Expansion: expansion}})
	})
	if err != nil {
		return err
	}
	err = doEveryLine(bdat, func(t string) {
		blanks := strings.Count(t, "_")
		if blanks == 0 {
			blanks = 1
		}
		*bd = append(*bd, BlackCard{Card: Card{Text: t, Expansion: expansion}, BlanksAmount: blanks})
	})
	log.Println("Successfully loaded cards from expansion " + expansion)
	return err
}
