package server

import (
	"net/http"
)

func getOpenGames(w http.ResponseWriter, req *http.Request) error {
	writeResponse(w, usecase.Game.AllOpen())
	return nil
}
