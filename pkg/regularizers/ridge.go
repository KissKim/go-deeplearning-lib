package regularizers

import "math"

type Ridge struct {
	Lambda  float64
	weights [][]float64
}

func