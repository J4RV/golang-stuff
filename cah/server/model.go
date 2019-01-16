package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/j4rv/golang-stuff/cah/game"
)

/* Client payloads */

type chooseWinnerPayload struct {
	Winner int `json:"winner"`
}

type playCardsPayload struct {
	CardIndexes []int `json:"cardIndexes"`
}

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* Server responses */

type playerInfo struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	HandSize         int              `json:"handSize"`
	WhiteCardsInPlay int              `json:"whiteCardsInPlay"`
	Points           []game.BlackCard `json:"points"`
}

type sinnerPlay struct {
	ID         int              `json:"id"`
	WhiteCards []game.WhiteCard `json:"whiteCards"`
}

type gameStateResponse struct {
	Phase           int              `json:"phase"`
	Players         []playerInfo     `json:"players"`
	CurrCzarID      int              `json:"currentCzarID"`
	BlackCardInPlay game.BlackCard   `json:"blackCardInPlay"`
	SinnerPlays     []sinnerPlay     `json:"sinnerPlays"`
	DiscardPile     []game.WhiteCard `json:"discardPile"`
	MyPlayer        game.Player      `json:"myPlayer"`
}

/* Write helpers */

func writeResponse(w http.ResponseWriter, obj interface{}) {
	j, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", j)
	}
}
