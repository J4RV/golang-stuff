package cah

type Phase uint8

const (
	NotStarted Phase = iota
	SinnersPlaying
	CzarChoosingWinner
	Finished
)

var phases = [...]string{
	"NotStarted",
	"SinnersPlaying",
	"CzarChoosingWinner",
	"Finished",
}

func (p Phase) String() string {
	return phases[p]
}
