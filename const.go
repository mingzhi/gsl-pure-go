package gsl

const FL_EPSILON = 2.2204460492503131e-16
const FL_MIN = 2.2250738585072014e-308

type Function interface {
	Evaluate(x float64) float64
}
