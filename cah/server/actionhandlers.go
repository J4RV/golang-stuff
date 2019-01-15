package main

import (
	"log"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/data"
	"github.com/j4rv/golang-stuff/cah/game"
)

type appHandler func(http.ResponseWriter, *http.Request) error

type gameHandler struct {
	loggedHandler
	game   serverGame
	action func(w http.ResponseWriter, r *http.Request, u data.User, s game.State) error
}

func (gh gameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.Printf("ServeHTTP error: %s", err)
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
	}
}

func simpleCAHActionHandler(f func(game.State) (game.State, error)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		g, err := getGame(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		newState, err := f(g.state)
		if err != nil {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			updateGameState(req, newState)
			writeJSONState(w, newState)
		}
	}
}

func giveBlackCardToWinner(w http.ResponseWriter, req *http.Request) error {
	/*var payload chooseWinnerPayload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	u, err := userFromSession(req)
	if err != nil {
		return err
	}
	game, err := getGame(req)
	pid, err := getPlayerIndex(req)
	if err != nil {
		return err
	}
	if pid != getState(req).CurrCzarIndex {
		return errors.New("Only the Czar can choose the winner!")
	}
	newS, err := game.GiveBlackCardToWinner(payload.Winner, getState(req))
	if err != nil {
		return err
	}
	updateState(req, newS)
	writeJSONState(w, newS)*/
	return nil
}

func playCards(w http.ResponseWriter, req *http.Request) error {
	/*var payload playCardsPayload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	pid, err := getPlayerIndex(req)
	if err != nil {
		return err
	}
	newS, err := game.PlayWhiteCards(pid, payload.CardIndexes, getState(req))
	if err != nil {
		return err
	} // oneline error handling when
	updateState(req, newS)
	writeJSONState(w, newS)*/
	return nil
}
