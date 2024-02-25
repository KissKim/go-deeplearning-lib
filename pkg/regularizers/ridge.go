package regularizers

import "math"

type Ridge struct {
	Lambda  float64
	w