package main

import (
	"github.com/j4rv/golang-stuff/cah"
	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/game"
	"github.com/j4rv/golang-stuff/cah/server"
)

func main() {
	controllers := cah.Controllers{
		game.GameController{},
	}
	dataServices := cah.DataServices{
		data.NewCardStore(),
	}
	server.Start(dataServices, controllers)
}
