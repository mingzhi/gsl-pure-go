package integration

import (
	"github.com/mingzhi/gsl"
	"math"
)

type IntegrateError struct {
	message string
}

func (err IntegrateError) Error() string {
	return err.message
}

func RescaleError(reserr, resabs, resasc float64) float64 {
	reserr = math.Abs(reserr)
	if resasc != 0 && reserr != 0 {
		scale := math.Pow((200 * reserr / resasc), 1.5)

		if scale < 1 {
			reserr = resasc * scale
		} else {
			reserr = resasc
		}
	}
	if resabs > gsl.FL_MIN/(50*gsl.FL_EPSILON) {
		min_err := 50 * gsl.FL_EPSILON * resabs

		if min_err > reserr {
			reserr = min_err
		}
	}
	return reserr
}

/*
func RescaleError(err, resabs, resasc float64) float64 {
	err = math.Abs(err)
	if resasc != 0 && err != 0 {
		scale := math.Pow((200 * err / resasc), 1.5)

		if scale < 1 {
			err = resasc * scale
		} else {
			err = resasc
		}
	}
	if resabs > gsl.FL_MIN/(50*gsl.FL_EPSILON) {
		min_err := 50 * gsl.FL_EPSILON * resabs

		if min_err > err {
			err = min_err
		}
	}
	return err
}
*/
