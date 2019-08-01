package main

import (
	"encoding/json"
	"flag"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	. "github.com/j4rv/gostuff/approximations"
	"github.com/j4rv/gostuff/log"
	"github.com/j4rv/gostuff/stopwatch"
)

const defaultPoints = `[
	{"x":0,  "y":0  },
	{"x":0.2,"y":30 },
	{"x":1,  "y":100},
	{"x":1.5,"y":118},
	{"x":2,  "y":160},
	{"x":2.3,"y":182},
	{"x":2.5,"y":180},
	{"x":2.6,"y":189},
	{"x":3,  "y":200},
	{"x":3.1,"y":199},
	{"x":3.3,"y":190},
	{"x":3.4,"y":146},
	{"x":4,  "y":220},
	{"x":4.8,"y":230},
	{"x":5,  "y":150},
	{"x":5.2,"y":160},
	{"x":5.3,"y":184},
	{"x":5.6,"y":151},
	{"x":6,  "y":310},	
	{"x":6.9,"y":410},
	{"x":7,  "y":428}
]`

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write mem profile to file")

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	params := parseFlagParams()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	cfg := configFromFlags(params)

	points := params.points

	log.Info("Approximating best function...")
	stop := stopwatch.Start()
	bestCell, bestFit := CalcBestCell(cfg, params.cellType, points)
	elapsed := stop()
	log.Info("Done in seconds:", elapsed.Seconds())

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		_ = pprof.WriteHeapProfile(f)
		_ = f.Close()
		return
	}

	log.Info("Best approximation: ", bestCell)
	log.Info("Best fit: ", bestFit)
	log.Info(bestCell.String())

	PlotResult(points, params.resPath, bestCell)
}

type flagParams struct {
	generations        int
	population         int
	mutationPercentage int
	initialTemp        float64

	cellType CellType
	points   *[]Point
	resPath  string

	verbose bool
}

func parseFlagParams() flagParams {
	var res flagParams

	typeFlag := flag.String("type", "sines3", "The type of function to approximate")
	pointsFlag := flag.String("points", defaultPoints, "The cloud of points to approximate")
	flag.StringVar(&res.resPath, "resPath", "./result.png", "Where the result image will be saved")
	flag.IntVar(&res.generations, "generations", 500, "Must be higher than 1")
	flag.IntVar(&res.population, "population", 20000, "Must be higher than 1")
	flag.IntVar(&res.mutationPercentage, "mutationPercentage", 20, "Must be in range [0, 100]")
	flag.Float64Var(&res.initialTemp, "initialTemp", 0,
		`Represents the initial 'temperature'.
	Higher values means the <?> numbers will start with higher absolute values, but can also cause overfitting.
	Must be higher than Zero. If not setted, it will be automatically calculated based on the points.`)
	flag.BoolVar(&res.verbose, "v", false, "Verbose")
	flag.Parse()

	t, err := TypeFromString(*typeFlag)
	if err != nil {
		log.Error(err.Error())
	}
	res.cellType = t

	points := pointsFromJSON(*pointsFlag)
	res.points = points

	return res
}

func randPoints(amount int) *[]Point {
	res := make([]Point, amount)
	f := func(x float64) float64 {
		return x
	}
	for i := 0; i < amount; i++ {
		x := rand.Float64() + float64(i)
		y := f(x) + (rand.Float64()-0.5)*0.1
		res[i] = Point{X: x, Y: y}
	}
	return &res
}

func configFromFlags(fp flagParams) Config {
	if fp.verbose {
		SetLogLevel(log.ALL)
	}
	cfg, err := NewConfig(fp.population, fp.mutationPercentage, fp.generations)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	SetInitialTemp(&cfg, fp.initialTemp)
	return cfg
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
