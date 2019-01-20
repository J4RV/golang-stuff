package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4rv/golang-stuff/cah"
)

func handleGameStates(r *mux.Router) {
	s := r.PathPrefix("/gamestate/{id}").Subrouter()
	s.Handle("/State", srvHandler(gameStateForUser)).Methods("GET")
	s.Handle("/ChooseWinner", srvHandler(chooseWinner)).Methods("POST")
	s.Handle("/PlayCards", srvHandler(playCards)).Methods("POST")
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

type fullPlayerInfo struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	Hand             []cah.WhiteCard `json:"hand" db:"hand"`
	WhiteCardsInPlay []cah.WhiteCard `json:"whiteCardsInPlay"`
	Points           []cah.BlackCard `json:"points"`
}

type sinnerPlay struct {
	ID         int             `json:"id"`
	WhiteCards []cah.WhiteCard `json:"whiteCards"`
}

type gameStateResponse struct {
	ID              int             `json:"id"`
	Phase           int             `json:"phase"`
	Players         []playerInfo    `json:"players"`
	CurrCzarID      int             `json:"currentCzarID"`
	BlackCardInPlay cah.BlackCard   `json:"blackCardInPlay"`
	SinnerPlays     []sinnerPlay    `json:"sinnerPlays"`
	DiscardPile     []cah.WhiteCard `json:"discardPile"`
	MyPlayer        fullPlayerInfo  `json:"myPlayer"`
}

func gameStateForUser(w http.ResponseWriter, req *http.Request) error {
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
		ID:              game.ID,
		Phase:           int(game.Phase),
		Players:         playersInfoFromGame(game),
		CurrCzarID:      game.Players[game.CurrCzarIndex].User.ID,
		BlackCardInPlay: game.BlackCardInPlay,
		SinnerPlays:     sinnerPlaysFromGame(game),
		DiscardPile:     game.DiscardPile,
		MyPlayer:        newFullPlayerInfo(*p),
	}
	writeResponse(w, response)
	return nil
}

func playersInfoFromGame(game cah.GameState) []playerInfo {
	ret := make([]playerInfo, len(game.Players))
	for i, p := range game.Players {
		ret[i] = newPlayerInfo(*p)
	}
	return ret
}

func newPlayerInfo(p cah.Player) playerInfo {
	return playerInfo{
		ID:               p.User.ID,
		Name:             p.User.Username,
		HandSize:         len(p.Hand),
		WhiteCardsInPlay: len(p.WhiteCardsInPlay),
		Points:           p.Points,
	}
}

func newFullPlayerInfo(player cah.Player) fullPlayerInfo {
	return fullPlayerInfo{
		ID:               player.User.ID,
		Name:             player.User.Username,
		Hand:             player.Hand,
		WhiteCardsInPlay: player.WhiteCardsInPlay,
	}
}

func sinnerPlaysFromGame(game cah.GameState) []sinnerPlay {
	if !usecase.GameState.AllSinnersPlayedTheirCards(game) {
		return []sinnerPlay{}
	}
	ret := make([]sinnerPlay, len(game.Players))
	for i, p := range game.Players {
		ret[i] = sinnerPlay{
			ID:         p.User.ID,
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

func chooseWinner(w http.ResponseWriter, req *http.Request) error {
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
	_, err = usecase.GameState.GiveBlackCardToWinner(payload.Winner, game)
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
	_, err = usecase.GameState.PlayWhiteCards(pid, payload.CardIndexes, game)
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
