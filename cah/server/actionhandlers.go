package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/game"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func simpleCAHActionHandler(f func(game.State) (game.State, error)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		s := getState(req)
		newS, err := f(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			updateState(req, newS)
			writeJSONState(w, newS)
		}
	}
}

func giveBlackCardToWinner(w http.ResponseWriter, req *http.Request) error {
	var payload chooseWinnerPayload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
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
	writeJSONState(w, newS)
	return nil
}

func playCards(w http.ResponseWriter, req *http.Request) error {
	var payload playCardsPayload
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
	writeJSONState(w, newS)
	return nil
}
