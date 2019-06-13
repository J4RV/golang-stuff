package approximations

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"

	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func PlotResult(points *[]Point, path string, cells ...Cell) {
	p, err := plot.New()
	checkPlotErr(err)

	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	min, max := minMaxX(points)
	expand := (max-min)*.2 + 0.00001 // 0.00001 for edge cases where min and max are the same. if that happens, the plot library freezes
	step := (max - min) * 0.002

	// Points plot
	scattered, err := plotter.NewScatter(pointsToPlotter(points))
	checkPlotErr(err)
	scattered.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	scattered.Shape = draw.CircleGlyph{}

	for i, c := range cells {
		// Best approximation plot
		function, err := plotter.NewLine(cellToPlotter(min-expand, max+expand, step, c))
		checkPlotErr(err)
		function.LineStyle.Width = vg.Points(1)
		function.LineStyle.Color = funcColor(i)

		p.Add(function)
		p.Legend.Add(c.String(), function)
	}

	p.Add(scattered)
	p.Legend.Add("points", scattered)

	err = p.Save(10*vg.Inch, 6*vg.Inch, path)
	checkPlotErr(err)
}

func funcColor(i int) color.RGBA {
	switch i {
	case 0:
		return color.RGBA{0, 184, 217, 255} // Tamarama cyan
	case 1:
		return color.RGBA{255, 86, 48, 255} // Poppy surprise red
	case 2:
		return color.RGBA{101, 84, 192, 255} // Da' juice purple
	case 3:
		return color.RGBA{0, 82, 204, 255} // Pacific bridge blue
	case 4:
		return color.RGBA{54, 179, 126, 255} // Fine pine green
	case 5:
		return color.RGBA{255, 171, 0, 255} // Golden state yellow
	default:
		return randColor()
	}
}

// make it more predictable, should never return colors close to white
func randColor() color.RGBA {
	return color.RGBA{
		uint8(rand.Float32() * 256),
		uint8(rand.Float32() * 256),
		uint8(rand.Float32() * 256),
		255,
	}
}

func minMaxX(points *[]Point) (float64, float64) {
	min, max := (*points)[0].X, (*points)[0].X
	for _, p := range *points {
		if p.X < min {
			min = p.X
		}
		if p.X > max {
			max = p.X
		}
	}
	return min, max
}

func pointsToPlotter(p *[]Point) plotter.XYs {
	pts := make(plotter.XYs, len(*p))
	for i := range pts {
		pts[i].X = (*p)[i].X
		pts[i].Y = (*p)[i].Y
	}
	return pts
}

func cellToPlotter(min, max, step float64, c Cell) plotter.XYs {
	pts := plotter.XYs{}
	for i := min; i <= max; i += step {
		p := plotter.XY{X: i, Y: c.Calc(i)}
		pts = append(pts, p)
	}
	return pts
}

func checkPlotErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
