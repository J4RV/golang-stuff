//go:generate stringer -type=Color,Value -linecomment=true

package unocards

import (
	"fmt"
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

func (c Card) String() string {
	return fmt.Sprintf("%s %s",
		c.Color.String(),
		c.Value.String())
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
