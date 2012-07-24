package integration

import (
	"github.com/mingzhi/go-utils/assert"
	"math"
	"testing"
)

type F1 struct {
	alpha float64
}

func (f F1) Evaluate(x float64) float64 {
	return math.Pow(x, f.alpha) * math.Log(1.0/x)
}

func TestQng(t *testing.T) {
	var (
		exp_result float64
		exp_abserr float64
		exp_neval  int
		result     float64
		abserr     float64
		neval      int
		err        error
	)
	exp_result = 7.716049379303083211E-02
	exp_abserr = 9.424302199601294244E-08
	exp_neval = 21
	f := F1{alpha: 2.6}
	result, abserr, neval, err = Qng(f, 0.0, 1.0, 1e-1, 0.0)
	if err != nil {
		t.Errorf("Don't expect Error: %v\n", err)
	}
	if !assert.EqualFloat(result, exp_result, 1e-15) {
		t.Errorf("Result: %f, expect: %f\n", result, exp_result)
	}
	if !assert.EqualFloat(abserr, exp_abserr, 1e-7) {
		t.Errorf("Abserr: %f, expect: %f\n", abserr, exp_abserr)
	}
	if neval != exp_neval {
		t.Errorf("Neval: %f, expect: %f\n", neval, exp_neval)
	}

	result, abserr, neval, err = Qng(f, 1.0, 0.0, 1e-1, 0.0)
	if err != nil {
		t.Errorf("Don't expect Error: %v\n", err)
	}
	if !assert.EqualFloat(result, -exp_result, 1e-15) {
		t.Errorf("Result: %f, expect: %f\n", result, exp_result)
	}
	if !assert.EqualFloat(abserr, exp_abserr, 1e-7) {
		t.Errorf("Abserr: %f, expect: %f\n", abserr, exp_abserr)
	}
	if neval != exp_neval {
		t.Errorf("Neval: %f, expect: %f\n", neval, exp_neval)
	}

	exp_result = 7.716049382706505200E-02
	exp_abserr = 2.666893044866214501E-12
	exp_neval = 43
	f2 := F1{alpha: 2.6}

	result, abserr, neval, err = Qng(f2, 0.0, 1.0, 1e-9, 0.0)
	if err != nil {
		t.Errorf("Don't expect Error: %v\n", err)
	}
	if !assert.EqualFloat(result, exp_result, 1e-15) {
		t.Errorf("Result: %f, expect: %f\n", result, exp_result)
	}
	if !assert.EqualFloat(abserr, exp_abserr, 1e-5) {
		t.Errorf("Abserr: %f, expect: %f\n", abserr, exp_abserr)
	}
	if neval != exp_neval {
		t.Errorf("Neval: %d, expect: %d\n", neval, exp_neval)
	}

	exp_result = -7.238969575482961938E-01
	exp_abserr = 1.277676889520056369E-14
	exp_neval = 43
	f3 := F1{alpha: 1.3}

	result, abserr, neval, err = Qng(f3, 0.3, 2.71, 0.0, 1e-12)
	if err != nil {
		t.Errorf("Don't expect Error: %v\n", err)
	}
	if !assert.EqualFloat(result, exp_result, 1e-15) {
		t.Errorf("Result: %f, expect: %f\n", result, exp_result)
	}
	if !assert.EqualFloat(abserr, exp_abserr, 1e-7) {
		t.Errorf("Abserr: %f, expect: %f\n", abserr, exp_abserr)
	}
	if neval != exp_neval {
		t.Errorf("Neval: %d, expect: %d\n", neval, exp_neval)
	}
}
