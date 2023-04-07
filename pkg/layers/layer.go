package layers

type layer struct {
	inputShape  Shape
	outputShape Shape
	output      []float64
}

func (l *layer) Minimize(loss []float64) []float64 {
	return loss
}

func (l *layer) Shape() []uint64 {
	return l.outputShape.Shape()
}

func (l *layer) SetShape(sha