package randist

type ProbabilityDistribution interface {
	Pdf(float64) float64
	Rand() float64
	Evaluate(float64) float64
}

type DiscreteDistribution interface {
	Pdf(int) float64
	Rand() int
	Evaluate(int) float64
}
