package optimizers

import "math"

type Adam struct {
	epsilon   float64
	