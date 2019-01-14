package data

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPassHashNotFailing(t *testing.T) {
	hash, err := getPassHash(commonPass, commonSalt)
	if err != nil {
		t.Error("generate password hash failed:", err)
	} else {
		t.Log("Hash:", string(hash))
	}
}

func TestPassHash(t *testing.T) {
	salted := append([]byte(commonPass), commonSalt...)
	err := bcrypt.CompareHashAndPassword(commonPassHash, []byte(salted))
	if err != nil {
		t.Error("password hash comparison failed:", err)
	}
}

func TestCardsInit(t *testing.T) {
	if len(whiteCards) == 0 {
		t.Error("white cards were not initialized correctly")
	} else {
		t.Log("whiteCards:\n", whiteCards)
	}
	if len(blackCards) == 0 {
		t.Error("black cards were not initialized correctly")
	} else {
		t.Log("blackCards:\n", blackCards)
	}
}
