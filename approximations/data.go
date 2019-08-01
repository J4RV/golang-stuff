package approximations

import (
	"math/rand"
	"runtime"
	"sync"
)

// Cell TODO DOCS
type Cell interface {
	Calc(x float64) float64
	Mutation(temp float64, out *Cell)
	String() string
	New(cfg Config) Cell
}

// Point TODO DOCS
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (p Point) aproxError(c Cell) float64 {
	aprox := c.Calc(p.X)
	diff := p.Y - aprox
	return diff * diff // better than math.Abs(diff)
}

var arrayCache = &[]float64{}

//mutationArray returns an array of zeros and a single element that's randomly in the range [-1.0, +1.0)
//not thread safe!!!
func mutationArray(resLen int, temp float64) *[]float64 {
	if resLen != len(*arrayCache) {
		newSlice := make([]float64, resLen)
		arrayCache = &newSlice
	}

	rng := rand.Float64()

	i := 0
	inc := 1.0 / float64(resLen)
	for interval := 0.0; interval < 1.0; interval += inc {
		if rng < interval {
			(*arrayCache)[i] = (rand.Float64() - .5) * 2 * temp
		} else {
			(*arrayCache)[i] = 0
		}
		i++
	}

	return arrayCache
}

func initialArray(resLen int, temp float64) *[]float64 {
	if resLen != len(*arrayCache) {
		newSlice := make([]float64, resLen)
		arrayCache = &newSlice
	}
	for i := 0; i < resLen; i++ {
		(*arrayCache)[i] = (rand.Float64() - .5) * 2 * temp
	}
	return arrayCache
}

func fitness(c Cell, points *[]Point) float64 {
	var sum float64

	// If there are "many" points, do it in parallel
	if len(*points) >= 2048 {
		return parallelFitness(c, points)
	}

	for _, p := range *points {
		sum += p.aproxError(c)
	}

	return sum
}

func parallelFitness(c Cell, points *[]Point) float64 {
	var sum float64
	numCPU := runtime.NumCPU()
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	chunkSize := (len(*points) + numCPU - 1) / numCPU

	for i := 0; i < len(*points); i += chunkSize {
		wg.Add(1)
		end := i + chunkSize
		if end > len(*points) {
			end = len(*points)
		}
		go func(start, end int) {
			var localSum float64
			for _, p := range (*points)[start:end] {
				localSum += p.aproxError(c)
			}
			mutex.Lock()
			sum += localSum
			mutex.Unlock()
			wg.Done()
		}(i, end)
	}

	wg.Wait()
	return sum
}
