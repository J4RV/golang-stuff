package hmackeyfinder

import (
	"math/rand"
	"time"
)

//Shuffle shuffles block indexes
func Shuffle(slice [256]uint8) [256]uint8 {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := len(slice); i > 0; i-- {
		randi := r.Intn(i)
		slice[i-1], slice[randi] = slice[randi], slice[i-1]
	}
	return slice
}
