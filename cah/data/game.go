package data

import (
	"errors"
	"log"

	"github.com/j4rv/golang-stuff/cah"
)

type gameMemStore struct {
	abstractMemStore
	stateStore cah.GameStateStore
	games      map[int]cah.Game
}

func NewGameStore(ss cah.GameStateStore) *gameMemStore {
	return &gameMemStore{
		stateStore: ss,
		games:      map[int]cah.Game{},
	}
}

func (store *gameMemStore) Create(g cah.Game) error {
	if g.ID != 0 {
		return errors.New("Tried to create a game but its ID was not zero")
	}
	g.ID = store.nextID()
	store.games[g.ID] = g
	return nil
}

func (store *gameMemStore) ByStatePhase(p cah.Phase) []cah.Game {
	ret := []cah.Game{}
	for _, g := range store.games {
		state, err := store.stateStore.ByID(g.StateID)
		if err != nil {
			log.Printf("Possible inconsistency in game with ID %d, pointing to non existing state %d\n", g.ID, g.StateID)
			continue
		}
		if state.Phase == p {
			ret = append(ret, g)
		}
	}
	return ret
}
