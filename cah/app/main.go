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
	}
	gameUsecases := usecase.NewGameUsecase(gameStore, usecases.User)
	usecases.Game = gameUsecases
	fixture.PopulateUsers(usecases.User)
	populateCards(usecases.Card)
	server.Start(usecases)
}

func populateCards(cardUC cah.CardUsecases) {
	cardUC.CreateFromFolder("./expansions/base-uk", "Base-UK")
	cardUC.CreateFromFolder("./expansions/anime", "Anime")
	cardUC.CreateFromFolder("./expansions/kikis", "Kikis")
	cardUC.CreateFromFolder("./expansions/expansion-1", "The First Expansion")
	cardUC.CreateFromFolder("./expansions/expansion-2", "The Second Expansion")
}
