package game

type Phase uint8

const (
	Starting Phase = iota
	SinnersPlaying
	CzarChoosingWinner
	Finished
)

type Card interface {
	GetText() string
}

type WhiteCard interface {
	Card
}

type BlackCard interface {
	Card
	GetBlanksAmount() int
}
