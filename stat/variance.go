package stat

import (
	"github.com/mingzhi/go-utils/number"
	"math"
	"reflect"
)

// Take a dataset and finds the variance
func computeVariance(data interface{}, stride, n int, mean float64) (variance float64) {
	arry := reflect.ValueOf(data)
	// find the sum of the squares
	for i := 0; i < n; i++ {
		v := arry.Index(i * stride)
		f := number.Float(v)
		delta := f - mean
		variance += (delta*delta - variance) / float64(i+1)
	}
	return
}

// Takes a dataset and finds the sum of squares about the mean
func computeTss(data interface{}, stride, n int, mean float64) (tss float64) {
	arry := reflect.ValueOf(data)
	// find the sum of the squares
	for i := 0; i < n; i++ {
		v := arry.Index(i * stride)
		f := number.Float(v)
		delta := f - mean
		tss += delta * delta
	}
	return
}

// Take a dataset and finds the variance giving a fixed mean
func VarianceWithFixedMean(data interface{}, stride, n int, mean float64) (variance float64) {
	variance = computeVariance(data, stride, n, mean)
	return
}

func SdWithFixedMean(data interface{}, stride, n int, mean float64) (sd float64) {
	variance := computeVariance(data, stride, n, mean)
	sd = math.Sqrt(variance)
	return
}

func VarianceM(data interface{}, stride, n int, mean float64) (variance float64) {
	variance = computeVariance(data, stride, n, mean)
	variance = variance * (float64(n) / float64(n-1))
	return
}

func SdM(data interface{}, stride, n int, mean float64) (sd float64) {
	variance := computeVariance(data, stride, n, mean)
	sd = math.Sqrt(variance * (float64(n) / float64(n-1)))
	return
}

func Variance(data interface{}, stride, n int) (variance float64) {
	mean := Mean(data, stride, n)
	variance = VarianceM(data, stride, n, mean)
	return
}

func Sd(data interface{}, stride, n int) (sd float64) {
	mean := Mean(data, stride, n)
	sd = SdM(data, stride, n, mean)
	return
}

func TssM(data interface{}, stride, n int, mean float64) (tss float64) {
	tss = computeTss(data, stride, n, mean)
	return
}

func Tss(data interface{}, stride, n int) (tss float64) {
	mean := Mean(data, stride, n)
	tss = TssM(data, stride, n, mean)
	return
}
