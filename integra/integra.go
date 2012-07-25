package integra

import (
	"math"
)

type Function interface {
	Evaluate(x float64) float64
}

func Evaluate(f Function, x float64) float64 {
	return f.Evaluate(x)
}

type QKFunction func(f Function, a, b float64) (float64, float64, float64, float64)

type float struct {
	value     float64
	precision float64
}

type f1 struct {
	alpha float64
	count int
}

func (f f1) Evaluate(x float64) float64 {
	return math.Pow(x, f.alpha) * math.Log(1.0/x)
}

type f3 struct {
	alpha float64
}

func (f f3) Evaluate(x float64) float64 {
	return math.Cos(math.Pow(2.0, f.alpha) * math.Sin(x))
}

type f16 struct {
	alpha float64
}

func (f f16) Evaluate(x float64) float64 {
	if x == 0 && f.alpha == 1 {
		return 1
	}
	if x == 0 && f.alpha > 1 {
		return 0
	}
	return math.Pow(x, f.alpha-1) / math.Pow(1+10.0*x, 2.0)
}
