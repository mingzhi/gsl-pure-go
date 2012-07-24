package integra

import (
	"math"
)

type Function interface {
	Evaluate(x float64) float64
}

type QKFunction func(f Function, a, b float64) (float64, float64, float64, float64)

type Float struct {
	value     float64
	precision float64
}

type F1 struct {
	alpha float64
}

func (f F1) Evaluate(x float64) float64 {
	return math.Pow(x, f.alpha) * math.Log(1.0/x)
}

type F3 struct {
	alpha float64
}

func (f F3) Evaluate(x float64) float64 {
	return math.Cos(math.Pow(2.0, f.alpha) * math.Sin(x))
}
