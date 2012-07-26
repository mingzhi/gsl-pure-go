package randist

import (
	"fmt"
	"github.com/mingzhi/gsl/integra"
	"math"
	"testing"
)

func TestExponential(t *testing.T) {
	rd := Exponential{mu: 2.0}
	name := "Exponential"
	err := testPDF(rd, name)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

const (
	N    = 100000
	BINS = 100
)

type TestError struct {
	message string
}

func (e TestError) Error() string {
	return e.message
}

func integrate(pdf integra.Function, a, b float64) float64 {
	n := 1000
	w, _ := integra.NewWorkspace(n)
	result, _, _ := integra.Qag(pdf, a, b, 1e-16, 1e-4, n, w, integra.Qk61)
	return result
}

func testPDF(rd ProbabilityDistribution, name string) (err error) {
	count := make([]float64, BINS)
	edge := make([]float64, BINS)
	p := make([]float64, BINS)

	a := -5.0
	b := +5.0
	dx := (b - a) / BINS
	var total, mean float64
	var attempts int
	n0 := 0
	n := N

	for i := 0; i < BINS; i++ {
		/* Compute the integral of p(x) from x to x+dx */

		x := a + float64(i)*dx

		if math.Abs(x) < 1e-10 {
			/* hit the origin exactly */
			x = 0.0
		}

		p[i] = integrate(rd, x, x+dx)
	}

trial:
	attempts++
	for i := n0; i < n; i++ {
		r := rd.Rand()
		total += r
		if a < r && r < b {
			u := (r - a) / dx
			bin, f := math.Modf(u)
			j := int(bin)
			if f == 0 {
				edge[j]++
			} else {
				count[j]++
			}
		}
	}

	/* Sort out where the hits on the edges should go */

	for i := 0; i < BINS; i++ {
		/* If the bin above is empty, its lower edge hits belong in the
		   lower bin */

		if i+1 < BINS && count[i+1] == 0 {
			count[i] += edge[i+1]
			edge[i+1] = 0
		}

		count[i] += edge[i]
		edge[i] = 0
	}

	mean = (total / float64(n))

	exception := math.IsInf(mean, 0) || math.IsNaN(mean)
	if exception {
		err = TestError{message: fmt.Sprintf("%s, finite mean, observed %g", name, mean)}
		return
	}
	exception_i := false
	for i := 0; i < BINS; i++ {
		x := a + float64(i)*dx
		d := math.Abs(count[i] - float64(n)*p[i])
		if math.IsNaN(p[i]) || math.IsInf(p[i], 0) {
			exception_i = true
		} else if p[i] != 0 {
			s := d / math.Sqrt(float64(n)*p[i])
			exception_i = (s > 5) && (d > 2)
		} else {
			exception_i = (count[i] != 0)
		}

		/* Extend the sample if there is an outlier on the first attempt
		   to avoid spurious failures when running large numbers of tests. */
		if exception_i && attempts < 10 {
			n0 = n
			n = 2.0 * n
			goto trial
		}

		exception = exception || exception_i
		if exception_i {
			err = TestError{message: fmt.Sprintf("%s [%g,%g) (%g/%d=%g observed vs %g expected)",
				name, x, x+dx, count[i], n, count[i]/float64(n), p[i])}
		}
	}

	if exception {
		err = TestError{message: fmt.Sprintf("%s, sampling against pdf over range [%g,%g) ",
			name, a, b)}
	}
	return
}
