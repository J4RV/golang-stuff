package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/game"

	"github.com/gorilla/mux"
)

var g game.State

func main() {
	router := mux.NewRouter()
	bd := data.GetBlackCards()
	wd := data.GetWhiteCards()
	p := game.GetRandomPlayers()
	g = game.NewGame(bd, wd, p, game.RandomStartingCzar)
	g, _ = game.PutBlackCardInPlay(g)
	router.HandleFunc("/blackcard", handleGetBlackCard(g)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func handleGetBlackCard(state game.State) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		bc := state.BlackCardInPlay
		fmt.Fprintf(w, "%s", bc.GetText())
		g, _ = game.PutBlackCardInPlay(state)
	}
}
