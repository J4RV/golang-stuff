package usecase

import (
	"errors"
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
	"golang.org/x/crypto/bcrypt"
)

type gameController struct {
	store cah.GameStore
	users cah.UserStore
}

func NewGameUsecase(store cah.GameStore, us cah.UserStore) *gameController {
	return &gameController{
		store: store,
		users: us,
	}
}

func (control gameController) Create(owner cah.User, name, pass string, expansions []string, state cah.GameState) error {
	if state.Phase != cah.NotStarted {
		return errors.New("Cannot create a new game with an already started State")
	}
	owner, ok := control.users.ByID(owner.ID)
	if !ok {
		return fmt.Errorf("No user find with owner ID %d", owner.ID)
	}
	game := cah.Game{
		OwnerID:    owner.ID,
		Name:       name,
		Expansions: expansions,
		StateID:    state.ID,
	}
	if pass != "" {
		hashed, err := gamePassHash(pass)
		if err != nil {
			return err
		}
		game.Password = hashed
	}
	return control.store.Create(game)
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
