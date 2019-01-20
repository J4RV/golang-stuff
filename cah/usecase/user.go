package usecase

import (
	"errors"

	"github.com/j4rv/golang-stuff/cah"
)

type UserController struct {
	Store cah.UserStore
}

func (uc UserController) Register(name, pass string) (cah.User, error) {
	_, ok := uc.Store.ByName(name)
	if ok {
		return cah.User{}, errors.New("That username already exists. Please try another.")
	}
	return uc.Store.Create(name, pass)
}

func (uc UserController) ByID(id int) (cah.User, bool) {
	return uc.Store.ByID(id)
}

func (uc UserController) Login(name, pass string) (cah.User, bool) {
	u, err := uc.Store.ByCredentials(name, pass)
	return u, err == nil
}
