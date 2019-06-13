package approximations

import (
	"errors"
	"math"

	"github.com/j4rv/gostuff/log"
)

const (
	population      = 20000
	mutationsPerGen = 16000
	generations     = 10
)

type CellType int8

const (
	//Cubes f(x) = ? + ?x + ?x^2 + ?x^3
	Cubes CellType = iota
	//Sines2 f(x) = ? + ?x + sin(?x+?)*? + sin(?x+?)*?
	Sines2
	//Sines3 f(x) = ? + ?x + sin(?x+?)*? + sin(?x+?)*? + sin(?x+?)*?
	Sines3
)

var initialTemp float64

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
	return max - min
}

func TypeFromString(s string) (CellType, error) {
	switch s {
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

func CalcBestCell(ct CellType, points *[]Point) (candidate Cell, fitting float64) {
	cell := typeToCell(ct)
	cells := make([]Cell, population)
	for i := range cells {
		cells[i] = cell.New()
	}

	candidate, fitting = findBestCandidate(points, &cells)

	for i := 0; i <= generations; i++ {
		log.Debug("Iteration nº:", i)
		temp := getTemp(i)
		newGeneration(temp, candidate, &cells)
		candidate, fitting = findBestCandidate(points, &cells)
	}

	return candidate, fitting
}

func newGeneration(temperature float64, best Cell, gens *[]Cell) {
	log.Trace("Adding last best candidate:", best)
	(*gens)[0] = best

	log.Trace("Mutating citizens:", mutationsPerGen-1)
	for i := 1; i < mutationsPerGen; i++ {
		(*gens)[i] = best.Mutation(temperature)
	}

	log.Trace("Mutating with high temperature:", population-mutationsPerGen)
	for i := mutationsPerGen; i < population; i++ {
		(*gens)[i] = best.Mutation(initialTemp)
	}
}

func findBestCandidate(points *[]Point, cells *[]Cell) (Cell, float64) {
	bestCell := (*cells)[0]
	bestFit := bestCell.Fitness(points)
	for i := 1; i < len(*cells); i++ {
		c := (*cells)[i]
		fit := c.Fitness(points)
		if bestFit == -1 || fit < bestFit {
			bestFit = fit
			bestCell = c
		}
	}
	log.Debug("Best candidate: ", bestCell)
	log.Debug("Fitness: ", bestFit)
	return bestCell, bestFit
}

func getTemp(iteration int) float64 {
	return initialTemp * math.Exp(-float64(iteration))
}
