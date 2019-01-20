package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleGames(r *mux.Router) {
	s := r.PathPrefix("/game").Subrouter()
	s.Handle("/ListOpen", srvHandler(openGames)).Methods("GET")
	/*s.Handle("/Join", srvHandler(playCards)).Methods("POST")
	s.Handle("/Leave", srvHandler(playCards)).Methods("POST")*/
}

type gameResponse struct {
	ID          int      `json:"id" `
	Owner       string   `json:"owner"`
	Name        string   `json:"name" `
	HasPassword bool     `json:"hasPassword" `
	Expansions  []string `json:"expansions" `
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
			Expansions:  g.Expansions,
			//StateID:     g.StateID,
		})
	}
	writeResponse(w, response)
	return nil
}
