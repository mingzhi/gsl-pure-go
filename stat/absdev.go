package stat

import (
	"math"
)

// Takes a dataset and finds the absolute deviation with a fixed mean
func AbsdevM(data []float64, stride, n int, mean float64) (absdev float64) {
	sum := 0.0
	for i := 0; i < n; i++ {
		delta := math.Abs(data[i*stride] - mean)
		sum += delta
	}
	absdev = sum / float64(n)
	return
}

// Takes a dataset and finds the absolute deviation
func Absdev(data []float64, stride, n int) (absdev float64) {
	mean := MeanFloats(data, stride, n)
	absdev = AbsdevM(data, stride, n, mean)
	return
}
