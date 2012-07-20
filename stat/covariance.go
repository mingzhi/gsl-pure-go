package stat

import (
	"github.com/mingzhi/go-utils/number"
	"math"
	"reflect"
)

func computeCovariance(data1 interface{}, stride1 int, data2 interface{}, stride2 int, n int, mean1, mean2 float64) (covariance float64) {
	arry1 := reflect.ValueOf(data1)
	arry2 := reflect.ValueOf(data2)
	// find the sum of the square
	for i := 0; i < n; i++ {
		v1 := number.Float(arry1.Index(i * stride1))
		v2 := number.Float(arry2.Index(i * stride2))
		delta1 := v1 - mean1
		delta2 := v2 - mean2
		covariance += (delta1*delta2 - covariance) / float64(i+1)
	}
	return
}

func CovarianceM(data1 interface{}, stride1 int, data2 interface{}, stride2 int, n int, mean1, mean2 float64) (covariance float64) {
	covariance = computeCovariance(data1, stride1, data2, stride2, n, mean1, mean2)
	covariance = covariance * (float64(n) / float64(n-1))
	return
}

func Covariance(data1 interface{}, stride1 int, data2 interface{}, stride2 int, n int) (covariance float64) {
	mean1 := Mean(data1, stride1, n)
	mean2 := Mean(data2, stride2, n)
	covariance = CovarianceM(data1, stride1, data2, stride2, n, mean1, mean2)
	return
}

func Correlation(data1 interface{}, stride1 int, data2 interface{}, stride2 int, n int) (r float64) {
	var ratio, delta_x, delta_y, sum_xsq, sum_ysq, sum_cross, mean_x, mean_y float64

	arry1 := reflect.ValueOf(data1)
	arry2 := reflect.ValueOf(data2)

	/*
	 Compute:
	 sum_xsq = Sum [ (x_i - mu_x)^2 ],
	 sum_ysq = Sum [ (y_i - mu_y)^2 ] and
	 sum_cross = Sum [ (x_i - mu_x) * (y_i - mu_y) ]
	 using the above relation from Welford's paper
	*/

	mean_x = number.Float(arry1.Index(0 * stride1))
	mean_y = number.Float(arry2.Index(0 * stride2))

	for i := 0; i < n; i++ {
		v1 := number.Float(arry1.Index(i * stride1))
		v2 := number.Float(arry2.Index(i * stride2))
		ratio = float64(i) / float64(i+1)
		delta_x = v1 - mean_x
		delta_y = v2 - mean_y
		sum_xsq += delta_x * delta_x * ratio
		sum_ysq += delta_y * delta_y * ratio
		sum_cross += delta_x * delta_y * ratio
		mean_x += delta_x / float64(i+1)
		mean_y += delta_y / float64(i+1)
	}

	r = sum_cross / (math.Sqrt(sum_xsq) * math.Sqrt(sum_ysq))

	return
}
