package data

import (
	"golang.org/x/crypto/bcrypt"
)

func getPassHash(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(b), err
}

func correctPass(pass string, storedhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedhash), []byte(pass))
	return err == nil
}
