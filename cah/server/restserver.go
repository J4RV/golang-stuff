package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/game"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	bd := data.GetBlackCards()
	wd := data.GetWhiteCards()
	p := game.GetRandomPlayers()
	s := game.NewGame(bd, wd, p, game.RandomStartingCzar)
	states["test"] = s
	stateRouter(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func stateRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/{stateid}").Subrouter()
	s.HandleFunc("/State", handleGetState()).Methods("GET")
	s.HandleFunc("/PutBlackCardInPlay", simpleCAHActionHandler(game.PutBlackCardInPlay)).Methods("POST")
	s.HandleFunc("/GiveBlackCardToWinner", giveBlackCardToWinner).Methods("POST")
	return r
}
func handleGetState() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		s := getState(req)
		writeJSONState(w, s)
	}
}

func writeJSONState(w http.ResponseWriter, s game.State) {
	j, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", j)
	}
}
