package usecase

import (
	"errors"
	"log"

	"github.com/j4rv/golang-stuff/cah"
	"golang.org/x/crypto/bcrypt"
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
	passHash, err := userPassHash(pass)
	if err != nil {
		//Never log passwords! But this one caused an error and will not be stored, its an ok exception
		log.Println("ERROR while trying to hash password.", pass, err)
		return cah.User{}, errors.New("That password could not be protected correctly. Please try another.")
	}
	return uc.store.Create(name, passHash)
}

func (uc userController) ByID(id int) (cah.User, bool) {
	return uc.store.ByID(id)
}

func (uc userController) Login(name, pass string) (cah.User, bool) {
	u, ok := uc.store.ByName(name)
	if !ok {
		return cah.User{}, false
	}
	if !userCorrectPass(pass, u.Password) {
		return cah.User{}, false
	}
	return u, true
}

// internal

const userPassCost = 10

func userPassHash(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), userPassCost)
	return string(b), err
}

func userCorrectPass(pass string, storedhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedhash), []byte(pass))
	return err == nil
}

// for testing

const commonPass = "dev"

func PopulateUsers(uuc cah.UserUsecases) {
	uuc.Register("Red", commonPass)
	uuc.Register("Green", commonPass)
	uuc.Register("Blue", commonPass)
	uuc.Register("Gold", commonPass)
	uuc.Register("Silver", commonPass)
	log.Print("Base users initialized")
}
