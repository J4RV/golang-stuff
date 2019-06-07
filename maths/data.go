package main

import (
	"fmt"
	"math/rand"
	"time"
)

type point struct {
	x, y float64
}

func (p point) calcFit(c cell) float64 {
	aprox := c.calc(p.x)
	diff := p.y - aprox
	return diff * diff // better than math.Abs(diff)
}

func getPoints() []point {
	// Ordered pls!, from left to right
	rand.Seed(time.Now().Unix())
	var res []point
	for i := 0; i < 8; i++ {
		res = append(res, point{float64(i), rand.Float64() * 100})
		//res = append(res, point{float64(i), float64(1 + 2*i + 3*i*i + 4*i*i*i)})
		//res = append(res, point{float64(i), float64(2*i + 5)})
	}
	return res
}

type cell struct {
	a, b, c, d float64
}

func (c cell) calc(x float64) float64 {
	sq := x * x
	return c.a + c.b*x + c.c*sq + c.d*x*sq
}

func (c cell) string() string {
	return fmt.Sprintf("%.3f + (%.3f)x + (%.3f)x^2 + (%.3f)x^3\n", c.a, c.b, c.c, c.d)
}
