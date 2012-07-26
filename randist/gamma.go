package randist

import (
	"math"
	"math/rand"
)

func GammaInt(a int) (x float64) {
	if a < 12 {
		prod := 1.0
		for i := 0; i < a; i++ {
			r := rand.Float64()
			for r == 0 {
				r = rand.Float64()
			}
			prod *= r
		}
		x = -math.Log(prod)
	} else {
		x = GammaLarge(a)
	}
	return
}

// Works only if a > 1, and is most efficient if a is large
// This algorithm, reported in Knuth, is attributed to Ahrens.  A
// faster one, we are told, can be found in: J. H. Ahrens and 
// U. Dieter, Computing 12 (1974) 223-246.
func GammaLarge(a int) float64 {
	var sqa, x, y, v float64
	sqa = math.Sqrt(2*float64(a) - 1)
	for {
		for {
			y = math.Tan(math.Pi * rand.Float64())
			x = sqa*y + float64(a) - 1
			if x > 0 {
				break
			}
		}
		v = rand.Float64()
		if v <= (1+y*y)*(math.Exp(float64(a)-1)*math.Log(x/(float64(a)-1))-sqa*y) {
			break
		}
	}
	return x
}
