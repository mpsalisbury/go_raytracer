package raytracer

import "sort"

type Group struct {
	shapes []Shape
}

	Intersect(r Ray) []MaterialIntersection
	Xform() *Matrix
	NormalAt(p Point) Vector
	Material() *Material

func (w *world) Intersect(r Ray) []MaterialIntersection {
	xs := []MaterialIntersection{}
	for _, s := range w.shapes {
		xs = append(xs, s.Intersect(r)...)
	}
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].t < xs[j].t
	})
	return xs
}

func (g *Group) AddShape(s Shape) {
	g.shapes = append(g.shapes, s)
}

func NewGroup() *Group {
	return &Group{[]Shape{}}
}
