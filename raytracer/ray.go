package raytracer

type Ray struct {
  orig Point
  dir Vector
}

func (r Ray) position(t float64) Point {
  return r.orig.PlusV(r.dir.Scale(t))
}

func (r Ray) xform(m *Matrix) Ray {
  return Ray{m.TimesP(r.orig), m.TimesV(r.dir)}
}

