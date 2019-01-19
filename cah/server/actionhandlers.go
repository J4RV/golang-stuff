package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/j4rv/golang-stuff/cah"
	"github.com/j4rv/golang-stuff/cah/game"
)

var gameControl cah.GameController

func init() {
	// should be injected
	gameControl = game.GameController{}
}

/*
GET GAME STATE
*/

type playerInfo struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	HandSize         int             `json:"handSize"`
	WhiteCardsInPlay int             `json:"whiteCardsInPlay"`
	Points           []cah.BlackCard `json:"points"`
}

type sinnerPlay struct {
	ID         int             `json:"id"`
	WhiteCards []cah.WhiteCard `json:"whiteCards"`
}

type gameStateResponse struct {
	Phase           int             `json:"phase"`
	Players         []playerInfo    `json:"players"`
	CurrCzarID      int             `json:"currentCzarID"`
	BlackCardInPlay cah.BlackCard   `json:"blackCardInPlay"`
	SinnerPlays     []sinnerPlay    `json:"sinnerPlays"`
	DiscardPile     []cah.WhiteCard `json:"discardPile"`
	MyPlayer        cah.Player      `json:"myPlayer"`
}

func getGameStateForUser(w http.ResponseWriter, req *http.Request) error {
	u, err := userFromSession(req)
	if err != nil {
		return err
	}
	game, err := getGame(req)
	if err != nil {
		return err
	}
	p, err := getPlayer(game, u)
	if err != nil {
		return err
	}
	response := gameStateResponse{
		Phase:           int(game.Phase),
		Players:         getPlayerInfo(game),
		CurrCzarID:      game.Players[game.CurrCzarIndex].ID,
		BlackCardInPlay: game.BlackCardInPlay,
		SinnerPlays:     getSinnerPlays(game),
		DiscardPile:     game.DiscardPile,
		MyPlayer:        *p,
	}
	writeResponse(w, response)
	return nil
}

func getPlayerInfo(game cah.Game) []playerInfo {
	ret := make([]playerInfo, len(game.Players))
	for i, p := range game.Players {
		ret[i] = gamePlayerToPlayerInfo(*p)
	}
	return ret
}

func gamePlayerToPlayerInfo(p cah.Player) playerInfo {
	return playerInfo{
		ID:               p.ID,
		Name:             p.User.Username,
		HandSize:         len(p.Hand),
		WhiteCardsInPlay: len(p.WhiteCardsInPlay),
		Points:           p.Points,
	}
}

func getSinnerPlays(game cah.Game) []sinnerPlay {
	if !gameControl.AllSinnersPlayedTheirCards(game) {
		return []sinnerPlay{}
	}
	ret := make([]sinnerPlay, len(game.Players))
	for i, p := range game.Players {
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
	game, err := getGame(req)
	if err != nil {
		return err
	}
	pid, err := getPlayerIndex(game, u)
	if err != nil {
		return err
	}
	if pid != game.CurrCzarIndex {
		return errors.New("Only the Czar can choose the winner")
	}
	newS, err := gameControl.GiveBlackCardToWinner(payload.Winner, game)
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
	game, err := getGame(req)
	if err != nil {
		return err
	}
	pid, err := getPlayerIndex(game, u)
	if err != nil {
		return err
	}
	newS, err := gameControl.PlayWhiteCards(pid, payload.CardIndexes, game)
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
