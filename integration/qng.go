package integration

import (
	"github.com/mingzhi/gsl"
	"math"
)

func Qng(f gsl.Function, a, b, epsabs, epsrel float64) (result, abserr float64, neval int, err error) {
	var (
		fv1, fv2, fv3, fv4         [5]float64
		savfun                     [21]float64 // array of function values which have been computed
		res10, res21, res43, res87 float64
		result_kronrod             float64
		resabs                     float64
		resasc                     float64
		reserr                     float64
	)

	half_length := 0.5 * (b - a)
	abs_half_length := math.Abs(half_length)
	center := 0.5 * (a + b)
	f_center := f.Evaluate(center)

	if epsabs <= 0 && (epsrel < 0.5e-28 || epsrel < 50*gsl.FL_EPSILON) {
		result = 0.0
		abserr = 0.0
		neval = 0
		err = IntegrateError{message: "tolerance cannot be acheived with given epsabs and epsrel"}
		return
	}

	// compute the integral using the 10- and 21-point formula

	res10 = 0
	res21 = w21b[5] * f_center
	resabs = w21b[5] * math.Abs(f_center)

	for k := 0; k < 5; k++ {
		abscissa := half_length * x1[k]
		fval1 := f.Evaluate(center + abscissa)
		fval2 := f.Evaluate(center - abscissa)
		fval := fval1 + fval2
		res10 += w10[k] * fval
		res21 += w21a[k] * fval
		resabs += w21a[k] * (math.Abs(fval1) + math.Abs(fval2))
		savfun[k] = fval
		fv1[k] = fval1
		fv2[k] = fval2
	}

	for k := 0; k < 5; k++ {
		abscissa := half_length * x2[k]
		fval1 := f.Evaluate(center + abscissa)
		fval2 := f.Evaluate(center - abscissa)
		fval := fval1 + fval2
		res21 += w21b[k] * fval
		resabs += w21b[k] * (math.Abs(fval1) + math.Abs(fval2))
		savfun[k+5] = fval
		fv3[k] = fval1
		fv4[k] = fval2
	}

	resabs *= abs_half_length

	mean := 0.5 * res21
	resasc = w21b[5] * math.Abs(f_center-mean)
	for k := 0; k < 5; k++ {
		resasc += (w21a[k]*(math.Abs(fv1[k]-mean)+math.Abs(fv2[k]-mean)) + w21b[k]*(math.Abs(fv3[k]-mean)+math.Abs(fv4[k]-mean)))
	}
	resasc *= abs_half_length

	result_kronrod = res21 * half_length

	reserr = RescaleError((res21-res10)*half_length, resabs, resasc)

	/*   test for convergence. */

	if reserr < epsabs || reserr < epsrel*math.Abs(result_kronrod) {
		result = result_kronrod
		abserr = reserr
		neval = 21
		return
	}

	/* compute the integral using the 43-point formula. */

	res43 = w43b[11] * f_center

	for k := 0; k < 10; k++ {
		res43 += savfun[k] * w43a[k]
	}

	for k := 0; k < 11; k++ {
		abscissa := half_length * x3[k]
		fval := f.Evaluate(center+abscissa) + f.Evaluate(center-abscissa)
		res43 += fval * w43b[k]
		savfun[k+10] = fval
	}

	/*  test for convergence */

	result_kronrod = res43 * half_length
	reserr = RescaleError((res43-res21)*half_length, resabs, resasc)

	if reserr < epsabs || reserr < epsrel*math.Abs(result_kronrod) {
		result = result_kronrod
		abserr = reserr
		neval = 43
		return
	}

	/* compute the integral using the 87-point formula. */

	res87 = w87b[22] * f_center

	for k := 0; k < 21; k++ {
		res87 += savfun[k] * w87a[k]
	}

	for k := 0; k < 22; k++ {
		abscissa := half_length * x4[k]
		res87 += w87b[k]*f.Evaluate(center+abscissa) + f.Evaluate(center-abscissa)
	}

	/*  test for convergence */

	result_kronrod = res87 * half_length

	reserr = RescaleError((res87-res43)*half_length, resabs, resasc)

	if reserr < epsabs || reserr < epsrel*math.Abs(result_kronrod) {
		result = result_kronrod
		abserr = reserr
		neval = 87
		return
	}

	/* failed to converge */

	result = result_kronrod
	abserr = reserr
	neval = 87
	err = IntegrateError{message: "failed to reach tolerance with highest-order rule"}

	return
}
