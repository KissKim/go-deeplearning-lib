package graph

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestFeaturesClassWeights(t *testing.T) {
	for k, v := range []struct {
		samples  int
		classes  int
		features int
		minority int
		weights  []float64
	}{
		{43400, 5, 21, 783, []float64{0.8159428463996992, 0.8102305610006534, 0.82345128545678