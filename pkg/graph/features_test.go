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
		{43400, 5, 21,