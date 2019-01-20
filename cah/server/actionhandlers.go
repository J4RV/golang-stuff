package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/j4rv/golang-stuff/cah"
)

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
	GameID          int             `json:"gameID"`
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
	game, err := gameFromRequest(req)
	if err != nil {
		return err
	}
	p, err := player(game, u)
	if err != nil {
		return err
	}
	response := gameStateResponse{
		GameID:          game.ID,
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
	if !usecase.Game.AllSinnersPlayedTheirCards(game) {
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
	game, err := gameFromRequest(req)
	if err != nil {
		return err
	}
	pid, err := playerIndex(game, u)
	if err != nil {
		return err
	}
	if pid != game.CurrCzarIndex {
		return errors.New("Only the Czar can choose the winner")
	}
	_, err = usecase.Game.GiveBlackCardToWinner(payload.Winner, game)
	if err != nil {
		return err
	}
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
	game, err := gameFromRequest(req)
	if err != nil {
		return err
	}
	pid, err := playerIndex(game, u)
	if err != nil {
		return err
	}
	_, err = usecase.Game.PlayWhiteCards(pid, payload.CardIndexes, game)
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
