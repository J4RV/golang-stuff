package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah"
)

func handleGames(r *mux.Router) {
	s := r.PathPrefix("/game").Subrouter()
	s.Handle("/{gameID}/room-state", srvHandler(roomState)).Methods("GET")
	s.Handle("/list-open", srvHandler(openGames)).Methods("GET")
	s.Handle("/create", srvHandler(createGame)).Methods("POST")
	s.Handle("/join", srvHandler(joinGame)).Methods("POST")
	//s.Handle("/Leave", srvHandler(playCards)).Methods("POST")
	s.Handle("/start", srvHandler(startGame)).Methods("POST")
	s.Handle("/available-expansions", srvHandler(availableExpansions)).Methods("GET")
}

/*
OPEN GAMES LIST
*/

type gameRoomResponse struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	HasPassword bool     `json:"hasPassword"`
	Players     []string `json:"players"`
	Phase       string   `json:"phase"`
	StateID     int      `json:"stateID"`
}

func roomState(w http.ResponseWriter, req *http.Request) error {
	g, err := gameFromRequest(req)
	if err != nil {
		return err
	}
	writeResponse(w, gameToResponse(g))
	return nil
}

func openGames(w http.ResponseWriter, req *http.Request) error {
	response := []gameRoomResponse{}
	for _, g := range usecase.Game.AllOpen() {
		response = append(response, gameToResponse(g))
	}
	writeResponse(w, response)
	return nil
}

func gameToResponse(g cah.Game) gameRoomResponse {
	players := make([]string, len(g.Users))
	for i := range g.Users {
		players[i] = g.Users[i].Username
	}
	return gameRoomResponse{
		ID:          g.ID,
		Owner:       g.Owner.Username,
		Name:        g.Name,
		HasPassword: g.Password != "",
		Players:     players,
		Phase:       g.State.Phase.String(),
		StateID:     g.State.ID,
	}
}

/*
CREATE GAME
*/

type createGamePayload struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}

func createGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload createGamePayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	err = usecase.Game.Create(u, payload.Name, payload.Password)
	if err != nil {
		return err
	}
	return nil
}

/*
JOIN GAME
*/

type joinGamePayload struct {
	ID int `json:"id"`
}

func joinGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload joinGamePayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	g, err := usecase.Game.ByID(payload.ID)
	if err != nil {
		return err
	}
	err = usecase.Game.UserJoins(u, g)
	if err != nil {
		return err
	}
	return nil
}

/*
JOIN GAME
*/

type startGamePayload struct {
	ID              int      `json:"id"`
	Expansions      []string `json:"expansions,omitempty"`
	HandSize        uint8    `json:"handSize,omitempty"`
	RandomFirstCzar bool     `json:"randomFirstCzar,omitempty"`
}

func startGame(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload startGamePayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	g, err := usecase.Game.ByID(payload.ID)
	if err != nil {
		return err
	}
	if g.Owner != u {
		return errors.New("Only the game owner can start the game")
	}
	state := usecase.GameState.Create()
	err = usecase.Game.Start(g, state,
		usecase.Game.Options().RandomStartingCzar(),
		usecase.Game.Options().BlackDeck(usecase.Card.ExpansionBlacks("The First Expansion")),
		usecase.Game.Options().WhiteDeck(usecase.Card.ExpansionWhites("The First Expansion")),
		usecase.Game.Options().HandSize(15),
	)
	if err != nil {
		return err
	}
	return nil
}

func availableExpansions(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	_, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	exps := usecase.Card.AvailableExpansions()
	writeResponse(w, exps)
	return nil
}

// Utils

func gameFromRequest(req *http.Request) (cah.Game, error) {
	strID := mux.Vars(req)["gameID"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		return cah.Game{}, err
	}
	g, err := usecase.Game.ByID(id)
	if err != nil {
		return g, fmt.Errorf("Could not get game with id %d", id)
	}
	return g, nil
}
