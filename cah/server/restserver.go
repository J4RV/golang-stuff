package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/game"

	"github.com/gorilla/mux"
)

func main() {
	var dir string
	port := 8000

	flag.StringVar(&dir, "dir", "./public/react/build", "the directory to serve files from. Defaults to './public'")
	flag.Parse()

	router := mux.NewRouter()
	bd := data.GetBlackCards()
	wd := data.GetWhiteCards()
	p := game.GetRandomPlayers()
	s := game.NewGame(bd, wd, p, game.RandomStartingCzar)
	states["test"] = s

	stateRouter(router)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Printf("Starting server in port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func stateRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/rest/{gameid}/{playerid}").Subrouter()
	s.HandleFunc("/State", handleGetState()).Methods("GET")
	s.HandleFunc("/PutBlackCardInPlay", simpleCAHActionHandler(game.PutBlackCardInPlay)).Methods("POST")
	s.Handle("/GiveBlackCardToWinner", appHandler(giveBlackCardToWinner)).Methods("POST")
	s.Handle("/PlayCards", appHandler(playCards)).Methods("POST")
	return r
}

func cleanStateForPlayer(s *game.State, p game.Player) {
	//todo: replace cards in other player hands with "Unknown" cards to prevent cheating
}

func handleGetState() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		s := getState(req)
		p, err := getPlayer(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		}
		cleanStateForPlayer(&s, p)
		writeJSONState(w, s)
	}
}

func writeJSONState(w http.ResponseWriter, s game.State) {
	j, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", j)
	}
}
