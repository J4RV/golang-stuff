package data

import (
	"bufio"
	"log"
	"strings"
)

func initCards() {
	blackCards = append(blackCards, BlackCard{Card: Card{Text: "How did I lose my virginity?"}, BlanksAmount: 1})
	blackCards = append(blackCards, BlackCard{Card: Card{Text: "Why can't I sleep at night?"}, BlanksAmount: 1})
	blackCards = append(blackCards, BlackCard{Card: Card{Text: "What's that smell?"}, BlanksAmount: 1})
	blackCards = append(blackCards, BlackCard{Card: Card{Text: "I got 99 problems but _ ain't one."}, BlanksAmount: 1})
	blackCards = append(blackCards, BlackCard{Card: Card{Text: "Maybe she's born with it. Maybe it's _."}, BlanksAmount: 1})

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
		whiteCards = append(whiteCards, WhiteCard{Card{Text: s.Text()}})
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
