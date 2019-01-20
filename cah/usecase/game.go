package usecase

import (
	"errors"
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

var nilBlackCard = cah.BlackCard{}

type errorEmptyBlackDeck struct{}

func (e errorEmptyBlackDeck) Error() string {
	return "Zero cards left in black deck"
}

type gameController struct {
	store   cah.GameStore
	options Options
}

func NewGameUsecase(store cah.GameStore) *gameController {
	return &gameController{store: store}
}

func (control gameController) ByID(id int) (cah.Game, error) {
	return control.store.ByID(id)
}

func playersDraw(s *cah.Game) {
	for _, p := range s.Players {
		for len(p.Hand) < s.HandSize {
			p.Hand = append(p.Hand, s.DrawWhite())
		}
	}
}

func (control gameController) Start(p []*cah.Player, g cah.Game, opts ...cah.Option) (cah.Game, error) {
	if g.Phase != cah.Starting {
		return g, fmt.Errorf("Tried to start the game but it has already started; current phase: '%s'", g.Phase)
	}
	if p == nil || len(p) < 3 {
		return g, fmt.Errorf("Wrong players argument: '%v'. Cannot be nil and the minimum amount of players is 3", p)
	}
	ret := g.Clone()
	ret, err := control.putBlackCardInPlay(g)
	if err != nil {
		return g, err
	}
	ret.Players = p
	playersDraw(&ret)
	applyOptions(&ret, opts...)
	err = control.store.Update(ret)
	if err != nil {
		return g, err
	}
	return ret, nil
}

func (control gameController) End(g cah.Game) (cah.Game, error) {
	if g.Phase == cah.Finished {
		return g, errors.New("Tried to end a game but it has already finished")
	}
	ret := g.Clone()
	ret.Phase = cah.Finished
	err := control.store.Update(ret)
	if err != nil {
		return g, err
	}
	return ret, nil
}

func (control gameController) putBlackCardInPlay(g cah.Game) (cah.Game, error) {
	if g.BlackCardInPlay != nilBlackCard {
		return g, errors.New("Tried to put a black card in play but there is already a black card in play")
	}
	if g.Phase == cah.Finished {
		return g, errors.New("Tried to put a black card in play but the game has already finished")
	}
	if len(g.BlackDeck) == 0 {
		return g, errorEmptyBlackDeck{}
	}
	ret := g.Clone()
	ret.BlackCardInPlay = ret.BlackDeck[0]
	ret.BlackDeck = ret.BlackDeck[1:]
	ret.Phase = cah.SinnersPlaying
	return ret, nil
}

func (control gameController) GiveBlackCardToWinner(wID int, g cah.Game) (cah.Game, error) {
	err := giveBlackCardToWinnerChecks(wID, g)
	if err != nil {
		return g, err
	}
	ret := g.Clone()
	var winner *cah.Player
	for _, p := range ret.Players {
		if p.User.ID == wID {
			winner = p
		}
	}
	if winner == nil {
		return g, fmt.Errorf("Invalid winner id %d", wID)
	}
	winner.Points = append(winner.Points, ret.BlackCardInPlay)
	ret.BlackCardInPlay = nilBlackCard
	for _, p := range g.Players {
		p.WhiteCardsInPlay = []cah.WhiteCard{}
	}
	ret, _ = control.nextCzar(ret)
	if (len(ret.BlackDeck)) == 0 {
		ret, err = control.End(ret)
		return ret, err
	}
	ret, err = control.putBlackCardInPlay(ret)
	if err != nil {
		_, ok := err.(*errorEmptyBlackDeck)
		if !ok {
			return g, err
		}
		ret, _ = control.End(ret)
	}
	playersDraw(&ret)
	err = control.store.Update(ret)
	if err != nil {
		return g, err
	}
	return ret, nil
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

func (control gameController) PlayWhiteCards(p int, cs []int, g cah.Game) (cah.Game, error) {
	if p < 0 || p >= len(g.Players) {
		return g, errors.New("Non valid sinner index")
	}
	if p == g.CurrCzarIndex {
		return g, errors.New("The Czar cannot play white cards")
	}
	if len(cs)+len(g.Players[p].WhiteCardsInPlay) > g.BlackCardInPlay.BlanksAmount {
		return g, fmt.Errorf("Invalid amount of white cards to play, expected %d but got %d",
			g.BlackCardInPlay.BlanksAmount,
			len(cs)+len(g.Players[p].WhiteCardsInPlay))
	}
	ret := g.Clone()
	player := ret.Players[p]
	newCardsPlayed, err := player.ExtractCardsFromHand(cs)
	if err != nil {
		return ret, err
	}
	player.WhiteCardsInPlay = append(player.WhiteCardsInPlay, newCardsPlayed...)
	if control.AllSinnersPlayedTheirCards(ret) {
		ret.Phase = cah.CzarChoosingWinner
	}
	err = control.store.Update(ret)
	if err != nil {
		return g, err
	}
	return ret, nil
}

func (_ gameController) AllSinnersPlayedTheirCards(s cah.Game) bool {
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

func (_ gameController) nextCzar(s cah.Game) (cah.Game, error) {
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
