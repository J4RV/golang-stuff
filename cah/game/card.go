package game

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/j4rv/golang-stuff/cah"
)

type CardController struct {
	Store cah.CardStore
}

func (cc CardController) AllBlacks() []cah.BlackCard {
	return cc.Store.AllBlacks()
}

func (cc CardController) AllWhites() []cah.WhiteCard {
	return cc.Store.AllWhites()
}

// CreateFromReaders creates and stores cards from two readers.
// The reader should provide a card per line. A line can contain "\n"s for card line breaks.
// Lines containing only whitespace are ignored
func (cc CardController) CreateFromReaders(wdat, bdat io.Reader, expansionName string) error {
	// Create cards from files
	var err error
	err = doEveryLine(wdat, func(t string) {
		if strings.TrimSpace(t) == "" {
			return
		}
		cc.Store.CreateWhite(t, expansionName)
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
		cc.Store.CreateBlack(t, expansionName, blanks)
	})
	log.Println("Successfully loaded cards from expansion " + expansionName)
	return err
}

// CreateFromFolder creates and stores cards from an expansion folder
// That folder should contain two files called 'white.md' and 'black.md'
// The files content is treated as explained for the CreateCards function
func (cc CardController) CreateFromFolder(folderPath, expansionName string) error {
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
	return cc.CreateFromReaders(wdat, bdat, expansionName)
}

func doEveryLine(r io.Reader, fun func(string)) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := strings.Replace(s.Text(), "\\n", "\n", -1)
		fun(t)
	}
	return s.Err()
}