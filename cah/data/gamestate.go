package data

import (
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

type stateMemStore struct {
	abstractMemStore
	games map[int]*cah.GameState
}

func NewGameStateStore() *stateMemStore {
	return &stateMemStore{
		games: make(map[int]*cah.GameState),
	}
}

func (store *stateMemStore) Create(g cah.GameState) (cah.GameState, error) {
	g.ID = store.nextID()
	store.games[g.ID] = &g
	return g, nil
}

func (store *stateMemStore) ByID(id int) (cah.GameState, error) {
	g, ok := store.games[id]
	if !ok {
		return *g, fmt.Errorf("No game found with ID %d", id)
	}
	return *g, nil
}

func (store *stateMemStore) Update(g cah.GameState) error {
	_, err := store.ByID(g.ID)
	if err != nil {
		return err
	}
	store.games[g.ID] = &g
	return nil
}

func (store *stateMemStore) Delete(id int) error {
	_, err := store.ByID(id)
	if err != nil {
		return err
	}
	delete(store.games, id)
	return nil
}