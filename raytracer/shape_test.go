package raytracer

import "testing"

func TestSphereIntersection(t *testing.T) {
  s := Sphere{}

  tests := []struct {
    name string
    r Ray
    want []Intersection
  }{
    {
      name: "ray hits sphere",
      r: Ray{Point{0,0,-5},Vector{0,0,1}},
      want: []Intersection{
        Intersection{4, s},
        Intersection{6, s},
      },
    },
    {
      name: "ray tangent to sphere",
      r: Ray{Point{0,1,-5},Vector{0,0,1}},
      want: []Intersection{
        Intersection{5, s},
        Intersection{5, s},
      },
    },
    {
      name: "ray misses sphere",
      r: Ray{Point{0,2,-5},Vector{0,0,1}},
      want: []Intersection{},
    },
    {
      name: "ray inside sphere",
      r: Ray{Point{0,0,0},Vector{0,0,1}},
      want: []Intersection{
        Intersection{-1, s},
        Intersection{1, s},
      },
    },
    {
      name: "sphere behind ray",
      r: Ray{Point{0,0,5},Vector{0,0,1}},
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

func TestHit(t *testing.T) {
  s := Sphere{}
  tests := []struct{
    name string
    xs []Intersection
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

