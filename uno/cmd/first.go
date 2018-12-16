package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/j4rv/golang-stuff/uno/cards"
	"github.com/j4rv/golang-stuff/uno/game"
)

func init() {
	seed := time.Now().UnixNano()
	log.Println("Seed used:", seed)
	rand.Seed(seed)
}

func initState() game.State {
	player := game.Player{
		Name: "Rojo",
	}
	deck := cards.NewVanilla(cards.Shuffle)
	return game.State{
		Players: []game.Player{player},
		Deck:    deck,
	}
}

func main() {
	state := initState()
	for state.Phase != game.Finished {
		state = RunPhase(state)
	}
	log.Println("GAME FINISHED!")
}

func RunPhase(s game.State) game.State {
	switch s.Phase {

	case game.Starting:
		return game.Start(s)

	case game.PlayerTurn:
		return AskPlayerInput(s)

	case game.ProcessingPlayedCard:
		if len(s.CurrPlayer().Hand) == 0 {
			return game.End(s)
		}
		return game.ProcessTopCard(s)

	case game.ChoosingColor:
		return AskPlayerInput(s)

	default:
		panic("Reached non controlled phase:" + string(s.Phase))
	}
}

func AskPlayerInput(s game.State) game.State {
	res := s
	player := s.CurrPlayerIndex

	for res.Phase == game.PlayerTurn || player != res.CurrPlayerIndex {
		res = AskPlayerTurn(res)
	}

	for res.Phase == game.ChoosingColor || player != res.CurrPlayerIndex {
		res = AskChooseColor(res)
	}

	return res
}

func AskPlayerTurn(s game.State) game.State {
	res := s

	log.Printf("Starting %s's turn\n", res.CurrPlayer().Name)
	log.Printf("Current top card: %v\n", res.TopCard())
	log.Printf("Your hand: %v\n", cards.IndexedCardsString(res.CurrPlayer().Hand))

	if res.DrawnThisTurn {
		log.Println("Do you want to (p)ass or play the nth card?")
	} else {
		log.Println("Do you want to (d)raw or play the nth card?")
	}

	var input string
	fmt.Scanf("%s\n", &input)
	switch input {
	case "d":
		res = game.Draw(res)
	case "p":
		res = game.Pass(res)
	case "debug":
		log.Printf("%+v", res)
	default:
		i, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			log.Println("Not a valid action:", input)
		} else {
			res = game.PlayCard(i, res)
		}
	}

	return res
}

func AskChooseColor(s game.State) game.State {
	res := s

	log.Println("Choose a color: (r)ed, (y)ellow, (g)reen or (b)lue")
	log.Printf("Your hand: %v\n", cards.IndexedCardsString(res.CurrPlayer().Hand))

	var input string
	fmt.Scanf("%s\n", &input)
	switch input {
	case "r":
		res = game.ChangeColor(cards.Red, res)
	case "y":
		res = game.ChangeColor(cards.Yellow, res)
	case "g":
		res = game.ChangeColor(cards.Green, res)
	case "b":
		res = game.ChangeColor(cards.Blue, res)
	default:
		log.Println("Not a valid action:", input)
	}

	return res
}

/*
func playerturn(s *State) {
	log.Printf("Starting %s's turn\n", s.CurrPlayer().Name)
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
			playerr := PlayCard(i, s)
			if playerr == nil {
				return
			} else {
				fmt.Print(err)
			}
		}
		fmt.Println("Not a valid action:", input)
	}
}
*/
