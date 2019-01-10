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
	if s.Phase != CzarChoosingWinner {
		return s, errors.New("Tried to choose a winner in a non valid phase")
	}
	if w < 0 || w >= len(s.Players) {
		return s, errors.New("Non valid player index")
	}
	res := s.Clone()
	res.Players[w].Points = append(res.Players[w].Points, res.BlackCardInPlay)
	res.BlackCardInPlay = nil
	return res, nil
}

func PlayWhiteCards(p int, cs []int, s State) (State, error) {
	if p < 0 || p >= len(s.Players) {
		return s, errors.New("Non valid player index")
	}
	player := &s.Players[p]
	for _, i := range cs {
		if i < 0 || i >= len(player.Hand) {
			return s, errors.New("Non valid white card index")
		}
		// TODO cs indexes: check not repeated indexes
	}
	if len(cs) != s.BlackCardInPlay.BlanksAmount() {
		return s, errors.New("Invalid amount of blank cards")
	}
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
