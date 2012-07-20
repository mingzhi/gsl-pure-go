package stat

import (
	"github.com/mingzhi/go-utils/assert"
	"math/rand"
	"testing"
)

type TestingFloats struct {
	raw, group               []float64
	stride, n                int
	mean                     float64
	variance_with_fixed_mean float64
	variance                 float64
	sd                       float64
	sd_with_fixed_mean       float64
	tss                      float64
	tss_m                    float64
	absdev                   float64
}

type TestingInts struct {
	raw, group               []int
	stride, n                int
	mean                     float64
	variance                 float64
	variance_with_fixed_mean float64
	sd                       float64
	sd_with_fixed_mean       float64
	tss                      float64
	tss_m                    float64
	absdev                   float64
}

var (
	f_groupa TestingFloats
	f_groupb TestingFloats
	i_groupa TestingInts
)

var (
	delta64 float64
)

func init() {
	delta64 = 1e-10
	stride := rand.Intn(10) + 1
	f_groupa = TestingFloats{
		raw: []float64{
			.0421, .0941, .1064, .0242, .1331,
			.0773, .0243, .0815, .1186, .0356,
			.0728, .0999, .0614, .0479,
		},
		mean: 0.0728,
		variance_with_fixed_mean: 0.00113837428571429,
		sd:                       0.0350134479659107,
		sd_with_fixed_mean:       0.0337398026922845,
		tss_m:                    1.59372400000000e-02,
		absdev:                   0.0287571428571429,
		n:                        14,
		stride:                   stride,
	}

	f_groupa.group = make([]float64, f_groupa.n*f_groupa.stride)
	for i, item := range f_groupa.raw {
		f_groupa.group[i*stride] = item
	}

	f_groupb = TestingFloats{
		raw: []float64{
			.1081, .0986, .1566, .1961, .1125,
			.1942, .1079, .1021, .1583, .1673,
			.1675, .1856, .1688, .1512,
		},
		variance: 0.00124956615384615,
		tss:      0.01624436,
		n:        14,
		stride:   stride,
	}
	f_groupb.group = make([]float64, f_groupb.n*f_groupa.stride)
	for i, item := range f_groupb.raw {
		f_groupb.group[i*stride] = item
	}

	i_groupa = TestingInts{
		raw: []int{
			17, 18, 16, 18, 12,
			20, 18, 20, 20, 22,
			20, 10, 8, 12, 16,
			16, 18, 20, 18, 21,
		},
		n:      20,
		stride: stride,
		mean:   17.0,
		variance_with_fixed_mean: 13.7,
		variance:                 14.4210526315789,
		sd_with_fixed_mean:       3.70135110466435,
		sd:                       3.79750610685209,
		absdev:                   2.9,
	}
	i_groupa.group = make([]int, i_groupa.n*stride)
	for i, item := range i_groupa.raw {
		i_groupa.group[i*stride] = item
	}
}

func TestMeanFloats(t *testing.T) {
	mean := MeanFloats(f_groupa.group, f_groupa.stride, f_groupa.n)
	correct := assert.EqualFloat(mean, f_groupa.mean, delta64)
	if !correct {
		t.Errorf("mean: %f, expected: %f\n", mean, f_groupa.mean)
	}
}

func TestMeanInts(t *testing.T) {
	mean := MeanInts(i_groupa.group, i_groupa.stride, i_groupa.n)
	correct := assert.EqualFloat(mean, i_groupa.mean, delta64)
	if !correct {
		t.Errorf("mean: %f, expected: %f\n", mean, i_groupa.mean)
	}
}

func TestMean(t *testing.T) {
	var mean float64
	var correct bool
	// test float
	mean = Mean(i_groupa.group, i_groupa.stride, i_groupa.n)
	correct = assert.EqualFloat(mean, i_groupa.mean, delta64)
	if !correct {
		t.Errorf("Int mean: %f, expected: %f\n", mean, i_groupa.mean)
	}
	// test int
	mean = Mean(f_groupa.group, f_groupa.stride, f_groupa.n)
	correct = assert.EqualFloat(mean, f_groupa.mean, delta64)
	if !correct {
		t.Errorf("Float mean: %f, expected: %f\n", mean, f_groupa.mean)
	}
	// test string
	data := []string{"1", "2", "3", "4"}
	stride := 1
	n := len(data)
	mean = Mean(data, stride, n)
	correct = assert.EqualFloat(mean, 2.5, delta64)
	if !correct {
		t.Errorf("String mean: %f, expected: %f\n", mean, 2.5)
	}
}

func TestVarianceWithFixedMean(t *testing.T) {
	var mean, variance float64
	var correct bool
	// test float
	mean = Mean(f_groupa.group, f_groupa.stride, f_groupa.n)
	variance = VarianceWithFixedMean(f_groupa.group, f_groupa.stride, f_groupa.n, mean)
	correct = assert.EqualFloat(variance, f_groupa.variance_with_fixed_mean, delta64)
	if !correct {
		t.Errorf("variance: %f, expected: %f\n", variance, f_groupa.variance)
	}
	// test int
	mean = Mean(i_groupa.group, i_groupa.stride, i_groupa.n)
	variance = VarianceWithFixedMean(i_groupa.group, i_groupa.stride, i_groupa.n, mean)
	correct = assert.EqualFloat(variance, i_groupa.variance_with_fixed_mean, delta64)
	if !correct {
		t.Errorf("variance_with_fixed_mean: %f, expected %f\n", variance, i_groupa.variance_with_fixed_mean)
	}
}

func TestVariance(t *testing.T) {
	var variance float64
	var correct bool
	// test float
	variance = Variance(f_groupb.group, f_groupb.stride, f_groupb.n)
	correct = assert.EqualFloat(variance, f_groupb.variance, delta64)
	if !correct {
		t.Errorf("variance: %f, expected: %f\n", variance, f_groupb.variance)
	}
	// test int
	variance = Variance(i_groupa.group, i_groupa.stride, i_groupa.n)
	correct = assert.EqualFloat(variance, i_groupa.variance, delta64)
	if !correct {
		t.Errorf("variance: %f, expected: %f\n", variance, i_groupa.variance)
	}
}

func TestSdWithFixedMean(t *testing.T) {
	var mean, sd float64
	var correct bool
	// test float
	mean = Mean(f_groupa.group, f_groupa.stride, f_groupa.n)
	sd = SdWithFixedMean(f_groupa.group, f_groupa.stride, f_groupa.n, mean)
	correct = assert.EqualFloat(sd, f_groupa.sd_with_fixed_mean, delta64)
	if !correct {
		t.Errorf("sd_with_fixed_mean: %f, but expected: %f\n", sd, f_groupa.sd_with_fixed_mean)
	}
	// test int
	mean = Mean(i_groupa.group, i_groupa.stride, i_groupa.n)
	sd = SdWithFixedMean(i_groupa.group, i_groupa.stride, i_groupa.n, mean)
	correct = assert.EqualFloat(sd, i_groupa.sd_with_fixed_mean, delta64)
	if !correct {
		t.Errorf("sd_with_fixed_mean: %f, but expected: %f\n", sd, i_groupa.sd_with_fixed_mean)
	}
}

func TestSd(t *testing.T) {
	var sd float64
	var correct bool
	// test float
	sd = Sd(f_groupa.group, f_groupa.stride, f_groupa.n)
	correct = assert.EqualFloat(sd, f_groupa.sd, delta64)
	if !correct {
		t.Errorf("sd: %f, expected: %f\n", sd, f_groupa.sd)
	}
	// test int
	sd = Sd(i_groupa.group, i_groupa.stride, i_groupa.n)
	correct = assert.EqualFloat(sd, i_groupa.sd, delta64)
	if !correct {
		t.Errorf("sd: %f, expected: %f\n", sd, i_groupa.sd)
	}
}

func TestTss(t *testing.T) {
	var tss float64
	var correct bool
	// test float
	tss = Tss(f_groupb.group, f_groupb.stride, f_groupb.n)
	correct = assert.EqualFloat(tss, f_groupb.tss, delta64)
	if !correct {
		t.Errorf("tss: %f, expected: %f\n", tss, f_groupb.tss)
	}
}

func TestTssM(t *testing.T) {
	mean := MeanFloats(f_groupa.group, f_groupa.stride, f_groupa.n)
	tss := TssM(f_groupa.group, f_groupa.stride, f_groupa.n, mean)
	correct := assert.EqualFloat(tss, f_groupa.tss_m, delta64)
	if !correct {
		t.Errorf("tss_m: %f, expected: %f\n", tss, f_groupa.tss_m)
	}
}

func TestAbsdev(t *testing.T) {
	var absdev float64
	var correct bool
	// test float
	absdev = Absdev(f_groupa.group, f_groupa.stride, f_groupa.n)
	correct = assert.EqualFloat(absdev, f_groupa.absdev, delta64)
	if !correct {
		t.Errorf("absdev: %f, expected: %f\n", absdev, f_groupa.absdev)
	}
	// test int
	absdev = Absdev(i_groupa.group, i_groupa.stride, i_groupa.n)
	correct = assert.EqualFloat(absdev, i_groupa.absdev, delta64)
	if !correct {
		t.Errorf("absdev: %f, expected: %f\n", absdev, i_groupa.absdev)
	}
}
