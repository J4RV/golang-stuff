package game

import (
	"errors"
	"fmt"
)

type Player struct {
	Name             string      `json:"name"`
	Hand             []WhiteCard `json:"hand"`
	WhiteCardsInPlay []WhiteCard `json:"whiteCardsInPlay"`
	Points           []BlackCard `json:"points"`
}

func (p *Player) extractCardFromHand(i int) (WhiteCard, error) {
	if i < 0 || i >= len(p.Hand) {
		msg := fmt.Sprintf("Index out of bounds at RemoveCardFromHand. Index: %d, Hand size: %d", i, len(p.Hand))
		return nil, errors.New(msg)
	}
	c := p.Hand[i]
	p.Hand[i] = p.Hand[len(p.Hand)-1]
	p.Hand[len(p.Hand)-1] = nil
	p.Hand = p.Hand[:len(p.Hand)-1]
	return c, nil
}

func GetRandomPlayers() []*Player {
	p := make([]*Player, 3)
	p[0] = &Player{Name: "Rojo", WhiteCardsInPlay: []WhiteCard{}, Points: []BlackCard{}}
	p[1] = &Player{Name: "Jury", WhiteCardsInPlay: []WhiteCard{}, Points: []BlackCard{}}
	p[2] = &Player{Name: "Paul", WhiteCardsInPlay: []WhiteCard{}, Points: []BlackCard{}}
	return p
}
