package maze

import (
	"math/rand"
	"time"
)

var defaultGenerator = new(randomGenerator)

// randomGeneratorI is an interface for randomGenerator
type randomGeneratorI interface {
	random(min int, max int) int
	random50() bool
}

// randomGenerator contains set of methods to generate random values
type randomGenerator struct{}

// random returns a random number from min to max
func (generator randomGenerator) random(min int, max int) int {
	if min == max {
		return min
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// random50 returns true with 50% probability. Otherwise returns false
func (generator randomGenerator) random50() bool {
	return generator.random(1, 100) < 50
}
