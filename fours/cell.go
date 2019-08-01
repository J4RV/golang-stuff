package main

type cell uint8

const (
	EMPTY cell = iota
	BLACK
	WHITE
	BoardVerticalSize   int = 6
	BoardHorizontalSize int = 9
)

func (c cell) String() string {
	switch c {
	case EMPTY:
		return " "
	case BLACK:
		return "■"
	case WHITE:
		return "□"
	default:
		panic("invalid cell")
	}
}

func (c cell) EnumString() string {
	switch c {
	case EMPTY:
		return "EMPTY (" + c.String() + ")"
	case BLACK:
		return "BLACK (" + c.String() + ")"
	case WHITE:
		return "WHITE (" + c.String() + ")"
	default:
		panic("invalid cell")
	}
}