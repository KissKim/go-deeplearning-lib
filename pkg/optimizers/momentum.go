package optimizers

type Momentum struct {
	gradients [][]float64
	momentum  float64
}

func (m Momentum) Apply(weights [][]float64) Momen