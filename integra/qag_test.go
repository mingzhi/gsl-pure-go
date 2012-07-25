package integra

import (
	"github.com/mingzhi/go-utils/assert"
	"testing"
)

func TestQag(t *testing.T) {
	{
		w, _ := NewWorkspace(1000)
		exp_result := 7.716049382715854665E-02
		exp_abserr := 6.679384885865053037E-12
		exp_last := 6

		a := []float64{0, 0.5, 0.25, 0.125, 0.0625, 0.03125}
		b := []float64{0.03125, 1, 0.5, 0.25, 0.125, 0.0625}
		r := []float64{3.966769831709074375E-06, 5.491842501998222409E-02,
			1.909827770934243926E-02, 2.776531175604360531E-03,
			3.280661030752063693E-04, 3.522704932261797744E-05}
		e := []float64{6.678528276336181873E-12, 6.097169993333454062E-16,
			2.120334764359736934E-16, 3.082568839745514608E-17,
			3.642265412331439511E-18, 3.910988124757650942E-19}
		order := []int{1, 2, 3, 4, 5, 6}

		alpha := 2.6
		f := f1{alpha: alpha}
		result, abserr, err := Qag(f, 0.0, 1.0, 0.0, 1e-10, w.limit, w, Qk15)
		if !assert.EqualFloat(result, exp_result, 1e-15) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, exp_result)
		}
		if !assert.EqualFloat(abserr, exp_abserr, 1e-6) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Smooth: last %d, expected %d\n", w.size, exp_last)
		}
		if err != nil {
			t.Errorf("Smooth: do not expected error %v\n", err)
		}

		for i := 0; i < 6; i++ {
			if !assert.EqualFloat(w.alist[i], a[i], 1e-15) {
				t.Errorf("Smooth: a[%d] %f, expected %f\n", i, w.alist[i], a[i])
			}
			if !assert.EqualFloat(w.blist[i], b[i], 1e-15) {
				t.Errorf("Smooth: b[%d] %f, expected %f\n", i, w.blist[i], b[i])
			}
			if !assert.EqualFloat(w.rlist[i], r[i], 1e-15) {
				t.Errorf("Smooth: r[%d] %f, expected %f\n", i, w.rlist[i], r[i])
			}
			if !assert.EqualFloat(w.elist[i], e[i], 1e-15) {
				t.Errorf("Smooth: e[%d] %f, expected %f\n", i, w.elist[i], e[i])
			}
			if !assert.EqualInt(w.order[i], order[i]-1) {
				t.Errorf("Smooth: order[%d] %d, expected %d\n", i, w.order[i], order[i]-1)
			}
		}

		result, abserr, err = Qag(f, 1.0, 0.0, 0.0, 1e-10, w.limit, w, Qk15)
		if !assert.EqualFloat(result, -exp_result, 1e-15) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, -exp_result)
		}
		if !assert.EqualFloat(abserr, exp_abserr, 1e-6) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Smooth: last %d, expected %d\n", w.size, exp_last)
		}
		if err != nil {
			t.Errorf("Smooth: do not expected error %v\n", err)
		}
	}

	/* Test the same function using an absolute error bound and the
	   21-point rule */

	{
		w, _ := NewWorkspace(1000)
		exp_result := 7.716049382716050342E-02
		exp_abserr := 2.227969521869139532E-15
		exp_last := 8

		a := []float64{0, 0.5, 0.25, 0.125, 0.0625, 0.03125, 0.015625, 0.0078125}
		b := []float64{0.0078125, 1, 0.5, 0.25, 0.125, 0.0625, 0.03125, 0.015625}
		r := []float64{3.696942726831556522E-08, 5.491842501998223103E-02,
			1.909827770934243579E-02, 2.776531175604360097E-03,
			3.280661030752062609E-04, 3.522704932261797744E-05,
			3.579060884684503576E-06, 3.507395216921808047E-07}
		e := []float64{1.371316364034059572E-15, 6.097169993333454062E-16,
			2.120334764359736441E-16, 3.082568839745514608E-17,
			3.642265412331439511E-18, 3.910988124757650460E-19,
			3.973555800712018091E-20, 3.893990926286736620E-21}
		order := []int{1, 2, 3, 4, 5, 6, 7, 8}

		alpha := 2.6
		f := f1{alpha: alpha}
		result, abserr, err := Qag(f, 0.0, 1.0, 1e-14, 0.0, w.limit, w, Qk21)
		if !assert.EqualFloat(result, exp_result, 1e-15) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, exp_result)
		}
		if !assert.EqualFloat(abserr, exp_abserr, 1e-6) {
			t.Errorf("Smooth: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Smooth: last %d, expected %d\n", w.size, exp_last)
		}
		if err != nil {
			t.Errorf("Smooth: do not expected error %v\n", err)
		}

		for i := 0; i < 6; i++ {
			if !assert.EqualFloat(w.alist[i], a[i], 1e-15) {
				t.Errorf("Smooth: a[%d] %f, expected %f\n", i, w.alist[i], a[i])
			}
			if !assert.EqualFloat(w.blist[i], b[i], 1e-15) {
				t.Errorf("Smooth: b[%d] %f, expected %f\n", i, w.blist[i], b[i])
			}
			if !assert.EqualFloat(w.rlist[i], r[i], 1e-15) {
				t.Errorf("Smooth: r[%d] %f, expected %f\n", i, w.rlist[i], r[i])
			}
			if !assert.EqualFloat(w.elist[i], e[i], 1e-6) {
				t.Errorf("Smooth: e[%d] %f, expected %f\n", i, w.elist[i], e[i])
			}
			if !assert.EqualInt(w.order[i], order[i]-1) {
				t.Errorf("Smooth: order[%d] %d, expected %d\n", i, w.order[i], order[i]-1)
			}
		}

		result, abserr, err = Qag(f, 1.0, 0.0, 1e-14, 0.0, w.limit, w, Qk21)
		if !assert.EqualFloat(result, -exp_result, 1e-15) {
			t.Errorf("Smooth: result %.10f, expected %.10f\n", result, -exp_result)
		}
		if !assert.EqualFloat(abserr, exp_abserr, 1e-6) {
			t.Errorf("Smooth: abserr %.20f, expected %.20f\n", abserr, exp_abserr)
		}
		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Smooth: last %d, expected %d\n", w.size, exp_last)
		}
		if err != nil {
			t.Errorf("Smooth: do not expected error %v\n", err)
		}
	}

	/* Adaptive integration of an oscillatory function which terminates because
	   of roundoff error, uses the 31-pt rule */
	{
		w, _ := NewWorkspace(1000)
		exp_result := -7.238969575482959717E-01
		exp_abserr := 1.285805464427459261E-14
		exp_last := 1

		alpha := 1.3
		f := f3{alpha: alpha}
		result, abserr, err := Qag(f, 0.3, 2.71, 1e-14, 0.0, w.limit, w, Qk31)
		if !assert.EqualFloat(result, exp_result, 1e-15) {
			t.Errorf("Oscill: result %.10f, expected %.10f\n", result, exp_result)
		}
		if !assert.EqualFloat(abserr, exp_abserr, 1e-6) {
			t.Errorf("Oscill: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Oscill: last %d, expected %d\n", w.size, exp_last)
		}
		if err == nil {
			t.Errorf("Oscill: expected error.\n")
		}

		result, abserr, err = Qag(f, 2.71, 0.3, 1e-14, 0.0, w.limit, w, Qk31)
		if !assert.EqualFloat(result, -exp_result, 1e-15) {
			t.Errorf("Oscill: result %.10f, expected %.10f\n", result, -exp_result)
		}
		if !assert.EqualFloat(abserr, exp_abserr, 1e-6) {
			t.Errorf("Oscill: abserr %.20f, expected %.20f\n", abserr, exp_abserr)
		}
		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Oscill: last %d, expected %d\n", w.size, exp_last)
		}
		if err == nil {
			t.Errorf("Oscill: expected error.\n")
		}
	}

	/* Check the singularity detection (singularity at x=-0.1 in this example) */
	{
		w, _ := NewWorkspace(1000)
		exp_last := 51

		alpha := 2.0
		f := f16{alpha: alpha}
		_, _, err := Qag(f, -1.0, 1.0, 1e-14, 0.0, w.limit, w, Qk51)

		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Oscill: last %d, expected %d\n", w.size, exp_last)
		}
		if err == nil {
			t.Errorf("Oscill: expected error.\n")
		}

		_, _, err = Qag(f, 1.0, -1.0, 1e-14, 0.0, w.limit, w, Qk51)

		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Oscill: last %d, expected %d\n", w.size, exp_last)
		}
		if err == nil {
			t.Errorf("Oscill: expected error.\n")
		}
	}
	/* Check for hitting the iteration limit */
	{
		w, _ := NewWorkspace(3)
		exp_result := 9.565151449233894709
		exp_abserr := 1.570369823891028460E+01
		exp_last := 3

		a := []float64{-5.000000000000000000E-01,
			0.000000000000000000,
			-1.000000000000000000}
		b := []float64{0.000000000000000000,
			1.000000000000000000,
			-5.000000000000000000E-01}
		r := []float64{9.460353469435913709,
			9.090909090909091161E-02,
			1.388888888888888812E-02}
		e := []float64{1.570369823891028460E+01,
			1.009293658750142399E-15,
			1.541976423090495140E-16}
		order := []int{1, 2, 3}

		alpha := 1.0
		f := f16{alpha: alpha}
		result, abserr, err := Qag(f, -1.0, 1.0, 1e-14, 0.0, w.limit, w, Qk61)
		if !assert.EqualFloat(result, exp_result, 1e-15) {
			t.Errorf("Limit: result %.10f, expected %.10f\n", result, exp_result)
		}
		if !assert.EqualFloat(abserr, exp_abserr, 1e-6) {
			t.Errorf("Limit: abserr %f, expected %f\n", abserr, exp_abserr)
		}
		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Limit: last %d, expected %d\n", w.size, exp_last)
		}
		if err == nil {
			t.Errorf("Limit: expect error\n")
		}

		for i := 0; i < 3; i++ {
			if !assert.EqualFloat(w.alist[i], a[i], 1e-15) {
				t.Errorf("Limit: a[%d] %f, expected %f\n", i, w.alist[i], a[i])
			}
			if !assert.EqualFloat(w.blist[i], b[i], 1e-15) {
				t.Errorf("Limit: b[%d] %f, expected %f\n", i, w.blist[i], b[i])
			}
			if !assert.EqualFloat(w.rlist[i], r[i], 1e-15) {
				t.Errorf("Limit: r[%d] %f, expected %f\n", i, w.rlist[i], r[i])
			}
			if !assert.EqualFloat(w.elist[i], e[i], 1e-6) {
				t.Errorf("Limit: e[%d] %f, expected %f\n", i, w.elist[i], e[i])
			}
			if !assert.EqualInt(w.order[i], order[i]-1) {
				t.Errorf("Limit: order[%d] %d, expected %d\n", i, w.order[i], order[i]-1)
			}
		}

		result, abserr, err = Qag(f, 1.0, -1.0, 1e-14, 0.0, w.limit, w, Qk61)
		if !assert.EqualFloat(result, -exp_result, 1e-15) {
			t.Errorf("Limit: result %.10f, expected %.10f\n", result, -exp_result)
		}
		if !assert.EqualFloat(abserr, exp_abserr, 1e-6) {
			t.Errorf("Limit: abserr %.20f, expected %.20f\n", abserr, exp_abserr)
		}
		if !assert.EqualInt(w.size, exp_last) {
			t.Errorf("Limit: last %d, expected %d\n", w.size, exp_last)
		}
		if err == nil {
			t.Errorf("Limit: expect error\n")
		}
	}
}
