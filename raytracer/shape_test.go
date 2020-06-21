package raytracer

import "testing"

func TestSphereDefaultTransform(t *testing.T) {
	s := NewSphere()
	got, want := s.Xform(), MakeIdentity()
	if !approxEq(got, want) {
		t.Errorf(approxError(got, want))
	}
}

func TestSphereSetTransform(t *testing.T) {
	s := NewSphere()
	xf := MakeTranslation(2, 3, 4)
	s.SetXform(xf)
	got, want := s.Xform(), xf
	if !approxEq(got, want) {
		t.Errorf(approxError(got, want))
	}
}

func TestSphereIntersection(t *testing.T) {
	s := NewSphere()

	tests := []struct {
		name string
		r    Ray
		want []Intersection
	}{
		{
			name: "ray hits sphere",
			r:    Ray{Point{0, 0, -5}, Vector{0, 0, 1}},
			want: []Intersection{
				Intersection{4, s},
				Intersection{6, s},
			},
		},
		{
			name: "ray tangent to sphere",
			r:    Ray{Point{0, 1, -5}, Vector{0, 0, 1}},
			want: []Intersection{
				Intersection{5, s},
				Intersection{5, s},
			},
		},
		{
			name: "ray misses sphere",
			r:    Ray{Point{0, 2, -5}, Vector{0, 0, 1}},
			want: []Intersection{},
		},
		{
			name: "ray inside sphere",
			r:    Ray{Point{0, 0, 0}, Vector{0, 0, 1}},
			want: []Intersection{
				Intersection{-1, s},
				Intersection{1, s},
			},
		},
		{
			name: "sphere behind ray",
			r:    Ray{Point{0, 0, 5}, Vector{0, 0, 1}},
			want: []Intersection{
				Intersection{-6, s},
				Intersection{-4, s},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, want := s.Intersect(test.r), test.want
			if !approxEq(got, want) {
				t.Errorf(approxError(got, want))
			}
		})
	}
}

func TestTransformedSphereIntersection(t *testing.T) {
	tests := []struct {
		name string
		r    Ray
		xf   *Matrix
		want []Intersection
	}{
		{
			name: "scaled sphere",
			r:    Ray{Point{0, 0, -5}, Vector{0, 0, 1}},
			xf:   MakeScaling(2, 2, 2),
			want: []Intersection{
				Intersection{3, nil},
				Intersection{7, nil},
			},
		},
		{
			name: "translated sphere",
			r:    Ray{Point{0, 0, -5}, Vector{0, 0, 1}},
			xf:   MakeTranslation(5, 0, 0),
			want: []Intersection{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewSphere()
			s.SetXform(test.xf)
			got, want := s.Intersect(test.r), test.want
			if !approxEq(got, want) {
				t.Errorf(approxError(got, want))
			}
		})
	}
}

func TestHit(t *testing.T) {
	s := NewSphere()
	tests := []struct {
		name string
		xs   []Intersection
		want *Intersection
	}{
		{
			name: "all positive",
			xs: []Intersection{
				Intersection{1, s},
				Intersection{2, s},
			},
			want: &Intersection{1, s},
		},
		{
			name: "some positive, some negative",
			xs: []Intersection{
				Intersection{-1, s},
				Intersection{1, s},
			},
			want: &Intersection{1, s},
		},
		{
			name: "all negative",
			xs: []Intersection{
				Intersection{-2, s},
				Intersection{-1, s},
			},
			want: nil,
		},
		{
			name: "lowest non-negative",
			xs: []Intersection{
				Intersection{5, s},
				Intersection{7, s},
				Intersection{-3, s},
				Intersection{2, s},
			},
			want: &Intersection{2, s},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, want := hit(test.xs), test.want
			if !approxEq(got, want) {
				t.Errorf(approxError(got, want))
			}
		})
	}
}
