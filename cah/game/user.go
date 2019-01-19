package game

import (
	"github.com/j4rv/golang-stuff/cah"
)

type UserController struct {
	Store cah.UserStore
}

func (uc UserController) ByID(id int) (cah.User, error) {
	return uc.Store.ByID(id)
}

func (uc UserController) Login(name, pass string) (cah.User, error) {
	return uc.Store.ByCredentials(name, pass)
}
