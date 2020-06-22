package raytracer

import (
	"github.com/google/go-cmp/cmp"
	"math"
)

type Intersection struct {
	t   float64
	Obj Shape
}

func IntersectionComparer() cmp.Option {
	return cmp.Comparer(func(x, y Intersection) bool { return x.t == y.t })
}

// Returns the intersection with the lowest non-negative t, or nil.
func hit(xs []Intersection) *Intersection {
	var lowestX *Intersection
	lowestT := math.Inf(1)

	for i, x := range xs {
		if x.t >= 0 && x.t < lowestT {
			lowestT = x.t
			lowestX = &xs[i]
		}
	}
	return lowestX
}
