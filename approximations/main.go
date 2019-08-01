package approximations

import (
	"errors"
	"math"
	"strings"

	"github.com/j4rv/gostuff/log"
)

// CellType indicates the structure of the function that will be approximated
type CellType int8

const (
	//Cubes f(x) = ? + ?x + ?x^2 + ?x^3
	Cubes CellType = iota
	//Sines2 f(x) = ? + ?x + sin(?x+?)*? + sin(?x+?)*?
	Sines2
	//Sines3 f(x) = ? + ?x + sin(?x+?)*? + sin(?x+?)*? + sin(?x+?)*?
	Sines3
)

func calcInitialTemp(points *[]Point) float64 {
	var min, max float64
	for _, p := range *points {
		if p.Y < min {
			min = p.Y
		}
		if p.Y > max {
			max = p.Y
		}
	}
	return (max - min) / 1000
}

// TypeFromString will return the corresponding CellType from its string identifier
// example: "sines2" (string) -> Sines2 (CellType)
func TypeFromString(s string) (CellType, error) {
	switch strings.ToLower(s) {
	case "cubes":
		return Cubes, nil
	case "sines2":
		return Sines2, nil
	case "sines3":
		return Sines3, nil
	default:
		return Sines3, errors.New(s + " is not a valid cell type. Returned Sines3.")
	}
}

func typeToCell(ct CellType) Cell {
	switch ct {
	case Cubes:
		return &cubes{}
	case Sines2:
		return &sines2{}
	case Sines3:
		return &sines3{}
	default:
		log.Error(ct, "is not a valid function type, using sines3")
		return &sines3{}
	}
}

// CalcBestCell will try to find the function with the celltype structure
// that best approximates the points given
func CalcBestCell(cfg Config, ct CellType, points *[]Point) (candidate Cell, fitting float64) {
	if cfg.initialTemp == 0 {
		cfg.initialTemp = calcInitialTemp(points)
	}

	cell := typeToCell(ct)
	cells := make([]Cell, cfg.population)
	for i := range cells {
		cells[i] = cell.New(cfg)
	}

	candidate, fitting = findBestCandidate(points, &cells)

	for i := 0; i <= cfg.generations; i++ {
		log.Debug("Iteration nº:", i)
		temp := getTemp(cfg, i)
		newGeneration(cfg, temp, candidate, &cells)
		candidate, fitting = findBestCandidate(points, &cells)
	}

	return candidate, fitting
}

// TODO this is a good candidate for goroutines optimization
func newGeneration(cfg Config, temperature float64, best Cell, gens *[]Cell) {
	log.Trace("Adding last best candidate:", best)
	(*gens)[0] = best

	var bigMutations = (cfg.population * cfg.mutationPercentage) / 100
	var smallMutations = cfg.population - bigMutations

	log.Trace("Mutating citizens:", smallMutations-1)
	for i := 1; i < smallMutations; i++ {
		best.Mutation(temperature, &(*gens)[i])
	}

	log.Trace("Mutating with high temperature:", bigMutations)
	for i := smallMutations; i < cfg.population; i++ {
		best.Mutation(cfg.initialTemp, &(*gens)[i])
	}
}

func findBestCandidate(points *[]Point, cells *[]Cell) (Cell, float64) {
	bestCell := (*cells)[0]
	bestFit := fitness(bestCell, points)
	for i := 1; i < len(*cells); i++ {
		c := (*cells)[i]
		fit := fitness(c, points)
		if bestFit == -1 || fit < bestFit {
			bestFit = fit
			bestCell = c
		}
	}
	log.Debug("Best candidate: ", bestCell)
	log.Debug("Fitness: ", bestFit)
	return bestCell, bestFit
}

func getTemp(cfg Config, iteration int) float64 {
	return cfg.initialTemp * math.Exp(-float64(iteration))
}
