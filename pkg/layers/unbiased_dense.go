package layers

import (
	"fmt"
	"math"
	"strings"
)

type UnbiasedDense struct {
	layer
	learner
	gradients []float64
	input     []float64
	Neurons   uint64
}

func (d *UnbiasedDense) Estimate(input []float64) []float64 {
	d.input = input
	for j := range d.weights {
		var z float64
		for k := range d.weights[j] {
			z = math.FMA(d.weights[j][k], input[k], z)
		}
		d.output[j] = z
	}
	return d.output
}

func (d *UnbiasedDense) Minimize(gradients []float64) []float64 {
	for k := range d.gradients {
		d.gradients[k] = 0
	}
	for j := range d.weights {
		for k := range d.weights[j] {
			d.localGradients[j][k] = gradients[j] * d.input[k]
			d.gradients[k] += gradients[j] * d.weights[j][k]
		}
	}
	return d.gradients
}

func (d *UnbiasedDense) SetShape(shape []uint64) {
	d.inputShape = shape
	d.outputShap