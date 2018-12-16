package cards

import (
	"fmt"
)

func ExampleCard_String() {
	fmt.Println(Card{Color: Red, Value: Reverse})
	fmt.Println(Card{Color: Yellow, Value: One})
	fmt.Println(Card{Color: Green, Value: Skip})
	fmt.Println(Card{Color: Blue, Value: Nine})
	fmt.Println(Card{Color: Wild, Value: DrawFour})
	fmt.Println(Card{Color: Wild, Value: Choose})

	// Output:
	// Red Reverse
	// Yellow One
	// Green Skip
	// Blue Nine
	// Wild Draw Four
	// Wild Choose
}
