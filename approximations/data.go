package approximations

import (
	"math/rand"
)

type Cell interface {
	Calc(x float64) float64
	Mutation(temp float64) Cell
	Fitness(points *[]Point) float64
	String() string
	New() Cell
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (p Point) aproxError(c Cell) float64 {
	aprox := c.Calc(p.X)
	diff := p.Y - aprox
	return diff * diff // better than math.Abs(diff)
}

//mutationArray returns an array of zeros and a single element that's randomly in the range [-1.0, +1.0)
func mutationArray(len int) []float64 {
	res := make([]float64, len)
	rng := rand.Float64()

	i := 0
	inc := 1.0 / float64(len)
	for interval := 0.0; interval < 1.0; interval += inc {
		if rng < interval {
			res[i] = (rand.Float64() - .5) * 2
			break
		}
		i++
	}

	return res
}

func initialArray(len int) []float64 {
	res := make([]float64, len)
	for i := 0; i < len; i++ {
		res[i] = (rand.Float64() - .5) * 2
	}
	return res
}
