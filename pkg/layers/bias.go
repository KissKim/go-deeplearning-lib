package layers

type Bias struct {
	layer
	learner
}

func (b *Bias) Estimate(input []fl