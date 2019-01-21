package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j4rv/golang-stuff/cah/db/mem"
	"github.com/j4rv/golang-stuff/cah/usecase/fixture"

	"github.com/j4rv/golang-stuff/cah"
)

var commonPass = "dev"
var commonPassHash, _ = userPassHash(commonPass)

var usecase cah.UserUsecases

func init() {
	store := mem.NewUserStore()
	usecase = NewUserUsecase(store)
	fixture.PopulateUsers(usecase)
}

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

func TestUserByID(t *testing.T) {
	var table = []struct {
		id    int
		name  string
		found bool
	}{
		{-1, "", false},
		{0, "", false},
		{1, "Red", true},
		{3, "Blue", true},
		{999, "", false},
	}
	for _, row := range table {
		u, ok := usecase.ByID(row.id)
		assert.Equal(t, row.found, ok)
		if ok {
			assert.Equal(t, row.name, u.Username)
		}
	}
}

func TestGetUserByLogin(t *testing.T) {
	u, ok := usecase.Login("Green", "Green")
	if !ok {
		t.Error("Could not login as Green")
	} else {
		if u.Username != "Green" {
			t.Fatal("GetUserByLogin is horribly broken")
		}
	}
	u, ok = usecase.Login("Green", "not green's password")
	if ok {
		t.Error("Error should not be nil")
	}
}
