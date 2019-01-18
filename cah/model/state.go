package model

type State struct {
	Phase           Phase       `json:"phase"`
	Players         []*Player   `json:"players"`
	BlackDeck       []BlackCard `json:"blackDeck"`
	WhiteDeck       []WhiteCard `json:"whiteDeck"`
	CurrCzarIndex   int         `json:"currentCzarIndex"`
	BlackCardInPlay BlackCard   `json:"blackCardInPlay"`
	DiscardPile     []WhiteCard `json:"discardPile"`
}

func (s *State) DrawWhite() WhiteCard {
	ret := s.WhiteDeck[0]
	s.WhiteDeck = s.WhiteDeck[1:]
	return ret
}

func (s State) CurrCzar() *Player {
	return s.Players[s.CurrCzarIndex]
}

func (s State) Clone() State {
	res := State{
		Phase:           s.Phase,
		Players:         make([]*Player, len(s.Players)),
		BlackDeck:       make([]BlackCard, len(s.BlackDeck)),
		WhiteDeck:       make([]WhiteCard, len(s.WhiteDeck)),
		CurrCzarIndex:   s.CurrCzarIndex,
		BlackCardInPlay: s.BlackCardInPlay,
		DiscardPile:     make([]WhiteCard, len(s.DiscardPile)),
	}
	copy(res.Players, s.Players)
	copy(res.BlackDeck, s.BlackDeck)
	copy(res.WhiteDeck, s.WhiteDeck)
	copy(res.DiscardPile, s.DiscardPile)
	return res
}
