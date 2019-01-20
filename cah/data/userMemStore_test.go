package data

import (
	"testing"
)

var store = NewUserStore()
var commonPassHash, _ = getPassHash(commonPass)

func init() {
	PopulateUsers(store)
}

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

func TestGetUserByID(t *testing.T) {
	u, ok := store.ByID(1)
	if !ok {
		t.Error("Did not find user with id 1")
	} else {
		if u.Username != "Red" && u.ID == 1 {
			t.Fatal("GetUserByID is horribly broken")
		}
	}
	u, ok = store.ByID(999)
	if ok {
		t.Error("Found an user with id 999 but expected none")
	}
}

// this one also testes the ByName method
func TestGetUserByLogin(t *testing.T) {
	u, err := store.ByCredentials("Green", commonPass)
	if err != nil {
		t.Error(err)
	} else {
		if u.Username != "Green" {
			t.Fatal("GetUserByLogin is horribly broken")
		}
	}
	u, err = store.ByCredentials("Green", "not green's password")
	if err == nil {
		t.Error("Error should not be nil")
	}
}
