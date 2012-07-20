package stat

import (
	"github.com/mingzhi/go-utils/number"
	"math"
	"reflect"
)

// Takes a dataset and finds the absolute deviation with a fixed mean
func AbsdevM(data interface{}, stride, n int, mean float64) (absdev float64) {
	arry := reflect.ValueOf(data)
	sum := 0.0
	for i := 0; i < n; i++ {
		v := arry.Index(i * stride)
		f := number.Float(v)
		delta := math.Abs(f - mean)
		sum += delta
	}
	absdev = sum / float64(n)
	return
}

// Takes a dataset and finds the absolute deviation
func Absdev(data interface{}, stride, n int) (absdev float64) {
	mean := Mean(data, stride, n)
	absdev = AbsdevM(data, stride, n, mean)
	return
}
