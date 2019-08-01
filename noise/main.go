package main

import (
	"fmt"
	"math"
	"math/rand"
)

type sinArgs struct {
	amplitude, freq float64
}

type noiser struct {
	sins *[]sinArgs
}

func (n *noiser) calc(x float64) float64 {
	var res float64
	for _, args := range *n.sins {
		res += math.Sin(x * args.freq) * args.amplitude
	}
	return res
}

func (n *noiser) String() string {
	var res string
	first := true
	for _, args := range *n.sins {
		if first {
			first = false
		} else {
			res += " + "
		}
		res += fmt.Sprintf("sin(x * %f) * %f", args.freq, args.amplitude)
	}
	return res
}

func newNoiser(n int) *noiser {
	var sins []sinArgs
	for i := 1.0; i <= float64(n); i+=1.0 {
		iCubed := i*i*i
		a := (rand.Float64()) /iCubed
		f := (rand.Float64()) *iCubed
		sins = append(sins, sinArgs{a, f})
	}
	return &noiser{&sins}
}

func main() {
	noiser := newNoiser(9)
	/*for i:=0.0; i<10.0; i+=0.1 {
		fmt.Printf("%.3f -> %.3f \n", i, noiser.calc(i))
	}*/
	fmt.Printf("%s \n", noiser.String())
}