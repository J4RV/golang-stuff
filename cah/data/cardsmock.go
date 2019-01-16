package data

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func doEveryLine(r io.Reader, fun func(string)) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := strings.Replace(s.Text(), "\\n", "\n", -1)
		fun(t)
	}
	return s.Err()
}

func LoadCards(expansionPath, expansionName string) error {
	wd := &whiteCards
	bd := &blackCards
	// Open both files
	wdat, err := os.Open(fmt.Sprintf("%s/white.md", expansionPath))
	if err != nil {
		return err
	}
	defer wdat.Close()
	bdat, err := os.Open(fmt.Sprintf("%s/black.md", expansionPath))
	if err != nil {
		return err
	}
	defer bdat.Close()
	// Create cards from files
	err = doEveryLine(wdat, func(t string) {
		*wd = append(*wd, WhiteCard{Card: Card{Text: t, Expansion: expansionName}})
	})
	if err != nil {
		return err
	}
	err = doEveryLine(bdat, func(t string) {
		blanks := strings.Count(t, "_")
		if blanks == 0 {
			blanks = 1
		}
		*bd = append(*bd, BlackCard{Card: Card{Text: t, Expansion: expansionName}, BlanksAmount: blanks})
	})
	log.Println("Successfully loaded cards from expansion " + expansionName)
	return err
}
