package approximations

import (
	"errors"

	"github.com/j4rv/gostuff/log"
)

var (
	ErrPopulationNotValid         = errors.New("Population cannot be less than 1")
	ErrMutationPercentageNotValid = errors.New("Mutation chance must be in the range [0, 100]")
	ErrGenerationsNotValid        = errors.New("Generations cannot be less than 1")
	ErrInitialTempNotValid        = errors.New("Generations cannot be less than 1")
)

type Config struct {
	population         int
	mutationPercentage int
	generations        int
	initialTemp        float64
}

func NewConfig(population, mutationPercentage, generations int) (Config, error) {
	if population < 1 {
		return Config{}, ErrPopulationNotValid
	}
	if !(0 < mutationPercentage || mutationPercentage < 100) {
		return Config{}, ErrMutationPercentageNotValid
	}
	if generations < 1 {
		return Config{}, ErrGenerationsNotValid
	}
	return Config{
		population:         population,
		mutationPercentage: mutationPercentage,
		generations:        generations,
	}, nil
}

func SetInitialTemp(cfg *Config, temp float64) error {
	if temp <= 0 {
		return ErrInitialTempNotValid
	}
	(*cfg).initialTemp = temp
	return nil
}

func SetLogLevel(l log.Level) {
	log.SetLevel(l)
}
