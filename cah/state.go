package cah

type GameController interface {
	Start(g Game) (Game, error)
	PutBlackCardInPlay(g Game) (Game, error)
	GiveBlackCardToWinner(wId int, g Game) (Game, error)
	PlayWhiteCards(p int, cs []int, g Game) (Game, error)
	AllSinnersPlayedTheirCards(g Game) bool
	NextCzar(g Game) (Game, error)
}

type Game struct {
	Phase           Phase       `json:"phase"`
	Players         []*Player   `json:"players"`
	BlackDeck       []BlackCard `json:"blackDeck"`
	WhiteDeck       []WhiteCard `json:"whiteDeck"`
	CurrCzarIndex   int         `json:"currentCzarIndex"`
	BlackCardInPlay BlackCard   `json:"blackCardInPlay"`
	DiscardPile     []WhiteCard `json:"discardPile"`
	HandSize        int         `json:"handSize"`
}

func (s *Game) DrawWhite() WhiteCard {
	ret := s.WhiteDeck[0]
	s.WhiteDeck = s.WhiteDeck[1:]
	return ret
}

func (s Game) CurrCzar() *Player {
	return s.Players[s.CurrCzarIndex]
}

func (s Game) Equal(other Game) bool {
	// Fast comparisons before checking lists
	equal := s.Phase == other.Phase
	equal = equal && s.CurrCzarIndex == other.CurrCzarIndex
	equal = equal && s.BlackCardInPlay == other.BlackCardInPlay
	equal = equal && s.HandSize == other.HandSize
	if !equal {
		return false
	}
	// Check lists one by one
	for i := range s.Players {
		if s.Players[i] != other.Players[i] {
			return false
		}
	}
	for i := range s.BlackDeck {
		if s.BlackDeck[i] != other.BlackDeck[i] {
			return false
		}
	}
	for i := range s.WhiteDeck {
		if s.WhiteDeck[i] != other.WhiteDeck[i] {
			return false
		}
	}
	for i := range s.DiscardPile {
		if s.DiscardPile[i] != other.DiscardPile[i] {
			return false
		}
	}
	return true
}

func (s Game) Clone() Game {
	res := Game{
		Phase:           s.Phase,
		CurrCzarIndex:   s.CurrCzarIndex,
		BlackCardInPlay: s.BlackCardInPlay,
		HandSize:        s.HandSize,
		Players:         make([]*Player, len(s.Players)),
		BlackDeck:       make([]BlackCard, len(s.BlackDeck)),
		WhiteDeck:       make([]WhiteCard, len(s.WhiteDeck)),
		DiscardPile:     make([]WhiteCard, len(s.DiscardPile)),
	}
	copy(res.Players, s.Players)
	copy(res.BlackDeck, s.BlackDeck)
	copy(res.WhiteDeck, s.WhiteDeck)
	copy(res.DiscardPile, s.DiscardPile)
	return res
}
