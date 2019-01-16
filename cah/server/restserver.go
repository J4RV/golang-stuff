package main

import (
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

	data.LoadCards("./expansions/base-uk", "Base UK")
	flag.StringVar(&dir, "dir", "./public/react/build", "the directory to serve files from. Defaults to './public'")
	flag.Parse()

	router := mux.NewRouter()
	createTestGame()
	stateRouter(router)
	router.HandleFunc("/user/login", processLogin).Methods("POST")
	router.HandleFunc("/user/logout", processLogout).Methods("POST", "GET")
	router.HandleFunc("/user/validcookie", validCookie).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Printf("Starting server in port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func createTestGame() {
	bd := data.GetBlackCards()
	wd := data.GetWhiteCards()
	p := getTestPlayers()
	s := game.NewGame(bd, wd, p, game.RandomStartingCzar)
	sg := serverGame{state: s}
	sg.userToPlayers = make(map[data.User]*game.Player)
	for i, p := range p {
		user, _ := data.GetUserById(i)
		sg.userToPlayers[user] = p
	}
	games["test"] = sg
}

func getTestPlayers() []*game.Player {
	users := make([]data.User, 3)
	for i := 0; i < 3; i++ {
		u, _ := data.GetUserById(i)
		users[i] = u
	}
	return getPlayersForUsers(users...)
}

func getPlayersForUsers(users ...data.User) []*game.Player {
	ret := make([]*game.Player, len(users))
	for i, u := range users {
		ret[i] = game.NewPlayer(u.ID, u.Username)
	}
	return ret
}

func stateRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/rest/{gameid}").Subrouter()
	s.Handle("/State", appHandler(getGameStateForUser)).Methods("GET")
	s.Handle("/GiveBlackCardToWinner", appHandler(giveBlackCardToWinner)).Methods("POST")
	s.Handle("/PlayCards", appHandler(playCards)).Methods("POST")
	return r
}

func getGameStateForUser(w http.ResponseWriter, req *http.Request) error {
	u, err := userFromSession(req)
	if err != nil {
		return err
	}
	sg, err := getGame(req)
	if err != nil {
		return err
	}
	p, err := getPlayer(sg, u)
	if err != nil {
		return err
	}
	state := sg.state
	response := gameStateResponse{
		Phase:           int(state.Phase),
		Players:         getPlayerInfo(sg),
		CurrCzarID:      state.Players[state.CurrCzarIndex].ID,
		BlackCardInPlay: state.BlackCardInPlay,
		SinnerPlays:     getSinnerPlays(sg),
		DiscardPile:     state.DiscardPile,
		MyPlayer:        *p,
	}
	writeResponse(w, response)
	return nil
}

func getPlayerInfo(sg serverGame) []playerInfo {
	ret := make([]playerInfo, len(sg.state.Players))
	for i, p := range sg.state.Players {
		ret[i] = gamePlayerToPlayerInfo(*p)
	}
	return ret
}

func gamePlayerToPlayerInfo(p game.Player) playerInfo {
	return playerInfo{
		ID:               p.ID,
		Name:             p.Name,
		HandSize:         len(p.Hand),
		WhiteCardsInPlay: len(p.WhiteCardsInPlay),
		Points:           p.Points,
	}
}

func getSinnerPlays(sg serverGame) []sinnerPlay {
	if !game.AllSinnersPlayedTheirCards(sg.state) {
		return []sinnerPlay{}
	}
	ret := make([]sinnerPlay, len(sg.state.Players))
	for i, p := range sg.state.Players {
		ret[i] = sinnerPlay{
			ID:         p.ID,
			WhiteCards: p.WhiteCardsInPlay,
		}
	}
	return ret
}
