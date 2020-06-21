package raytracer

import (
	"math"
	"testing"
)

func TestSphereDefaultTransform(t *testing.T) {
	s := NewSphere(MakeIdentity())
	got, want := s.Xform(), MakeIdentity()
	if !approxEq(got, want) {
		t.Errorf(approxError(got, want))
	}
}

func TestSphereSetTransform(t *testing.T) {
	xf := MakeTranslation(2, 3, 4)
	s := NewSphere(xf)
	got, want := s.Xform(), xf
	if !approxEq(got, want) {
		t.Errorf(approxError(got, want))
	}
}

func TestSphereIntersection(t *testing.T) {
	s := NewSphere(MakeIdentity())

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
			s := NewSphere(test.xf)
			got, want := s.Intersect(test.r), test.want
			if !approxEq(got, want) {
				t.Errorf(approxError(got, want))
			}
		})
	}
}

func TestHit(t *testing.T) {
	s := NewSphere(MakeIdentity())
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

func TestSphereNormalAt(t *testing.T) {
	tests := []struct {
		name string
		s    Shape
		p    Point
		want Vector
	}{
		{
			name: "on x axis",
			s:    NewSphere(MakeIdentity()),
			p:    Point{1, 0, 0},
			want: Vector{1, 0, 0},
		},
		{
			name: "on y axis",
			s:    NewSphere(MakeIdentity()),
			p:    Point{0, 1, 0},
			want: Vector{0, 1, 0},
		},
		{
			name: "on z axis",
			s:    NewSphere(MakeIdentity()),
			p:    Point{0, 0, 1},
			want: Vector{0, 0, 1},
		},
		{
			name: "on nonaxial point",
			s:    NewSphere(MakeIdentity()),
			p:    Point{math.Sqrt(3) / 3.0, math.Sqrt(3) / 3.0, math.Sqrt(3) / 3.0},
			want: Vector{math.Sqrt(3) / 3.0, math.Sqrt(3) / 3.0, math.Sqrt(3) / 3.0},
		},
		{
			name: "on translated sphere",
			s:    NewSphere(MakeTranslation(0, 1, 0)),
			p:    Point{0, 1.70711, -0.70711},
			want: Vector{0, 0.70711, -0.70711},
		},
		{
			name: "on transformed sphere",
			s:    NewSphere(MakeRotationZ(math.Pi/5).Scale(1, 0.5, 1)),
			p:    Point{0, 0.70711, -0.70711},
			want: Vector{0, 0.97014, -0.24254},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, want := test.s.NormalAt(test.p), test.want
			if !approxEq(got, want) {
				t.Errorf(approxError(got, want))
			}
		})
	}
}
