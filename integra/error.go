package integra

import (
	"github.com/mingzhi/gsl"
	"math"
)

type IError struct {
	message string
}

func (err IError) Error() string {
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
	if resabs > math.SmallestNonzeroFloat64/(50*gsl.EpsilonFloat64) {
		min_err := 50 * gsl.EpsilonFloat64 * resabs

		if min_err > reserr {
			reserr = min_err
		}
	}
	return reserr
}
