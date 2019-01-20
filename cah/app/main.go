package main

import (
	"math/rand"
	"time"

	"github.com/j4rv/golang-stuff/cah"
	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/server"
	"github.com/j4rv/golang-stuff/cah/usecase"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	run()
}

func run() {
	gameStore := data.NewGameStore()
	cardStore := data.NewCardStore()
	userStore := data.NewUserStore()
	data.PopulateUsers(userStore)
	usecases := cah.Usecases{
		Game: usecase.NewGameUsecase(gameStore),
		Card: usecase.NewCardUsecase(cardStore),
		User: usecase.NewUserUsecase(userStore),
	}
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
