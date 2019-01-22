package mem

import (
	"errors"
	"fmt"

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
	store.Lock()
	defer store.Unlock()
	if g.ID != 0 {
		return errors.New("Tried to create a game but its ID was not zero")
	}
	g.ID = store.nextID()
	store.games[g.ID] = g
	return nil
}

func (store *gameMemStore) ByID(id int) (cah.Game, error) {
	store.Lock()
	defer store.Unlock()
	g, ok := store.games[id]
	if !ok {
		return g, fmt.Errorf("No game found with id %d", id)
	}
	return g, nil
}

func (store *gameMemStore) ByStatePhase(p cah.Phase) []cah.Game {
	store.Lock()
	defer store.Unlock()
	ret := []cah.Game{}
	for _, g := range store.games {
		if g.StateID == 0 {
			ret = append(ret, g)
		}
	}
	return ret
}

func (store *gameMemStore) Update(g cah.Game) error {
	store.Lock()
	defer store.Unlock()
	store.games[g.ID] = g
	return nil
}
