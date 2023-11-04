package raytracer

type Shader interface {
	ColorAt(s Shape, r Ray) Color
}

type PrimitiveShader struct{}

func (ps PrimitiveShader) ColorAt(s Shape, r Ray) Color {
  mxs := s.Intersect(r)
  xs := Intersections(mxs)
	if hit(xs) != nil {
		return Red()
	}
	return Black()
}

type lightShader struct {
	L Light
}

func NewLightShader() Shader {
	l := NewPointLight(Point{-10, 10, -10}, White())
	return lightShader{l}
}

func (ls lightShader) ColorAt(s Shape, r Ray) Color {
  mxs := s.Intersect(r)
  xs := Intersections(mxs)
	x := hit(xs)
	if x == nil {
		return Black()
	}

	point := r.position(x.t)
	return Lighting(x.material, ls.L, point, r.dir.Negate(), s.NormalAt(point))
}
