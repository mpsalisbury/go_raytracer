package raytracer

import (
	"github.com/google/go-cmp/cmp"
	"math"
)

// The raw intersection returned by intersect().
type MaterialIntersection struct {
  ray Ray
  t float64
  normalV Vector
  material *Material
}

func NewMaterialIntersection(r Ray, t float64, normalV Vector, m *Material) MaterialIntersection {
	return MaterialIntersection{r, t, normalV, m}
}

func Intersections(mxs []MaterialIntersection) []Intersection {
  xs := []Intersection{}
  for _, mx := range mxs {
    xs = append(xs, NewIntersection(mx))
  }
  return xs
}

// An intersection with additional computed values.
type Intersection struct {
	t float64
	point   Point
	eyeV    Vector
	normalV Vector
  inside bool
  material *Material
}

func NewIntersection(mi MaterialIntersection) Intersection {
  point := mi.ray.position(mi.t)
  eyeV := mi.ray.dir.Negate()
  normalV := mi.normalV
  inside := false
  if normalV.Dot(eyeV) < 0.0 {
    inside = true
    normalV = normalV.Negate()
  }

	return Intersection{mi.t, point, eyeV, normalV, inside, mi.material}
}

func IntersectionTComparer() cmp.Option {
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
