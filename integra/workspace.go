package integra

type Workspace struct {
	limit         int
	size          int
	nrmax         int
	i             int
	maximum_level int
	alist         []float64
	blist         []float64
	rlist         []float64
	elist         []float64
	order         []int
	level         []int
}

func NewWorkspace(n int) (w *Workspace, err error) {
	if n == 0 {
		err = IError{message: "workspace length n must be positive integer."}
		return
	}

	w = &Workspace{
		size:          0,
		limit:         n,
		maximum_level: 0,
		alist:         make([]float64, n),
		blist:         make([]float64, n),
		elist:         make([]float64, n),
		rlist:         make([]float64, n),
		order:         make([]int, n),
		level:         make([]int, n),
	}

	return
}
