package layers

import (
	"fmt"
	"math"
	"strings"
)

type UnbiasedDense struct {
	layer
	learner
	gradi