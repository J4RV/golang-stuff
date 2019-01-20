package usecase

import (
	"github.com/j4rv/golang-stuff/cah"
	"golang.org/x/crypto/bcrypt"
)

type gameController struct {
	store cah.GameStore
}

func (control gameController) AllOpen() []cah.Game {
	return control.store.ByStatePhase(cah.NotStarted)
}

// crypto

const gamePassCost = 4

func gamePassHash(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), gamePassCost)
	return string(b), err
}

func gameCorrectPass(pass string, storedhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedhash), []byte(pass))
	return err == nil
}
