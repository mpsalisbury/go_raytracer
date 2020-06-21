package raytracer

import (
	"github.com/google/go-cmp/cmp"
	"math"
)

type Shape interface {
	Intersect(r Ray) []Intersection
	Xform() *Matrix
	NormalAt(p Point) Vector
	Material() *Material
}

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

type sphere struct {
	xf       *Matrix
	ixf      *Matrix
	tixf     *Matrix
	material *Material
}

func NewSphere(m *Matrix) Shape {
	xf := m.Copy()
	ixf := xf.inverse()
	tixf := ixf.transpose()
	material := NewMaterial()
	return &sphere{xf, ixf, tixf, material}
}

func (s *sphere) Intersect(r Ray) []Intersection {
	xfray := r.xform(s.xf.inverse())
	sphereToRay := xfray.orig.Minus(Point{0, 0, 0})
	a := xfray.dir.Dot(xfray.dir)
	b := 2.0 * xfray.dir.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	disc := b*b - 4*a*c
	if disc < 0 {
		return []Intersection{}
	}
	t1 := (-b - math.Sqrt(disc)) / (2 * a)
	t2 := (-b + math.Sqrt(disc)) / (2 * a)
	return []Intersection{
		Intersection{t1, s},
		Intersection{t2, s},
	}
}

func (s *sphere) Xform() *Matrix {
	return s.xf
}

func (s *sphere) Material() *Material {
	return s.material
}

func (s *sphere) NormalAt(worldPoint Point) Vector {
	objectPoint := s.ixf.TimesP(worldPoint)
	objectNormal := Vector{objectPoint.X, objectPoint.Y, objectPoint.Z}
	worldNormal := s.tixf.TimesV(objectNormal)
	return worldNormal.Norm()
}
