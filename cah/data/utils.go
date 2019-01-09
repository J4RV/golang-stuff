package data

import "golang.org/x/crypto/bcrypt"

func getPassHash(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), 10)
}

func correctPass(pass string, storedhash string) bool {
	hp, err := getPassHash(pass)
	if err != nil {
		return false
	}
	return nil != bcrypt.CompareHashAndPassword(hp, []byte(storedhash))
}
