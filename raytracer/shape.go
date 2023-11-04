package raytracer

import (
	"math"
)

type Shape interface {
	Intersect(r Ray) []MaterialIntersection
	Xform() *Matrix
	NormalAt(p Point) Vector
	Material() *Material
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

func (s *sphere) Intersect(r Ray) []MaterialIntersection {
  var xs []MaterialIntersection
  for _, t := range s.intersectPoints(r) {
    p := r.position(t)
    normalV := s.NormalAt(p)
    xs := append(xs, NewMaterialIntersection(r, t, normalV, s.material))
  }
  return xs
}

func (s *sphere) intersectPoints(r Ray) []float64 {
	xfray := r.xform(s.xf.inverse())
	sphereToRay := xfray.orig.Minus(Point{0, 0, 0})
	a := xfray.dir.Dot(xfray.dir)
	b := 2.0 * xfray.dir.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	disc := b*b - 4*a*c
	if disc < 0 {
		return []float64{}
	}
	t1 := (-b - math.Sqrt(disc)) / (2 * a)
	t2 := (-b + math.Sqrt(disc)) / (2 * a)
	return []float64{t1, t2}
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
