package data

import "golang.org/x/crypto/bcrypt"

func getPassHash(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(b), err
}

func correctPass(pass string, storedhash string) bool {
	hp, err := getPassHash(pass)
	if err != nil {
		return false
	}
	return nil != bcrypt.CompareHashAndPassword([]byte(hp), []byte(storedhash))
}
