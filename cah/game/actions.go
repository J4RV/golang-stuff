package game

import "errors"

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
	// TODO w index in range of players check
	res := s.Clone()
	res.Players[w].Points = append(res.Players[w].Points, res.BlackCardInPlay)
	res.BlackCardInPlay = nil
	return res, nil
}

func PlayWhiteCards(p int, cs []int, s State) (State, error) {
	// TODO p index in range of players check
	// TODO cs indexes in range of players hand and not repeated indexes
	// TODO cs indexes len == current black card blanks
	player := &s.Players[p]
	for _, i := range cs {
		err := player.removeCardFromHand(i)
		if err != nil {
			return s, err
		}
	}
	return s, nil //TODO
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
