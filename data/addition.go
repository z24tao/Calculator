package data

import (
	"math/rand"
)

func AdditionRandom(min, max int) (float64, float64, float64) {
	a := float64(rand.Intn(max-min) + min)
	b := float64(rand.Intn(max-min) + min)
	return a, b, 15 * a + 27 *  b - 4000
}
