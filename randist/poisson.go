package randist

import (
	"math"
	"math/rand"
)

func Poisson(mu float64) (k int) {
	prod := 1.0

	emu := math.Exp(-mu)
	for {
		prod *= rand.Float64()
		k++
		if prod <= emu {
			break
		}
	}
	k = k - 1
	return
}
