package main

import (
	"runtime"
	"sync"

	"github.com/j4rv/gostuff/log"
	"github.com/j4rv/gostuff/stopwatch"

	"math/rand"
)

const (
	population      = 200
	mutationsPerGen = 150
	generations     = 60000
	initialTemp     = 10
	finalTemp       = 0.01
)

func main() {
	log.SetLevelFromFlag("log")
	points := getPoints()
	log.Debug("Points: ", points)

	log.Info("Approximating best function...")
	stop := stopwatch.Start()
	bestCell, bestFit := calcBestCell(&points)
	elapsed := stop()

	log.Info("Done in seconds:", elapsed.Seconds())
	log.Info("Best approximation: ", bestCell)
	log.Info("Best fit: ", bestFit)
	log.Info(bestCell.string())

	plotResult(points, bestCell)
}

// Failed attempt at trying to make it better
func multiApproximate(points *[]point) (candidate cell, fitting float64) {
	n := runtime.NumCPU()
	cells := make(chan cell, n)

	// multiple goroutines making their approximations
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			log.Info("Starting worker", i)
			c, _ := calcBestCell(points)
			cells <- c
			wg.Done()
			log.Info("Finished worker", i)
		}(i)
	}
	wg.Wait()
	close(cells)

	// init candidate and fitting using the first candidate cell available
	c := <-cells
	candidate, fitting = c, calcFit(points, c)

	// choose the best approximation by checking the rest of the candidates
	for c := range cells {
		f := calcFit(points, c)
		if f < fitting {
			candidate, fitting = c, f
		}
	}

	return candidate, fitting
}

func calcBestCell(points *[]point) (candidate cell, fitting float64) {
	cells := make([]cell, population)
	for i := range cells {
		cells[i] = newRandomCell()
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

func newGeneration(temperature float64, best cell, gens *[]cell) {
	log.Trace("Adding last best candidate:", best)
	(*gens)[0] = best

	log.Trace("Mutating citizens:", mutationsPerGen-1)
	for i := 1; i < mutationsPerGen; i++ {
		(*gens)[i] = mutateCell(best, temperature)
	}

	log.Trace("Replacing citizens with new ones:", population-mutationsPerGen)
	for i := mutationsPerGen; i < population; i++ {
		(*gens)[i] = newRandomCell()
	}
}

func findBestCandidate(points *[]point, cells *[]cell) (cell, float64) {
	bestCell := (*cells)[0]
	bestFit := calcFit(points, bestCell)
	for i := 1; i < len(*cells); i++ {
		c := (*cells)[i]
		fit := calcFit(points, c)
		if bestFit == -1 || fit < bestFit {
			bestFit = fit
			bestCell = c
		}
	}
	log.Debug("Best candidate: ", bestCell)
	log.Debug("Fitness: ", bestFit)
	return bestCell, bestFit
}

func newRandomCell() cell {
	return mutateCell(cell{}, initialTemp)
}

func getTemp(iteration int) float64 {
	pos := float64(generations-iteration) / float64(generations)
	return initialTemp*(pos) + finalTemp*(1-pos)
}

func mutateCell(c cell, temp float64) cell {
	return cell{
		c.a + (rand.Float64()-.5)*temp,
		c.b + (rand.Float64()-.5)*temp,
		c.c + (rand.Float64()-.5)*temp,
		c.d + (rand.Float64()-.5)*temp,
	}
}

func calcFit(points *[]point, c cell) float64 {
	var sum float64
	for _, p := range *points {
		sum += p.calcFit(c)
	}
	return sum
}
