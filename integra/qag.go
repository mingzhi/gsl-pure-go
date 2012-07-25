package integra

import (
	"github.com/mingzhi/gsl"
	"github.com/mingzhi/gsl/err"
	"math"
)

func Qag(f Function, a, b, epsabs, epsrel float64, limit int, workspace *Workspace, q QKFunction) (result, abserr float64, e error) {
	var (
		area, errsum                       float64
		result0, abserr0, resabs0, resasc0 float64
		tolerance                          float64
		iteration                          int
		roundoff_type1                     int
		roundoff_type2                     int
		error_type                         int
		round_off                          float64
	)

	// Initialize results
	workspace.Initialize(a, b)

	result = 0
	abserr = 0

	if limit > workspace.limit {
		e = err.Error{Message: "iteration limit exceeds available workspace", Status: err.EINVAL}
		return
	}

	if epsabs <= 0 && (epsrel < 50*gsl.EpsilonFloat64 || epsrel < 0.5e-28) {
		e = err.Error{Message: "tolerance cannot be acheived with given epsabs and epsrel", Status: err.EBADTOL}
		return
	}

	// perform the first integration
	result0, abserr0, resabs0, resasc0 = q(f, a, b)
	workspace.SetInitialResult(result0, abserr0)

	// Test on accuracy
	tolerance = math.Max(epsabs, epsrel*math.Abs(result0))
	round_off = 50 * gsl.EpsilonFloat64 * resabs0
	if abserr0 <= round_off && abserr0 > tolerance {
		result = result0
		abserr = abserr0
		e = err.Error{Message: "cannot reach tolerance because of roundoff error on first attempt", Status: err.EROUND}
		return
	} else if (abserr0 <= tolerance && abserr0 != resasc0) || abserr0 == 0.0 {
		result = result0
		abserr = abserr0
		return
	} else if limit == 1 {
		result = result0
		abserr = abserr0
		e = err.Error{Message: "a maximum of one iteration was insufficient", Status: err.EMAXITER}
		return
	}

	area = result0
	errsum = abserr0

	iteration = 1
	for iteration < limit && errsum > tolerance && error_type == 0 {
		var (
			a1, b1, a2, b2          float64
			a_i, b_i, r_i, e_i      float64
			area1, area2, area12    float64
			error1, error2, error12 float64
			resasc1, resasc2        float64
		)

		a_i, b_i, r_i, e_i = workspace.Retrieve()
		a1 = a_i
		b1 = 0.5 * (a_i + b_i)
		a2 = b1
		b2 = b_i

		area1, error1, _, resasc1 = q(f, a1, b1)
		area2, error2, _, resasc2 = q(f, a2, b2)

		area12 = area1 + area2
		error12 = error1 + error2

		errsum += (error12 - e_i)
		area += area12 - r_i

		if resasc1 != error1 && resasc2 != error2 {
			delta := r_i - area12

			if math.Abs(delta) <= 1.0e-5*math.Abs(area12) && error12 >= 0.99*e_i {
				roundoff_type1++
			}
			if iteration >= 10 && error12 > e_i {
				roundoff_type2++
			}
		}

		tolerance = math.Max(epsabs, epsrel*math.Abs(area))

		if errsum > tolerance {
			if roundoff_type1 >= 6 || roundoff_type2 >= 20 {
				error_type = 2 /* round off error */
			}

			/* set error flag in the case of bad integrand behaviour at
			   a point of the integration range */

			if subintervalTooSmall(a1, a2, b2) {
				error_type = 3
			}
		}

		workspace.Update(a1, b1, area1, error1, a2, b2, area2, error2)

		a_i, b_i, r_i, e_i = workspace.Retrieve()

		iteration++
	}

	result = workspace.SumResults()
	abserr = errsum

	if errsum <= tolerance {
		return
	} else if error_type == 2 {
		e = err.Error{Message: "roundoff error prevents tolerance from being achieved", Status: err.EROUND}
	} else if error_type == 3 {
		e = err.Error{Message: "bad integrand behavior found in the integration interval", Status: err.ESING}
	} else if iteration == limit {
		e = err.Error{Message: "maximum number of subdivisions reached", Status: err.EMAXITER}
	} else {
		e = err.Error{Message: "could not integrate function", Status: err.EFAILED}
	}

	return
}

func subintervalTooSmall(a1, a2, b2 float64) (status bool) {
	e := gsl.EpsilonFloat64
	u := math.SmallestNonzeroFloat64

	tmp := (1 + 100*e) * (math.Abs(a2) + 1000*u)

	status = math.Abs(a1) <= tmp && math.Abs(b2) <= tmp
	return
}
