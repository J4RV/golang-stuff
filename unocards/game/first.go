package main

import (
	"errors"
	"fmt"

	uno "github.com/j4rv/golang-stuff/unocards"
)

func main() {
	deck := uno.NewVanilla()
	player := Player{
		Name: "Rojo",
	}
	board := Board{
		Deck: deck,
	}
	state := State{
		Players:    []Player{player},
		Currplayer: player,
		Board:      board,
	}

	// start game
	for i := 0; i < 7; i++ {
		draw(&state.Board.Deck, &state.Currplayer.Hand)
	}
	draw(&state.Board.Deck, &state.Board.Played)

	fmt.Printf("%+v", state)
}

func draw(from *uno.Deck, to *[]uno.Card) error {
	if len(*from) == 0 {
		return errors.New("no cards to draw")
	}
	c := (*from)[0]
	*from = (*from)[1:]
	*to = append(*to, c)
	return nil
}
