package raytracer

type Shader interface {
	ColorAt(s Shape, r Ray) Color
}

type PrimitiveShader struct{}

func (ps PrimitiveShader) ColorAt(s Shape, r Ray) Color {
	if hit(s.Intersect(r)) != nil {
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
	x := hit(s.Intersect(r))
	if x == nil {
		return Black()
	}

	point := r.position(x.t)
	return Lighting(x.Obj.Material(), ls.L, point, r.dir.Negate(), s.NormalAt(point))
}
