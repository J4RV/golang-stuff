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
	createTestGame()
	stateRouter(router)
	router.HandleFunc("/rest/login", processLogin).Methods("POST")
	router.HandleFunc("/rest/validcookie", validCookie).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Printf("Starting server in port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func createTestGame() {
	bd := data.GetBlackCards()
	wd := data.GetWhiteCards()
	p := game.GetRandomPlayers()
	s := game.NewGame(bd, wd, p, game.RandomStartingCzar)
	sg := serverGame{state: s}
	sg.userToPlayers = make(map[data.User]*game.Player)
	for i, p := range p {
		user, _ := data.GetUserById(i)
		sg.userToPlayers[user] = p
	}
	games["test"] = sg
}

func stateRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/rest/{gameid}").Subrouter()
	s.HandleFunc("/State", handleGetState()).Methods("GET")
	s.Handle("/GiveBlackCardToWinner", appHandler(giveBlackCardToWinner)).Methods("POST")
	s.Handle("/PlayCards", appHandler(playCards)).Methods("POST")
	return r
}

func handleGetState() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		game, err := getGame(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		writeJSONState(w, game.state)
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
