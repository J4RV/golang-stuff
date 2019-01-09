package data

import (
	"bufio"
	"log"
	"strings"
)

func initCards() {
	blackCards = append(blackCards, Card{IsBlack: true, Text: "How did I lose my virginity?"})
	blackCards = append(blackCards, Card{IsBlack: true, Text: "Why can't I sleep at night?"})
	blackCards = append(blackCards, Card{IsBlack: true, Text: "What's that smell?"})
	blackCards = append(blackCards, Card{IsBlack: true, Text: "I got 99 problems but _ ain't one."})
	blackCards = append(blackCards, Card{IsBlack: true, Text: "Maybe she's born with it. Maybe it's _."})

	allWhiteCards := `Seeing Granny naked
Elderly Japanese men.
Free samples.
Estrogen.
Sexual tension.
Famine.
A stray pube.
Men.
Heartwarming orphans.
Genuine human connection.
A bag of magic beans.
Repression.
Prancing.
My relationship status.
Overcompensation.
Peeing a little bit.
Pooping back and forth. Forever.
A ginger's freckled ballsack.
Testicular torsion.
The Devil himself.
The World of Warcraft.
Some bloody peace and quiet.
MechaHitler.
Being fabulous.
Pictures of boobs.
A gentle caress of the inner thigh.
Wiping her bum.
Doing a shit in Pudsey Bear's eyehole.
Lance Armstrong's missing testicle.
England.
The Pope.
Flying sex snakes.
Emma Watson.
My ex-wife.
Sexy pillow fights.
A Fleshlight.
Cybernetic enhancements.
Civilian casualties.
Magnets.
The female orgasm.
Bitches.
Madeline McCann.
Auschwitz.
Finger painting.`

	// This can be reutilized to read a file with all black/white cards
	s := bufio.NewScanner(strings.NewReader(allWhiteCards))
	for s.Scan() {
		whiteCards = append(whiteCards, Card{IsBlack: false, Text: s.Text()})
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
