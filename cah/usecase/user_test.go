package usecase

import "testing"

var commonPass = "dev"
var commonPassHash, _ = userPassHash(commonPass)

func TestPassHashNotFailing(t *testing.T) {
	hash, err := userPassHash(commonPass)
	if err != nil {
		t.Error("generate password hash failed:", err)
	} else {
		t.Log("Hash:", string(hash))
	}
}

func TestCorrectPass(t *testing.T) {
	ok := userCorrectPass(commonPass, commonPassHash)
	if !ok {
		t.Error("password check failed")
	}
	ok = userCorrectPass("asdwwas12354", commonPassHash)
	if ok {
		t.Error("password check should have failed")
	}
}

/*
Needs store mock, dont want to do it right now tbh
var usecase = NewUserUsecase()

// this one also testes the ByName method
func TestGetUserByLogin(t *testing.T) {
	u, err := usecase.Login("Green", commonPass)
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
}*/
