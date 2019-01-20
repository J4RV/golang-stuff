package data

import (
	"errors"

	"github.com/j4rv/golang-stuff/cah"
)

type gameMemStore struct {
	abstractMemStore
	games map[int]cah.Game
}

func (store gameMemStore) Create(g cah.Game) error {
	if g.ID != 0 {
		return errors.New("Tried to create a game but its ID was not zero")
	}
	g.ID = store.nextID()
	store.games[g.ID] = g
	return nil
}

func (store gameMemStore) ByStatePhase(p cah.Phase) []cah.Game {
	ret := []cah.Game{}
	for _, g := range store.games {
		if g.State.Phase == p {
			ret = append(ret, g)
		}
	}
	return ret
}
