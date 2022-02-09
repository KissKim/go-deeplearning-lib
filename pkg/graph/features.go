
package graph

import (
	"math/rand"
)

type Features struct {
	ClassWeights        []float64
	DisableClassWeights bool
	DisableShuffle      bool
	X                   [][]float64
	Y                   [][]float64
}

func (f *Features) Balance() {
	if len(f.Y) == 0 || len(f.Y[0]) == 0 {
		return
	}
	n := len(f.Y)
	m := len(f.Y[0])
	f.ClassWeights = make([]float64, m)