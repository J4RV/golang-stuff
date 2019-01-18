package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/game"
	"github.com/j4rv/golang-stuff/cah/model"
)

/*
GET GAME STATE
*/

type playerInfo struct {
	ID               int               `json:"id"`
	Name             string            `json:"name"`
	HandSize         int               `json:"handSize"`
	WhiteCardsInPlay int               `json:"whiteCardsInPlay"`
	Points           []model.BlackCard `json:"points"`
}

type sinnerPlay struct {
	ID         int               `json:"id"`
	WhiteCards []model.WhiteCard `json:"whiteCards"`
}

type gameStateResponse struct {
	Phase           int               `json:"phase"`
	Players         []playerInfo      `json:"players"`
	CurrCzarID      int               `json:"currentCzarID"`
	BlackCardInPlay model.BlackCard   `json:"blackCardInPlay"`
	SinnerPlays     []sinnerPlay      `json:"sinnerPlays"`
	DiscardPile     []model.WhiteCard `json:"discardPile"`
	MyPlayer        model.Player      `json:"myPlayer"`
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

func gamePlayerToPlayerInfo(p model.Player) playerInfo {
	return playerInfo{
		ID:               p.ID,
		Name:             p.User.Username,
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

/*
CHOOSE WINNER
*/

type chooseWinnerPayload struct {
	Winner int `json:"winner"`
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
	return nil
}

/*
PLAY CARDS
*/

type playCardsPayload struct {
	CardIndexes []int `json:"cardIndexes"`
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
	err = updateGameState(req, newS)
	if err != nil {
		return err
	}
	return nil
}

// Utils

func writeResponse(w http.ResponseWriter, obj interface{}) {
	j, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", j)
	}
}
