package gsl

const EpsilonFloat64 = 2.2204460492503131e-16

type Function interface {
	Evaluate(x float64) float64
}
