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
	fgroupa TestingFloats
	fgroupb TestingFloats
	igroupa TestingInts
	igroupb TestingInts
)

var (
	delta64 float64
)

func init() {
	delta64 = 1e-10
	stride := rand.Intn(10) + 1
	// float group A
	fgroupa = TestingFloats{
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

	fgroupa.group = make([]float64, fgroupa.n*fgroupa.stride)
	for i, item := range fgroupa.raw {
		fgroupa.group[i*stride] = item
	}

	// float group B
	fgroupb = TestingFloats{
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
	fgroupb.group = make([]float64, fgroupb.n*fgroupa.stride)
	for i, item := range fgroupb.raw {
		fgroupb.group[i*stride] = item
	}

	// int group A
	igroupa = TestingInts{
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
	igroupa.group = make([]int, igroupa.n*stride)
	for i, item := range igroupa.raw {
		igroupa.group[i*stride] = item
	}

	// int group B
	igroupb = TestingInts{
		raw: []int{
			19, 20, 22, 24, 10,
			25, 20, 22, 21, 23,
			20, 10, 12, 14, 12,
			20, 22, 24, 23, 17,
		},
		n:      20,
		stride: stride,
	}
	igroupb.group = make([]int, igroupb.n*stride)
	for i, item := range igroupb.raw {
		igroupb.group[i*stride] = item
	}
}

func TestMeanFloats(t *testing.T) {
	mean := MeanFloats(fgroupa.group, fgroupa.stride, fgroupa.n)
	correct := assert.EqualFloat(mean, fgroupa.mean, delta64)
	if !correct {
		t.Errorf("mean: %f, expected: %f\n", mean, fgroupa.mean)
	}
}

func TestMeanInts(t *testing.T) {
	mean := MeanInts(igroupa.group, igroupa.stride, igroupa.n)
	correct := assert.EqualFloat(mean, igroupa.mean, delta64)
	if !correct {
		t.Errorf("mean: %f, expected: %f\n", mean, igroupa.mean)
	}
}

func TestMean(t *testing.T) {
	var mean float64
	var correct bool
	// test float
	mean = Mean(igroupa.group, igroupa.stride, igroupa.n)
	correct = assert.EqualFloat(mean, igroupa.mean, delta64)
	if !correct {
		t.Errorf("Int mean: %f, expected: %f\n", mean, igroupa.mean)
	}
	// test int
	mean = Mean(fgroupa.group, fgroupa.stride, fgroupa.n)
	correct = assert.EqualFloat(mean, fgroupa.mean, delta64)
	if !correct {
		t.Errorf("Float mean: %f, expected: %f\n", mean, fgroupa.mean)
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
	mean = Mean(fgroupa.group, fgroupa.stride, fgroupa.n)
	variance = VarianceWithFixedMean(fgroupa.group, fgroupa.stride, fgroupa.n, mean)
	correct = assert.EqualFloat(variance, fgroupa.variance_with_fixed_mean, delta64)
	if !correct {
		t.Errorf("variance: %f, expected: %f\n", variance, fgroupa.variance)
	}
	// test int
	mean = Mean(igroupa.group, igroupa.stride, igroupa.n)
	variance = VarianceWithFixedMean(igroupa.group, igroupa.stride, igroupa.n, mean)
	correct = assert.EqualFloat(variance, igroupa.variance_with_fixed_mean, delta64)
	if !correct {
		t.Errorf("variance_with_fixed_mean: %f, expected %f\n", variance, igroupa.variance_with_fixed_mean)
	}
}

func TestVariance(t *testing.T) {
	var variance float64
	var correct bool
	// test float
	variance = Variance(fgroupb.group, fgroupb.stride, fgroupb.n)
	correct = assert.EqualFloat(variance, fgroupb.variance, delta64)
	if !correct {
		t.Errorf("variance: %f, expected: %f\n", variance, fgroupb.variance)
	}
	// test int
	variance = Variance(igroupa.group, igroupa.stride, igroupa.n)
	correct = assert.EqualFloat(variance, igroupa.variance, delta64)
	if !correct {
		t.Errorf("variance: %f, expected: %f\n", variance, igroupa.variance)
	}
}

func TestSdWithFixedMean(t *testing.T) {
	var mean, sd float64
	var correct bool
	// test float
	mean = Mean(fgroupa.group, fgroupa.stride, fgroupa.n)
	sd = SdWithFixedMean(fgroupa.group, fgroupa.stride, fgroupa.n, mean)
	correct = assert.EqualFloat(sd, fgroupa.sd_with_fixed_mean, delta64)
	if !correct {
		t.Errorf("sd_with_fixed_mean: %f, but expected: %f\n", sd, fgroupa.sd_with_fixed_mean)
	}
	// test int
	mean = Mean(igroupa.group, igroupa.stride, igroupa.n)
	sd = SdWithFixedMean(igroupa.group, igroupa.stride, igroupa.n, mean)
	correct = assert.EqualFloat(sd, igroupa.sd_with_fixed_mean, delta64)
	if !correct {
		t.Errorf("sd_with_fixed_mean: %f, but expected: %f\n", sd, igroupa.sd_with_fixed_mean)
	}
}

func TestSd(t *testing.T) {
	var sd float64
	var correct bool
	// test float
	sd = Sd(fgroupa.group, fgroupa.stride, fgroupa.n)
	correct = assert.EqualFloat(sd, fgroupa.sd, delta64)
	if !correct {
		t.Errorf("sd: %f, expected: %f\n", sd, fgroupa.sd)
	}
	// test int
	sd = Sd(igroupa.group, igroupa.stride, igroupa.n)
	correct = assert.EqualFloat(sd, igroupa.sd, delta64)
	if !correct {
		t.Errorf("sd: %f, expected: %f\n", sd, igroupa.sd)
	}
}

func TestTss(t *testing.T) {
	var tss float64
	var correct bool
	// test float
	tss = Tss(fgroupb.group, fgroupb.stride, fgroupb.n)
	correct = assert.EqualFloat(tss, fgroupb.tss, delta64)
	if !correct {
		t.Errorf("tss: %f, expected: %f\n", tss, fgroupb.tss)
	}
}

func TestTssM(t *testing.T) {
	mean := MeanFloats(fgroupa.group, fgroupa.stride, fgroupa.n)
	tss := TssM(fgroupa.group, fgroupa.stride, fgroupa.n, mean)
	correct := assert.EqualFloat(tss, fgroupa.tss_m, delta64)
	if !correct {
		t.Errorf("tss_m: %f, expected: %f\n", tss, fgroupa.tss_m)
	}
}

func TestAbsdev(t *testing.T) {
	var absdev float64
	var correct bool
	// test float
	absdev = Absdev(fgroupa.group, fgroupa.stride, fgroupa.n)
	correct = assert.EqualFloat(absdev, fgroupa.absdev, delta64)
	if !correct {
		t.Errorf("absdev: %f, expected: %f\n", absdev, fgroupa.absdev)
	}
	// test int
	absdev = Absdev(igroupa.group, igroupa.stride, igroupa.n)
	correct = assert.EqualFloat(absdev, igroupa.absdev, delta64)
	if !correct {
		t.Errorf("absdev: %f, expected: %f\n", absdev, igroupa.absdev)
	}
}

func TestCovariance(t *testing.T) {
	var covariance, expected float64
	var correct bool
	// test float
	covariance = Covariance(fgroupa.group, fgroupa.stride, fgroupb.group, fgroupb.stride, fgroupb.n)
	expected = -0.000139021538461539
	correct = assert.EqualFloat(covariance, expected, delta64)
	if !correct {
		t.Errorf("covariance: %f, expected: %f\n", covariance, expected)
	}

	// test int
	covariance = Covariance(igroupa.group, igroupa.stride, igroupb.group, igroupb.stride, igroupb.n)
	expected = 14.5263157894737
	correct = assert.EqualFloat(covariance, expected, delta64)
	if !correct {
		t.Errorf("covariance: %f, expected: %f\n", covariance, expected)
	}
}

func TestCorrelation(t *testing.T) {
	var correlation, expected float64
	var correct bool
	// test float
	correlation = Correlation(fgroupa.group, fgroupa.stride, fgroupb.group, fgroupb.stride, fgroupb.n)
	expected = -0.112322712666074171
	correct = assert.EqualFloat(correlation, expected, delta64)
	if !correct {
		t.Errorf("correlation: %f, expected: %f\n", correlation, expected)
	}

	// test int
	correlation = Correlation(igroupa.group, igroupa.stride, igroupb.group, igroupb.stride, igroupb.n)
	expected = 0.793090350710101
	correct = assert.EqualFloat(correlation, expected, delta64)
	if !correct {
		t.Errorf("correlation: %f, expected: %f\n", correlation, expected)
	}
}
