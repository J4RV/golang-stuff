package usecase

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/j4rv/golang-stuff/cah"
)

type cardController struct {
	store cah.CardStore
}

func NewCardUsecase(store cah.CardStore) *cardController {
	return &cardController{store: store}
}

func (cc cardController) AllBlacks() []cah.BlackCard {
	return cc.store.AllBlacks()
}

func (cc cardController) AllWhites() []cah.WhiteCard {
	return cc.store.AllWhites()
}

func (cc cardController) ExpansionWhites(exps ...string) []cah.WhiteCard {
	return cc.store.ExpansionWhites(exps...)
}

func (cc cardController) ExpansionBlacks(exps ...string) []cah.BlackCard {
	return cc.store.ExpansionBlacks(exps...)
}

func (cc cardController) AvailableExpansions() []string {
	return cc.store.AvailableExpansions()
}

// CreateFromReaders creates and stores cards from two readers.
// The reader should provide a card per line. A line can contain "\n"s for card line breaks.
// Lines containing only whitespace are ignored
func (cc cardController) CreateFromReaders(wdat, bdat io.Reader, expansionName string) error {
	// Create cards from files
	var err error
	err = doEveryLine(wdat, func(t string) {
		if strings.TrimSpace(t) == "" {
			return
		}
		cc.store.CreateWhite(t, expansionName)
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
		cc.store.CreateBlack(t, expansionName, blanks)
	})
	log.Println("Successfully loaded cards from expansion " + expansionName)
	return err
}

// CreateFromFolder creates and stores cards from an expansion folder
// That folder should contain two files called 'white.md' and 'black.md'
// The files content is treated as explained for the CreateCards function
func (cc cardController) CreateFromFolder(folderPath, expansionName string) error {
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
