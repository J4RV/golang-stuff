package main

import (
	"encoding/json"
	"flag"
	"math/rand"
	"os"
	"time"

	. "github.com/j4rv/gostuff/approximations"
	"github.com/j4rv/gostuff/log"
	"github.com/j4rv/gostuff/stopwatch"
)

const defaultPoints = `[
	{"x":0,"y":0},
	{"x":1,"y":100},
	{"x":2,"y":160},
	{"x":3,"y":200},
	{"x":3.3,"y":190},
	{"x":4,"y":220},
	{"x":5,"y":150},
	{"x":6,"y":310},
	{"x":7,"y":428}
]`

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	params := parseFlagParams()
	points := params.points
	log.Info("Points: ", points)

	log.Info("Approximating best function...")
	stop := stopwatch.Start()
	bestCell, bestFit := CalcBestCell(params.cellType, points)
	elapsed := stop()

	log.Info("Done in seconds:", elapsed.Seconds())
	log.Info("Best approximation: ", bestCell)
	log.Info("Best fit: ", bestFit)
	log.Info(bestCell.String())

	bestCubes, _ := CalcBestCell(Cubes, points)
	bestSines2, _ := CalcBestCell(Sines2, points)

	PlotResult(points, params.resPath, bestCell, bestSines2, bestCubes)
}

type flagParams struct {
	cellType CellType
	points   *[]Point
	resPath  string
}

func parseFlagParams() flagParams {
	var res flagParams

	typeFlag := flag.String("type", "sines3", "The type of function to approximate")
	pointsFlag := flag.String("points", defaultPoints, "The cloud of points to approximate")
	resPathFlag := flag.String("resPath", "./result.png", "Where the result image will be saved")
	flag.Parse()

	log.Info("Using function type:", *typeFlag)
	t, err := TypeFromString(*typeFlag)
	if err != nil {
		log.Error(err.Error())
	}
	res.cellType = t

	points := pointsFromJSON(*pointsFlag)
	res.points = points

	res.resPath = *resPathFlag

	return res
}

func pointsFromJSON(s string) *[]Point {
	points := make([]Point, 0)
	err := json.Unmarshal([]byte(s), &points)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	return &points
}
