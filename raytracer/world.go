package raytracer

import "sort"

type World interface {
	Intersect(r Ray) []Intersection

	AddShape(s Shape)
	AddLight(l Light)

	Shapes() []Shape
	Lights() []Light
}

type world struct {
	shapes []Shape
	lights []Light
}

func (w *world) Intersect(r Ray) []Intersection {
	xs := []Intersection{}
	for _, s := range w.shapes {
		xs = append(xs, s.Intersect(r)...)
	}
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].t < xs[j].t
	})
	return xs
}

func (w *world) AddShape(s Shape) {
	w.shapes = append(w.shapes, s)
}

func (w *world) AddLight(l Light) {
	w.lights = append(w.lights, l)
}

func (w *world) Shapes() []Shape {
	return w.shapes
}

func (w *world) Lights() []Light {
	return w.lights
}

func NewEmptyWorld() World {
	return &world{[]Shape{}, []Light{}}
}

func NewDefaultWorld() World {
	w := NewEmptyWorld()

	w.AddLight(NewPointLight(Point{-10, 10, -10}, White()))

	s1 := NewSphere(MakeIdentity())
	m1 := s1.Material()
	m1.Color = Color{0.8, 1.0, 0.6}
	m1.Diffuse = 0.7
	m1.Specular = 0.2
	w.AddShape(s1)

	s2 := NewSphere(MakeScaling(0.5, 0.5, 0.5))
	w.AddShape(s2)

	return w
}
