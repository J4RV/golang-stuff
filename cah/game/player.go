package game

import (
	"errors"
	"fmt"
	"sort"
)

type Player struct {
	ID               int
	Name             string      `json:"name"`
	Hand             []WhiteCard `json:"hand"`
	WhiteCardsInPlay []WhiteCard `json:"whiteCardsInPlay"`
	Points           []BlackCard `json:"points"`
}

func NewPlayer(id int, name string) *Player {
	return &Player{
		ID:               id,
		Name:             name,
		WhiteCardsInPlay: []WhiteCard{},
		Points:           []BlackCard{}}
}

func (p *Player) removeCardFromHand(i int) error {
	if i < 0 || i >= len(p.Hand) {
		msg := fmt.Sprintf("Index out of bounds. Index: %d, Hand size: %d", i, len(p.Hand))
		return errors.New(msg)
	}
	p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
	return nil
}

func (p *Player) extractCardsFromHand(indexes []int) ([]WhiteCard, error) {
	ret := make([]WhiteCard, len(indexes))

	for iter, index := range indexes {
		if index < 0 || index >= len(p.Hand) {
			return nil, fmt.Errorf("Non valid white card index: %d", index)
		}
		c := p.Hand[index]
		ret[iter] = c
	}

	// Order ints from highest to lowest to prevent indexes out of bounds
	// since we will be altering the hand slice by removing cards from it
	iOrdered := make([]int, len(indexes))
	copy(iOrdered, indexes)
	sort.Sort(sort.Reverse(sort.IntSlice(iOrdered)))
	for _, index := range iOrdered {
		err := p.removeCardFromHand(index)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (p *Player) extractCardFromHand(i int) (WhiteCard, error) {
	ret, err := p.extractCardsFromHand([]int{i})
	if err != nil {
		return nil, err
	}
	return ret[0], nil
}
