package data

import (
	"testing"
)

func TestPassHashNotFailing(t *testing.T) {
	hash, err := getPassHash(commonPass)
	if err != nil {
		t.Error("generate password hash failed:", err)
	} else {
		t.Log("Hash:", string(hash))
	}
}

func TestCorrectPass(t *testing.T) {
	ok := correctPass(commonPass, commonPassHash)
	if !ok {
		t.Error("password check failed")
	}
	ok = correctPass("asdwwas12354", commonPassHash)
	if ok {
		t.Error("password check should have failed")
	}
}

func TestGetUserByLogin(t *testing.T) {
	u, err := GetUserByLogin("Green", commonPass)
	if err != nil {
		t.Error(err)
	} else {
		if u.Username != "Green" {
			t.Fatal("GetUserByLogin is horribly broken")
		}
	}
	u, err = GetUserByLogin("Green", "not green's password")
	if err == nil {
		t.Error("Error should not be nil")
	}
}
