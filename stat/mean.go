package stat

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
