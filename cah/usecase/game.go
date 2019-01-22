package usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/j4rv/golang-stuff/cah"
)

type gameController struct {
	store cah.GameStore
	users cah.UserUsecases
}

func NewGameUsecase(store cah.GameStore, uuc cah.UserUsecases) *gameController {
	return &gameController{
		store: store,
		users: uuc,
	}
}

func (control gameController) Create(owner cah.User, name, pass string) error {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return errors.New("A game name cannot be blank")
	}
	owner, ok := control.users.ByID(owner.ID)
	if !ok {
		return fmt.Errorf("No user find with owner ID %d", owner.ID)
	}
	game := cah.Game{
		Owner:  owner,
		UserID: owner.ID,
		Name:   trimmed,
		Users:  []cah.User{owner},
	}
	trimmedPass := strings.TrimSpace(pass)
	if trimmedPass != "" {
		game.Password = trimmedPass
	}
	return control.store.Create(game)
}

func (control gameController) ByID(id int) (cah.Game, error) {
	return control.store.ByID(id)
}

func (control gameController) AllOpen() []cah.Game {
	return control.store.ByStatePhase(cah.NotStarted)
}

func (control gameController) UserJoins(user cah.User, game cah.Game) error {
	for _, u := range game.Users {
		if u.ID == user.ID {
			return nil // don't add the user if they already joined
		}
	}
	game.Users = append(game.Users, user)
	return control.store.Update(game)
}
