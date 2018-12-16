package main

import (
	"fmt"

	. "github.com/j4rv/golang-stuff/unocards"
)

func main() {
	deck := NewVanilla()
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
	fmt.Printf("%+v", state)
}
