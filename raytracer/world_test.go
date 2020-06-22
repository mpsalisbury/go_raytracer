package raytracer

import "testing"

func TestWorldIntersect(t *testing.T) {
	w := NewDefaultWorld()
	r := Ray{Point{0, 0, -5}, Vector{0, 0, 1}}
	got := w.Intersect(r)
	want := []Intersection{
		Intersection{4, nil},
		Intersection{4.5, nil},
		Intersection{5.5, nil},
		Intersection{6, nil},
	}
	if !approxEq(got, want) {
		t.Errorf(approxError(got, want))
	}
}
