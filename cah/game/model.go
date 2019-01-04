package game

type Phase uint8

const (
	Starting Phase = iota
	SinnersPlaying
	CzarChoosingWinner
	Finished
)

type Card interface {
	text() string
}

type WhiteCard interface {
	text() string
}

type BlackCard interface {
	text() string
	blanksAmount() int
}
