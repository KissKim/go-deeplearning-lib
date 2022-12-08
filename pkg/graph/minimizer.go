package graph

type Minimizer struct {
	Layer
	cursor       int
	batch        [][][]float64
	gradients    [][]float64
	learningRate LearningRate
	optimizer    Optimizer
	regularizer  Regularizer
	weights      [][]float64
}

func (m *Minimizer) Minimize(gradients []float64) []float64 {
	gradients = m.Layer.Minimize(gradients)
	for i := range