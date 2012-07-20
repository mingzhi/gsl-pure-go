package stat

import (
	"github.com/mingzhi/go-utils/number"
	"reflect"
)

func MedianFromSortedData(sorted_data interface{}, stride, n int) (median float64) {
	arry := reflect.ValueOf(sorted_data)
	lhs := (n - 1) / 2
	rhs := n / 2
	if n == 0 {
		median = 0.0
	} else {
		if lhs == rhs {
			median = number.Float(arry.Index(lhs * stride))
		} else {
			median = (number.Float(arry.Index(lhs*stride)) + number.Float(arry.Index(rhs*stride))) / 2.0
		}
	}
	return
}
