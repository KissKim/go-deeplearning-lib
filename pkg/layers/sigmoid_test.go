package layers

import (
	"math"
	"strconv"
	"testing"
)

func TestSigmoidActivate(t *testing.T) {
	var sigmoid = func