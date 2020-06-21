package raytracer

import (
	"fmt"
	"testing"
)

func TestRayAtT(t *testing.T) {
	r := Ray{Point{2, 3, 4}, Vector{1, 0, 0}}

	tests := []struct {
		t    float64
		want Point
	}{{
		t:    0,
		want: Point{2, 3, 4},
	},
		{
			t:    1,
			want: Point{3, 3, 4},
		},
		{
			t:    -1,
			want: Point{1, 3, 4},
		},
		{
			t:    2.5,
			want: Point{4.5, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("t=%f", test.t), func(t *testing.T) {
			got, want := r.position(test.t), test.want
			if !approxEq(got, want) {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}

func TestRayXform(t *testing.T) {
	tests := []struct {
		name string
		r    Ray
		xf   *Matrix
		want Ray
	}{
		{
			name: "translate",
			r:    Ray{Point{1, 2, 3}, Vector{0, 1, 0}},
			xf:   MakeTranslation(3, 4, 5),
			want: Ray{Point{4, 6, 8}, Vector{0, 1, 0}},
		},
		{
			name: "scale",
			r:    Ray{Point{1, 2, 3}, Vector{0, 1, 0}},
			xf:   MakeScaling(2, 3, 4),
			want: Ray{Point{2, 6, 12}, Vector{0, 3, 0}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, want := test.r.xform(test.xf), test.want
			if !approxEq(got, want) {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}
