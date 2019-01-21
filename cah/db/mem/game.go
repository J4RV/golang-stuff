package mem

import (
	"errors"

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
	store.lock()
	defer store.release()
	if g.ID != 0 {
		return errors.New("Tried to create a game but its ID was not zero")
	}
	g.ID = store.nextID()
	store.games[g.ID] = g
	return nil
}

func (store *gameMemStore) ByStatePhase(p cah.Phase) []cah.Game {
	store.lock()
	defer store.release()
	ret := []cah.Game{}
	for _, g := range store.games {
		if g.StateID == 0 {
			ret = append(ret, g)
		}
	}
	return ret
}
