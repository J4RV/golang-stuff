package approximations

import (
	"fmt"
)

type cubes struct {
	a, b, c, d float64
}

func (c *cubes) Calc(x float64) float64 {
	sq := x * x
	return c.a + c.b*x + c.c*sq + c.d*x*sq
}

func (c *cubes) Mutation(temp float64, out *Cell) {
	res, ok := (*out).(*cubes)
	if !ok {
		panic("Mutation error")
	}
	*res = *c
	(*res).mutate(*mutationArray(4, temp))
	*out = res
}

func (c *cubes) String() string {
	return fmt.Sprintf("%.3f + x*%.3f + (x^2)*%.3f + (x^3)*%.3f\n", c.a, c.b, c.c, c.d)
}

func (c *cubes) New(cfg Config) Cell {
	res := *c
	res.mutate(*initialArray(4, cfg.initialTemp))
	return &res
}

func (c *cubes) mutate(mutation []float64) {
	c.a += mutation[0]
	c.b += mutation[1]
	c.c += mutation[2]
	c.d += mutation[3]
}
