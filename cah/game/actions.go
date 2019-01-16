package game

import (
	"errors"
	"fmt"
)

func playersDraw(s *State) {
	for _, p := range s.Players {
		for len(p.Hand) < playerHandSize {
			p.Hand = append(p.Hand, s.drawWhite())
		}
	}
}

func PutBlackCardInPlay(s State) (State, error) {
	if s.BlackCardInPlay != nil {
		return s, errors.New("Tried to put a black card in play but there is already a black card in play")
	}
	if s.Phase == Finished {
		return s, errors.New("Tried to put a black card in play but the game has already finished")
	}
	if len(s.BlackDeck) == 0 {
		return s, errors.New("Tried to put a black card in play but the black deck is empty")
	}
	res := s.Clone()
	res.BlackCardInPlay = res.BlackDeck[0]
	res.BlackDeck = res.BlackDeck[1:]
	res.Phase = SinnersPlaying
	return res, nil
}

func GiveBlackCardToWinner(w int, s State) (State, error) {
	err := giveBlackCardToWinnerChecks(w, s)
	if err != nil {
		return s, err
	}
	res := s.Clone()
	res.Players[w].Points = append(res.Players[w].Points, res.BlackCardInPlay)
	res.BlackCardInPlay = nil
	for _, p := range s.Players {
		p.WhiteCardsInPlay = []WhiteCard{}
	}
	res, _ = NextCzar(res)
	res, _ = PutBlackCardInPlay(res)
	playersDraw(&res)
	return res, nil
}

func giveBlackCardToWinnerChecks(w int, s State) error {
	if s.Phase != CzarChoosingWinner {
		return fmt.Errorf("Tried to choose a winner in a non valid phase '%d'", s.Phase)
	}
	if w < 0 || w >= len(s.Players) {
		return errors.New("Non valid player index")
	}
	for i, p := range s.Players {
		if i == s.CurrCzarIndex {
			continue
		}
		if len(p.WhiteCardsInPlay) != s.BlackCardInPlay.GetBlanksAmount() {
			return errors.New("Not all sinners have played their cards")
		}
	}
	return nil
}

func PlayWhiteCards(p int, cs []int, s State) (State, error) {
	if p < 0 || p >= len(s.Players) {
		return s, errors.New("Non valid sinner index")
	}
	if p == s.CurrCzarIndex {
		return s, errors.New("The Czar cannot play white cards")
	}
	if len(cs)+len(s.Players[p].WhiteCardsInPlay) > s.BlackCardInPlay.GetBlanksAmount() {
		return s, fmt.Errorf("Invalid amount of white cards to play, expected %d but got %d",
			s.BlackCardInPlay.GetBlanksAmount(),
			len(cs)+len(s.Players[p].WhiteCardsInPlay))
	}
	res := s.Clone()
	player := res.Players[p]
	newCardsPlayed, err := player.extractCardsFromHand(cs)
	if err != nil {
		return res, err
	}
	player.WhiteCardsInPlay = append(player.WhiteCardsInPlay, newCardsPlayed...)
	if AllSinnersPlayedTheirCards(res) {
		res.Phase = CzarChoosingWinner
	}
	return res, nil
}

func AllSinnersPlayedTheirCards(s State) bool {
	for i, p := range s.Players {
		if i == s.CurrCzarIndex {
			continue
		}
		if len(p.WhiteCardsInPlay) != s.BlackCardInPlay.GetBlanksAmount() {
			return false
		}
	}
	return true
}

func NextCzar(s State) (State, error) {
	if s.BlackCardInPlay != nil {
		return s, errors.New("Tried to rotate to the next Czar but there is still a black card in play")
	}
	if s.Phase == Finished {
		return s, errors.New("Tried to rotate to the next Czar but the game has already finished")
	}
	res := s.Clone()
	res.CurrCzarIndex++
	if res.CurrCzarIndex == len(s.Players) {
		res.CurrCzarIndex = 0
	}
	return res, nil
}
