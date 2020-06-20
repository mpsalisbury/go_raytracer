package raytracer

import "math"

type Shape interface {
	Intersect(r Ray) []Intersection
}

type Intersection struct {
  t float64
  obj Shape
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

type Sphere struct {}

func (s Sphere) Intersect(r Ray) []Intersection {
  sphereToRay := r.orig.Minus(Point{0,0,0})
  a := r.dir.Dot(r.dir)
  b := 2.0 * r.dir.Dot(sphereToRay)
  c := sphereToRay.Dot(sphereToRay) - 1
  disc := b*b - 4*a*c
  if disc < 0 {
    return []Intersection{}
  }
  t1 := (-b - math.Sqrt(disc)) / (2*a)
  t2 := (-b + math.Sqrt(disc)) / (2*a)
  return []Intersection{
    Intersection{t1, s},
    Intersection{t2, s},
  }
}
