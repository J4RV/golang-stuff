package main

import (
	"os"

	"github.com/j4rv/gostuff/log"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func plotResult(points []point, c cell) {
	p, err := plot.New()
	checkPlotErr(err)

	p.Title.Text = c.string()
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	min, max := points[0].x, points[len(points)-1].x
	expand := (max - min) * 0.2
	step := (max - min) * 0.01
	err = plotutil.AddLinePoints(p,
		"Points", pointsToPlotter(points),
		"Best approximation", cellToPlotter(min-expand, max+expand, step, c))
	checkPlotErr(err)

	err = p.Save(6*vg.Inch, 6*vg.Inch, "result.png")
	checkPlotErr(err)
}

func pointsToPlotter(p []point) plotter.XYs {
	pts := make(plotter.XYs, len(p))
	for i := range pts {
		pts[i].X = p[i].x
		pts[i].Y = p[i].y
	}
	return pts
}

func cellToPlotter(min, max, step float64, c cell) plotter.XYs {
	pts := plotter.XYs{}
	for i := min; i <= max; i += step {
		p := plotter.XY{X: i, Y: c.calc(i)}
		pts = append(pts, p)
	}
	return pts
}

func checkPlotErr(err error) {
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
