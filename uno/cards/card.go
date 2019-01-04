//go:generate stringer -type=Color,Value -linecomment=true

package cards

import (
	"fmt"

	"github.com/fatih/color"
)

//Color represents the color of a UNO card
type Color uint8

//UNO card colors, Wild represents black cards
const (
	Wild Color = iota
	Red
	Yellow
	Green
	Blue
)

//Value represents the value of a UNO card
type Value uint8

//UNO card values
const (
	Zero Value = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Skip
	Reverse
	DrawTwo  //Draw Two
	DrawFour //Draw Four
	Choose
)

//Card UNO card type
type Card struct {
	Color
	Value
}

var red = color.New(color.BgHiRed).Add(color.FgBlack).SprintFunc()
var yellow = color.New(color.BgHiYellow).Add(color.FgBlack).SprintFunc()
var green = color.New(color.BgHiGreen).Add(color.FgBlack).SprintFunc()
var blue = color.New(color.BgHiBlue).Add(color.FgBlack).SprintFunc()
var black = color.New(color.FgHiBlack).Add(color.FgHiWhite).SprintFunc()

func (c Card) String() string {
	uncolored := fmt.Sprintf("%s %s",
		c.Color.String(),
		c.Value.String())
	switch c.Color {
	case Red:
		return red(uncolored)
	case Yellow:
		return yellow(uncolored)
	case Green:
		return green(uncolored)
	case Blue:
		return blue(uncolored)
	default:
		return black(uncolored)
	}
}

var nonwildcolors = [...]Color{
	Red, Yellow, Green, Blue,
}
var allvalues = [...]Value{
	Zero,
	One, Two, Three, Four, Five, Six, Seven, Eight, Nine,
	Skip, Reverse, DrawTwo,
	DrawFour, Choose,
}
