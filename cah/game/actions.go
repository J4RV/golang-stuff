package game

import (
	"errors"
	"fmt"

	"github.com/j4rv/golang-stuff/cah/model"
)

func playersDraw(s *model.State) {
	for _, p := range s.Players {
		for len(p.Hand) < playerHandSize {
			p.Hand = append(p.Hand, s.DrawWhite())
		}
	}
}

func PutBlackCardInPlay(s model.State) (model.State, error) {
	if s.BlackCardInPlay != nilBlackCard {
		return s, errors.New("Tried to put a black card in play but there is already a black card in play")
	}
	if s.Phase == model.Finished {
		return s, errors.New("Tried to put a black card in play but the game has already finished")
	}
	if len(s.BlackDeck) == 0 {
		return s, errors.New("Tried to put a black card in play but the black deck is empty")
	}
	res := s.Clone()
	res.BlackCardInPlay = res.BlackDeck[0]
	res.BlackDeck = res.BlackDeck[1:]
	res.Phase = model.SinnersPlaying
	return res, nil
}

func GiveBlackCardToWinner(wId int, s model.State) (model.State, error) {
	err := giveBlackCardToWinnerChecks(wId, s)
	if err != nil {
		return s, err
	}
	res := s.Clone()
	var winner *model.Player
	for _, p := range res.Players {
		if p.ID == wId {
			winner = p
		}
	}
	if winner == nil {
		return s, fmt.Errorf("Invalid winner id %d", wId)
	}
	winner.Points = append(winner.Points, res.BlackCardInPlay)
	res.BlackCardInPlay = nilBlackCard
	for _, p := range s.Players {
		p.WhiteCardsInPlay = []model.WhiteCard{}
	}
	res, _ = NextCzar(res)
	res, _ = PutBlackCardInPlay(res)
	playersDraw(&res)
	return res, nil
}

func giveBlackCardToWinnerChecks(w int, s model.State) error {
	if s.Phase != model.CzarChoosingWinner {
		return fmt.Errorf("Tried to choose a winner in a non valid phase '%d'", s.Phase)
	}
	for i, p := range s.Players {
		if i == s.CurrCzarIndex {
			continue
		}
		if len(p.WhiteCardsInPlay) != s.BlackCardInPlay.BlanksAmount {
			return errors.New("Not all sinners have played their cards")
		}
	}
	return nil
}

func PlayWhiteCards(p int, cs []int, s model.State) (model.State, error) {
	if p < 0 || p >= len(s.Players) {
		return s, errors.New("Non valid sinner index")
	}
	if p == s.CurrCzarIndex {
		return s, errors.New("The Czar cannot play white cards")
	}
	if len(cs)+len(s.Players[p].WhiteCardsInPlay) > s.BlackCardInPlay.BlanksAmount {
		return s, fmt.Errorf("Invalid amount of white cards to play, expected %d but got %d",
			s.BlackCardInPlay.BlanksAmount,
			len(cs)+len(s.Players[p].WhiteCardsInPlay))
	}
	res := s.Clone()
	player := res.Players[p]
	newCardsPlayed, err := player.ExtractCardsFromHand(cs)
	if err != nil {
		return res, err
	}
	player.WhiteCardsInPlay = append(player.WhiteCardsInPlay, newCardsPlayed...)
	if AllSinnersPlayedTheirCards(res) {
		res.Phase = model.CzarChoosingWinner
	}
	return res, nil
}

func AllSinnersPlayedTheirCards(s model.State) bool {
	for i, p := range s.Players {
		if i == s.CurrCzarIndex {
			continue
		}
		if len(p.WhiteCardsInPlay) != s.BlackCardInPlay.BlanksAmount {
			return false
		}
	}
	return true
}

func NextCzar(s model.State) (model.State, error) {
	if s.BlackCardInPlay != nilBlackCard {
		return s, errors.New("Tried to rotate to the next Czar but there is still a black card in play")
	}
	if s.Phase == model.Finished {
		return s, errors.New("Tried to rotate to the next Czar but the game has already finished")
	}
	res := s.Clone()
	res.CurrCzarIndex++
	if res.CurrCzarIndex == len(s.Players) {
		res.CurrCzarIndex = 0
	}
	return res, nil
}
