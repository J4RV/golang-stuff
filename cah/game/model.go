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
	GetText() string
}

type BlackCard interface {
	GetText() string
	BlanksAmount() int
}
