package stat

import (
	"github.com/mingzhi/go-utils/number"
	"reflect"
)

func MeanFloats(data []float64, stride, n int) (mean float64) {
	for i := 0; i < n; i++ {
		mean += (data[i*stride] - mean) / float64(i+1)
	}
	return
}

func MeanInts(data []int, stride, n int) (mean float64) {
	for i := 0; i < n; i++ {
		mean += (float64(data[i*stride]) - mean) / float64(i+1)
	}
	return
}

func Mean(data interface{}, stride, n int) (mean float64) {
	arry := reflect.ValueOf(data)
	for i := 0; i < n; i++ {
		v := arry.Index(i * stride)
		fv := number.Float(v)
		mean += (fv - mean) / float64(i+1)
	}
	return
}
