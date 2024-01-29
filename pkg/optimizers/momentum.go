package optimizers

type Momentum struct {
	gradients [][]float64
	momentum  float64
}

func (m Momentum) Apply(weights [][]float64) Momentum {
	gradients := make([][]float64, len(weights))
	for i := range gradients {
		gradie