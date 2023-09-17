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
	Neurons 