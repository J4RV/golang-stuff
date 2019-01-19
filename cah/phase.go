package cah

type Phase uint8

const (
	Starting Phase = iota
	SinnersPlaying
	CzarChoosingWinner
	Finished
)

var phases = [...]string{
	"Starting",
	"SinnersPlaying",
	"CzarChoosingWinner",
	"Finished",
}

func (p Phase) String() string {
	return phases[p]
}
