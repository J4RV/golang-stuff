package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/game"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := fn(w, req); err != nil {
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
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload chooseWinnerPayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	sg, err := getGame(req)
	if err != nil {
		return err
	}
	pid, err := getPlayerIndex(sg, u)
	if err != nil {
		return err
	}
	if pid != sg.state.CurrCzarIndex {
		return errors.New("Only the Czar can choose the winner")
	}
	newS, err := game.GiveBlackCardToWinner(payload.Winner, sg.state)
	if err != nil {
		return err
	}
	updateGameState(req, newS)
	writeJSONState(w, newS)
	return nil
}

func playCards(w http.ResponseWriter, req *http.Request) error {
	// User is logged
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	// Decode user's payload
	var payload playCardsPayload
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	sg, err := getGame(req)
	if err != nil {
		return err
	}
	pid, err := getPlayerIndex(sg, u)
	if err != nil {
		return err
	}
	newS, err := game.PlayWhiteCards(pid, payload.CardIndexes, sg.state)
	if err != nil {
		return err
	} // oneline error handling when
	updateGameState(req, newS)
	writeJSONState(w, newS)
	return nil
}
