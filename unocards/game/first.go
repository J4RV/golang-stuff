package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	uno "github.com/j4rv/golang-stuff/unocards"
)

func init() {
	seed := time.Now().UnixNano()
	log.Println("Seed used:", seed)
	rand.Seed(seed)
}

func main() {
	deck := uno.NewVanilla(uno.Shuffle)
	player := Player{
		Name: "Rojo",
	}
	state := State{
		Players:    []*Player{&player},
		Currplayer: &player,
		Deck:       deck,
	}
	// start game
	for i := 0; i < 7; i++ {
		draw(&state.Deck, &(state.Currplayer.Hand))
	}
	draw(&state.Deck, &state.Played)
	fmt.Printf("First top card: %v\n", topcard(state))

	//prettyPrint(state)
	for len(state.Currplayer.Hand) > 0 {
		playerturn(&state)
	}
	fmt.Println("GAME FINISHED!")
}

func playerturn(s *State) {
	fmt.Printf("Starting %s's turn\n", s.Currplayer.Name)
	var input string
	var carddrawn bool
	for true {
		fmt.Printf("Current top card: %v\n", topcard(*s))
		fmt.Printf("Your hand: %v\n", uno.IndexedCardsString(s.Currplayer.Hand))
		if carddrawn {
			fmt.Println("Do you want to (p)ass or play the nth card?")
		} else {
			fmt.Println("Do you want to (d)raw or play the nth card?")
		}
		fmt.Scanf("%s\n", &input)
		switch input {
		case "d":
			if !carddrawn {
				drawn, err := draw(&s.Deck, &s.Currplayer.Hand)
				if err != nil {
					fmt.Println(err)
				} else {
					carddrawn = true
					fmt.Printf("Card drawn: %v\n", drawn)
				}
			}
			continue
		case "p":
			if carddrawn {
				return // do nothing, player is passing
			}
		default:
			i, err := strconv.ParseUint(input, 10, 32)
			if err != nil {
				fmt.Printf("Not a valid card index: %v\n", input)
			}
			playerr := playcardact(i, s)
			if playerr == nil {
				return
			} else {
				fmt.Print(err)
			}
		}
		fmt.Println("That is not a valid action:", input)
	}
}

func playcardact(i uint64, s *State) error {
	if i < 0 || int(i) >= len(s.Currplayer.Hand) {
		return fmt.Errorf("card index out of hand: %v", i)
	}
	c := s.Currplayer.Hand[i]
	if !canplay(c, *s) {
		return fmt.Errorf("not a valid card to play: %v", c)
	}
	log.Printf("playing card: %v", c)
	playcard(i, s)
	return nil
}

func playcard(i uint64, s *State) {
	h := &s.Currplayer.Hand
	c := s.Currplayer.Hand[i]
	*h = append((*h)[:i], (*h)[i+1:]...)
	s.Played = append(s.Played, c)
	processtopcard(s)
}

func canplay(c uno.Card, s State) bool {
	topc := topcard(s)
	if s.Drawacum == 0 {
		if c.Color == uno.Wild {
			return true
		}
		if topc.Color == uno.Wild {
			return c.Color == s.Currcolor
		}
		if c.Color == topc.Color {
			return true
		}
	}
	// even if Drawacum is not zero, you can answer a draw two with another draw two
	// same for wild draw four
	return c.Value == topc.Value
}

func processtopcard(s *State) {
	state := *s
	topc := topcard(state)
	switch topc.Value {
	case uno.Skip:
		state.Skip = true
	case uno.Reverse:
		state.Orderreversed = true
	case uno.DrawTwo:
		state.Drawacum += 2
	case uno.DrawFour:
		state.Drawacum += 4
	}
	state.Currcolor = topc.Color
}

func topcard(s State) uno.Card {
	cards := s.Played
	if len(cards) == 0 {
		panic("no cards to process, should never happen")
	}
	return cards[len(cards)-1]
}

func draw(from *[]uno.Card, to *[]uno.Card) (uno.Card, error) {
	var c uno.Card
	if len(*from) == 0 {
		return c, errors.New("no cards to draw")
	}
	c = (*from)[0]
	*from = (*from)[1:]
	*to = append(*to, c)
	return c, nil
}
