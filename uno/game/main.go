package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/j4rv/golang-stuff/uno/cards"
)

func init() {
	seed := time.Now().UnixNano()
	log.Println("Seed used:", seed)
	rand.Seed(seed)
}

func initState() State {
	player := Player{Name: "Rojo"}
	player2 := Player{Name: "Jurado"}
	deck := cards.NewVanilla(cards.Shuffle)
	return State{
		Players: []Player{player, player2},
		Deck:    deck,
	}
}

func main() {
	state := initState()
	for state.Phase != Finished {
		state = RunPhase(state)
	}
	log.Println("GAME FINISHED!")
	log.Printf("%s won!", state.CurrPlayer().Name)
}

func RunPhase(s State) State {
	switch s.Phase {

	case Starting:
		return Start(s)

	case PlayerTurnStarting:
		if s.DrawAcum > 0 {
			return AskPlayerToDraw(s)
		} else {
			return AskPlayerTurn(s)
		}

	case PlayerThinkingAfterDraw:
		return AskPlayerAfterDraw(s)

	case PlayerChoosingColor:
		return AskChooseColor(s)

	case ProcessingCard:
		return ProcessTopCard(s)

	case ProcessingForcedDraw:
		return ProcessForcedDraw(s)

	default:
		panic("Reached non controlled phase:" + string(s.Phase))
	}
}

func AskPlayerToDraw(s State) State {
	res := s

	log.Printf("Starting %s's turn\n", res.CurrPlayer().Name)
	log.Printf("Current top card: %v\n", res.TopCard())
	log.Printf("Your hand: %v\n", cards.IndexedCardsString(res.CurrPlayer().Hand))
	log.Printf("Do you want to (d)raw %d cards or play the nth card?", res.DrawAcum)

	var input string
	fmt.Scanf("%s\n", &input)
	switch input {
	case "d":
		res = ForcedDraw(res)
	case "debug":
		log.Printf("%+v", res)
	default:
		i, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			log.Println("Not a valid action:", input)
		} else {
			res = PlayCard(i, res)
		}
	}

	return res
}

func AskPlayerAfterDraw(s State) State {
	res := s

	log.Printf("Starting %s's turn\n", res.CurrPlayer().Name)
	log.Printf("Current top card: %v\n", res.TopCard())
	log.Printf("Your hand: %v\n", cards.IndexedCardsString(res.CurrPlayer().Hand))
	log.Println("Do you want to (p)ass or play the nth card?")

	var input string
	fmt.Scanf("%s\n", &input)
	switch input {
	case "p":
		res = Pass(res)
	case "debug":
		log.Printf("%+v", res)
	default:
		i, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			log.Println("Not a valid action:", input)
		} else {
			res = PlayCard(i, res)
		}
	}

	return res
}

func AskPlayerTurn(s State) State {
	res := s

	log.Printf("Starting %s's turn\n", res.CurrPlayer().Name)
	log.Printf("Current top card: %v\n", res.TopCard())
	log.Printf("Your hand: %v\n", cards.IndexedCardsString(res.CurrPlayer().Hand))
	log.Println("Do you want to (d)raw or play the nth card?")

	var input string
	fmt.Scanf("%s\n", &input)
	switch input {
	case "d":
		res = DrawOne(res)
	case "debug":
		log.Printf("%+v", res)
	default:
		i, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			log.Println("Not a valid action:", input)
		} else {
			res = PlayCard(i, res)
		}
	}

	return res
}

func AskChooseColor(s State) State {
	res := s

	log.Printf("Your hand: %v\n", cards.IndexedCardsString(res.CurrPlayer().Hand))
	log.Println("Choose a color: (r)ed, (y)ellow, (g)reen or (b)lue")

	var input string
	fmt.Scanf("%s\n", &input)
	for true {
		switch input {
		case "r":
			return ChangeColor(cards.Red, res)
		case "y":
			return ChangeColor(cards.Yellow, res)
		case "g":
			return ChangeColor(cards.Green, res)
		case "b":
			return ChangeColor(cards.Blue, res)
		default:
			log.Println("Not a valid action:", input)
		}
	}
	panic("how did you get here")
}
