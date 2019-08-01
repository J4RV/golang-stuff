package approximations

import (
	"fmt"
	"math"
)

type sines2 struct {
	// a, b represent one Sine function each
	// 1: the sine's frequency
	// 2: the sine's displacement
	// 3: the sine's amplitude
	y, z, a1, a2, a3, b1, b2, b3 float64
}

func (c *sines2) Calc(x float64) float64 {
	return c.y + x*c.z + math.Sin(x*c.a1+c.a2)*c.a3 + math.Sin(x*c.b1+c.b2)*c.b3
}

func (c *sines2) Mutation(temp float64, out *Cell) {
	res, ok := (*out).(*sines2)
	if !ok {
		panic("Mutation error")
	}
	*res = *c
	(*res).mutate(*mutationArray(8, temp))
	*out = res
}

func (c *sines2) String() string {
	return fmt.Sprintf("%.2f + %.2fx + sin(%.2fx+%.2f)*%.2f + sin(%.2fx+%.2f)*%.2f",
		c.y, c.z,
		c.a1, c.a2, c.a3,
		c.b1, c.b2, c.b3,
	)
}

func (c *sines2) New(cfg Config) Cell {
	new := *c
	new.mutate(*initialArray(8, cfg.initialTemp))
	return &new
}

func (c *sines2) mutate(mutation []float64) {
	c.y += mutation[0]
	c.z += mutation[1]
	c.a1 += mutation[2]
	c.a2 += mutation[3]
	c.a3 += mutation[4]
	c.b1 += mutation[5]
	c.b2 += mutation[6]
	c.b3 += mutation[7]
}
