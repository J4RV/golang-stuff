package approximations

import (
	"fmt"
	"math"
)

type sines3 struct {
	// a, b and c represent one Sine function each
	// 1: the sine's frequency
	// 2: the sine's displacement
	// 3: the sine's amplitude
	y, z, a1, a2, a3, b1, b2, b3, c1, c2, c3 float64
}

func (c *sines3) Calc(x float64) float64 {
	return c.y + x*c.z + math.Sin(x*c.a1+c.a2)*c.a3 + math.Sin(x*c.b1+c.b2)*c.b3 + math.Sin(x*c.c1+c.c2)*c.c3
}

func (c *sines3) Mutation(temp float64) Cell {
	new := *c
	new.mutate(mutationArray(11))
	return &new
}

func (c *sines3) Fitness(points *[]Point) float64 {
	var sum float64
	for _, p := range *points {
		sum += p.aproxError(c)
	}
	return sum
}

func (c *sines3) String() string {
	return fmt.Sprintf("%.2f + %.2fx + sin(%.2fx+%.2f)*%.2f + sin(%.2fx+%.2f)*%.2f + sin(%.2fx+%.2f)*%.2f",
		c.y, c.z,
		c.a1, c.a2, c.a3,
		c.b1, c.b2, c.b3,
		c.c1, c.c2, c.c3,
	)
}

func (c *sines3) New() Cell {
	new := *c
	new.mutate(initialArray(11))
	return &new
}

func (c *sines3) mutate(mutation []float64) {
	c.y += mutation[0]
	c.z += mutation[1]
	c.a1 += mutation[2]
	c.a2 += mutation[3]
	c.a3 += mutation[4]
	c.b1 += mutation[5]
	c.b2 += mutation[6]
	c.b3 += mutation[7]
	c.c1 += mutation[8]
	c.c2 += mutation[9]
	c.c3 += mutation[10]
}
