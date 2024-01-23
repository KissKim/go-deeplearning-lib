package optimizers

type Momentum struct {
	gradients [][]float64
	momentum  float64
}