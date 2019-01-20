package data

import (
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

type gameMemStore struct {
	abstractMemStore
	games map[int]*cah.Game
}

func NewGameStore() *gameMemStore {
	return &gameMemStore{
		games: make(map[int]*cah.Game),
	}
}

func (store *gameMemStore) Create(g cah.Game) (cah.Game, error) {
	g.ID = store.nextID()
	store.games[g.ID] = &g
	return g, nil
}

func (store *gameMemStore) ByID(id int) (cah.Game, error) {
	g, ok := store.games[id]
	if !ok {
		return *g, fmt.Errorf("No game found with ID %d", id)
	}
	return *g, nil
}

func (store *gameMemStore) Update(g cah.Game) error {
	_, err := store.ByID(g.ID)
	if err != nil {
		return err
	}
	store.games[g.ID] = &g
	return nil
}

func (store *gameMemStore) Delete(id int) error {
	_, err := store.ByID(id)
	if err != nil {
		return err
	}
	delete(store.games, id)
	return nil
}
