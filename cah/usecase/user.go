package usecase

import (
	"errors"

	"github.com/j4rv/golang-stuff/cah"
)

type userController struct {
	store cah.UserStore
}

func NewUserUsecase(store cah.UserStore) *userController {
	return &userController{store: store}
}

func (uc userController) Register(name, pass string) (cah.User, error) {
	_, ok := uc.store.ByName(name)
	if ok {
		return cah.User{}, errors.New("That username already exists. Please try another.")
	}
	return uc.store.Create(name, pass)
}

func (uc userController) ByID(id int) (cah.User, bool) {
	return uc.store.ByID(id)
}

func (uc userController) Login(name, pass string) (cah.User, bool) {
	u, err := uc.store.ByCredentials(name, pass)
	return u, err == nil
}
