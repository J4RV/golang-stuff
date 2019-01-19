package game

import (
	"errors"
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

var nilBlackCard = cah.BlackCard{}

type GameController struct{}

func playersDraw(s *cah.Game) {
	for _, p := range s.Players {
		for len(p.Hand) < s.HandSize {
			p.Hand = append(p.Hand, s.DrawWhite())
		}
	}
}

func (control GameController) Start(g cah.Game) (cah.Game, error) {
	if g.Phase != cah.Starting {
		return g, fmt.Errorf("Tried to start the game but it has already started; current phase: '%s'", g.Phase)
	}
	ret, err := control.PutBlackCardInPlay(g)
	if err != nil {
		return g, err
	}
	playersDraw(&ret)
	return ret, nil
}

func (_ GameController) PutBlackCardInPlay(s cah.Game) (cah.Game, error) {
	if s.BlackCardInPlay != nilBlackCard {
		return s, errors.New("Tried to put a black card in play but there is already a black card in play")
	}
	if s.Phase == cah.Finished {
		return s, errors.New("Tried to put a black card in play but the game has already finished")
	}
	if len(s.BlackDeck) == 0 {
		return s, errors.New("Tried to put a black card in play but the black deck is empty")
	}
	res := s.Clone()
	res.BlackCardInPlay = res.BlackDeck[0]
	res.BlackDeck = res.BlackDeck[1:]
	res.Phase = cah.SinnersPlaying
	return res, nil
}

func (control GameController) GiveBlackCardToWinner(wId int, s cah.Game) (cah.Game, error) {
	err := giveBlackCardToWinnerChecks(wId, s)
	if err != nil {
		return s, err
	}
	res := s.Clone()
	var winner *cah.Player
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
		p.WhiteCardsInPlay = []cah.WhiteCard{}
	}
	res, _ = control.NextCzar(res)
	res, _ = control.PutBlackCardInPlay(res)
	playersDraw(&res)
	return res, nil
}

func giveBlackCardToWinnerChecks(w int, s cah.Game) error {
	if s.Phase != cah.CzarChoosingWinner {
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

func (control GameController) PlayWhiteCards(p int, cs []int, s cah.Game) (cah.Game, error) {
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
	if control.AllSinnersPlayedTheirCards(res) {
		res.Phase = cah.CzarChoosingWinner
	}
	return res, nil
}

func (_ GameController) AllSinnersPlayedTheirCards(s cah.Game) bool {
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

func (_ GameController) NextCzar(s cah.Game) (cah.Game, error) {
	if s.BlackCardInPlay != nilBlackCard {
		return s, errors.New("Tried to rotate to the next Czar but there is still a black card in play")
	}
	if s.Phase == cah.Finished {
		return s, errors.New("Tried to rotate to the next Czar but the game has already finished")
	}
	res := s.Clone()
	res.CurrCzarIndex++
	if res.CurrCzarIndex == len(s.Players) {
		res.CurrCzarIndex = 0
	}
	return res, nil
}
