package layers

type Bias struct {
	layer
	learner
}

func (b *Bias) Estimate(input []float64) []float64 {
	f