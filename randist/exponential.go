package randist

import (
	"math"
	"math/rand"
)

type Exponential struct {
	mu float64
}

func (e Exponential) Rand() (x float64) {
	return ExponentialRand(e.mu)
}

func (e Exponential) Pdf(x float64) (p float64) {
	return ExponentialPdf(x, e.mu)
}

func (e Exponential) Evaluate(x float64) (p float64) {
	return e.Pdf(x)
}

// The exponential distribution has the form
// p(x) dx = exp(-x/mu) dx/mu
// for x = 0 ... + infty
func ExponentialRand(mu float64) (x float64) {
	u := rand.Float64()
	x = -mu * math.Log1p(-u)
	return
}

func ExponentialPdf(x, mu float64) (p float64) {
	if x < 0 {
		p = 0.0
	} else {
		p = math.Exp(-x/mu) / mu
	}
	return
}
