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
	s.Handle("/ListOpen", srvHandler(openGames)).Methods("GET")
	s.Handle("/Create", srvHandler(createGame)).Methods("POST")
	/*s.Handle("/Join", srvHandler(playCards)).Methods("POST")
	s.Handle("/Leave", srvHandler(playCards)).Methods("POST")*/
}

/*
OPEN GAMES LIST
*/

type gameResponse struct {
	ID          int    `json:"id" `
	Owner       string `json:"owner"`
	Name        string `json:"name" `
	HasPassword bool   `json:"hasPassword" `
	//StateID     int      `json:"stateID"`
}

func openGames(w http.ResponseWriter, req *http.Request) error {
	response := []gameResponse{}
	for _, g := range usecase.Game.AllOpen() {
		owner, ok := usecase.User.ByID(g.OwnerID)
		if !ok {
			log.Printf("Game '%s' owner with ID %d not found", g.Name, g.OwnerID)
			continue
		}
		response = append(response, gameResponse{
			ID:          g.ID,
			Owner:       owner.Username,
			Name:        g.Name,
			HasPassword: g.Password != "",
			//StateID:     g.StateID,
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
