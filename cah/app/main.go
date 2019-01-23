package main

import (
	"math/rand"
	"time"

	"github.com/j4rv/golang-stuff/cah"
	db "github.com/j4rv/golang-stuff/cah/db/mem"
	"github.com/j4rv/golang-stuff/cah/server"
	"github.com/j4rv/golang-stuff/cah/usecase"
	"github.com/j4rv/golang-stuff/cah/usecase/fixture"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	run()
}

func run() {
	stateStore := db.NewGameStateStore()
	gameStore := db.NewGameStore(stateStore)
	cardStore := db.NewCardStore()
	userStore := db.NewUserStore()
	usecases := cah.Usecases{
		GameState: usecase.NewGameStateUsecase(stateStore),
		Card:      usecase.NewCardUsecase(cardStore),
		User:      usecase.NewUserUsecase(userStore),
		Game:      usecase.NewGameUsecase(gameStore),
	}
	fixture.PopulateUsers(usecases.User)
	createTestGames(usecases)
	populateCards(usecases.Card)
	server.Start(usecases)
}

// For quick prototyping

func createTestGames(usecase cah.Usecases) {
	users := getTestUsers(usecase)
	usecase.Game.Create(users[1], "A long and descriptive game name", "")
	usecase.Game.Create(users[0], "Amo a juga", "1234")
}

func getTestUsers(usecase cah.Usecases) []cah.User {
	users := make([]cah.User, 4)
	for i := 0; i < 4; i++ {
		u, _ := usecase.User.ByID(i + 1)
		users[i] = u
	}
	return users
}

func populateCards(cardUC cah.CardUsecases) {
	cardUC.CreateFromFolder("./expansions/base-uk", "Base-UK")
	cardUC.CreateFromFolder("./expansions/anime", "Anime")
	cardUC.CreateFromFolder("./expansions/kikis", "Kikis")
	cardUC.CreateFromFolder("./expansions/expansion-1", "The First Expansion")
	cardUC.CreateFromFolder("./expansions/expansion-2", "The Second Expansion")
	// to check that it does not break the app
	cardUC.CreateFromFolder("./expansinos/undefined", "Non existant")
}
