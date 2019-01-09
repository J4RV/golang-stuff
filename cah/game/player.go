package game

import (
	"errors"
	"fmt"
)

type Player struct {
	Name             string
	Hand             []WhiteCard
	WhiteCardsInPlay []WhiteCard
	Points           []BlackCard
}

func (p *Player) removeCardFromHand(i int) error {
	if i < 0 || i >= len(p.Hand) {
		msg := fmt.Sprintf("Index out of bounds at RemoveCardFromHand. Index: %d, Hand size: %d", i, len(p.Hand))
		return errors.New(msg)
	}
	p.Hand[i] = p.Hand[len(p.Hand)-1]
	p.Hand[len(p.Hand)-1] = nil
	p.Hand = p.Hand[:len(p.Hand)-1]
	return nil
}

func GetRandomPlayers() []Player {
	p := make([]Player, 3)
	p[0] = Player{Name: "One"}
	p[0] = Player{Name: "Two"}
	p[0] = Player{Name: "Three"}
	return p
}
