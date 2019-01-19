package main

import (
	"math/rand"
	"time"

	"github.com/j4rv/golang-stuff/cah"
	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/game"
	"github.com/j4rv/golang-stuff/cah/server"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	run()
}

func run() {
	cardStore := data.NewCardStore()
	userStore := data.NewUserStore()
	data.PopulateUsers(&userStore)
	usecases := cah.Usecases{
		game.GameController{
			//Store: gameStore,
		},
		game.CardController{
			Store: cardStore,
		},
		game.UserController{
			Store: userStore,
		},
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