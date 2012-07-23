package randist

import (
	"math"
	"math/rand"
)

// The exponential distribution has the form
// p(x) dx = exp(-x/mu) dx/mu
// for x = 0 ... + infty
func RandExponential(r *rand.Rand, mu float64) (x float64) {
	u := r.Float64()
	x = -mu * math.Log1p(-u)
	return
}

func RandExponentialPdf(x, mu float64) (p float64) {
	if x < 0 {
		p = 0.0
	} else {
		p = math.Exp(-x/mu) / mu
	}
	return
}
