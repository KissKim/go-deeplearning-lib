package layers

import (
	"fmt"
	"math"
	"strings"
)

type Sigmoid struct {
	layer
}

func (s *Sigmoid) Ac