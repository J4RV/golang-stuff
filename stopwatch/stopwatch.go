package stopwatch

import "time"

// Start the stopwatch.
// Returns a function that stops the clockwatch and returns the interval.
func Start() func() time.Duration {
	before := time.Now()
	return func() time.Duration {
		after := time.Now()
		interval := after.Sub(before)
		return interval
	}
}
