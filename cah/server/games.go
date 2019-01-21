package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleGames(r *mux.Router) {
	s := r.PathPrefix("/game").Subrouter()
	s.Handle("/list-open", srvHandler(openGames)).Methods("GET")
	s.Handle("/create", srvHandler(createGame)).Methods("POST")
	s.Handle("/join", srvHandler(joinGame)).Methods("POST")
	//s.Handle("/Leave", srvHandler(playCards)).Methods("POST")
}

/*
OPEN GAMES LIST
*/

type openGameResponse struct {
	ID          int      `json:"id"`
	Owner       string   `json:"owner"`
	Name        string   `json:"name"`
	HasPassword bool     `json:"hasPassword"`
	Players     []string `json:"players"`
}

func openGames(w http.ResponseWriter, req *http.Request) error {
	response := []openGameResponse{}
	for _, g := range usecase.Game.AllOpen() {
		owner, ok := usecase.User.ByID(g.OwnerID)
		if !ok {
			log.Printf("Game '%s' owner with ID %d not found", g.Name, g.OwnerID)
			continue
		}
		players := make([]string, len(g.Users))
		for i := range g.Users {
			players[i] = g.Users[i].Username
		}
		response = append(response, openGameResponse{
			ID:          g.ID,
			Owner:       owner.Username,
			Name:        g.Name,
			HasPassword: g.Password != "",
			Players:     players,
		})
	}
	writeResponse(w, response)
	return nil
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
