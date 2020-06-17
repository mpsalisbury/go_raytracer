package raytracer

import (
	"math"
	"testing"
)

func TestPointAsTuple(t *testing.T) {
	x := 4.0
	y := -4.0
	z := 3.0
	w := 1.0

	p := Point{x, y, z}

	if p.tx() != x {
		t.Errorf("tx() = %f; want %f", p.tx(), x)
	}
	if p.ty() != y {
		t.Errorf("ty() = %f; want %f", p.ty(), y)
	}
	if p.tz() != z {
		t.Errorf("tz() = %f; want %f", p.tz(), z)
	}
	if p.tw() != w {
		t.Errorf("tw() = %f; want %f", p.tw(), w)
	}
}

func TestVectorAsTuple(t *testing.T) {
	x := 4.0
	y := -4.0
	z := 3.0
	w := 0.0

	v := Vector{x, y, z}

	if v.tx() != x {
		t.Errorf("tx() = %f; want %f", v.tx(), x)
	}
	if v.ty() != y {
		t.Errorf("ty() = %f; want %f", v.ty(), y)
	}
	if v.tz() != z {
		t.Errorf("tz() = %f; want %f", v.tz(), z)
	}
	if v.tw() != w {
		t.Errorf("tw() = %f; want %f", v.tw(), w)
	}
}

func TestVectorPlusPoint(t *testing.T) {
	v := Vector{3.0, -2.0, 5.0}
	p := Point{-2.0, 3.0, 1.0}
	got, want := v.plusPoint(p), Point{1.0, 1.0, 6.0}
	if got != want {
		t.Errorf("%+v.plusPoint(%+v) = %+v; want %+v", v, p, got, want)
	}
}

func TestVectorPlusVector(t *testing.T) {
	v1 := Vector{3.0, -2.0, 5.0}
	v2 := Vector{-2.0, 3.0, 1.0}
	got, want := v1.plus(v2), Vector{1.0, 1.0, 6.0}
	if got != want {
		t.Errorf("%+v.plus(%+v) = %+v; want %+v", v1, v2, got, want)
	}
}

func TestPointMinusPoint(t *testing.T) {
	p1 := Point{3.0, 2.0, 1.0}
	p2 := Point{5.0, 6.0, 7.0}
	got, want := p1.minus(p2), Vector{-2.0, -4.0, -6.0}
	if got != want {
		t.Errorf("%+v.minus(%+v) = %+v; want %+v", p1, p2, got, want)
	}
}

func TestPointMinusVector(t *testing.T) {
	p := Point{3.0, 2.0, 1.0}
	v := Vector{5.0, 6.0, 7.0}
	got, want := p.minusVector(v), Point{-2.0, -4.0, -6.0}
	if got != want {
		t.Errorf("%+v.minusVector(%+v) = %+v; want %+v", p, v, got, want)
	}
}

func TestVectorMinusVector(t *testing.T) {
	v1 := Vector{3.0, 2.0, 1.0}
	v2 := Vector{5.0, 6.0, 7.0}
	got, want := v1.minus(v2), Vector{-2.0, -4.0, -6.0}
	if got != want {
		t.Errorf("%+v.minus(%+v) = %+v; want %+v", v1, v2, got, want)
	}
}

func TestVectorNegate(t *testing.T) {
	v := Vector{1.0, -2.0, 3.0}
	got, want := v.negate(), Vector{-1.0, 2.0, -3.0}
	if got != want {
		t.Errorf("%+v.negate() = %+v; want %+v", v, got, want)
	}
}

func TestVectorScale(t *testing.T) {
	v := Vector{1.0, -2.0, 3.0}
	got, want := v.scale(3.5), Vector{3.5, -7.0, 10.5}
	if got != want {
		t.Errorf("%+v.scale(3.5) = %+v; want %+v", v, got, want)
	}
}

func TestVectorMagnitude(t *testing.T) {
	tests := []struct {
		v    Vector
		want float64
	}{
		{
			v:    Vector{1.0, 0.0, 0.0},
			want: 1.0,
		},
		{
			v:    Vector{0.0, 1.0, 0.0},
			want: 1.0,
		},
		{
			v:    Vector{0.0, 0.0, 1.0},
			want: 1.0,
		},
		{
			v:    Vector{1.0, 2.0, 3.0},
			want: math.Sqrt(14.0),
		},
		{
			v:    Vector{-1.0, -2.0, -3.0},
			want: math.Sqrt(14.0),
		},
	}
	for _, test := range tests {
		t.Run("Simple", func(t *testing.T) {
			got, want := test.v.magnitude(), test.want
			if got != want {
				t.Errorf("%+v.magnitude() = %f; want %f", test.v, got, want)
			}
		})
	}
}

func TestVectorNorm(t *testing.T) {
	tests := []struct {
		v    Vector
		want Vector
	}{
		{
			v:    Vector{4.0, 0.0, 0.0},
			want: Vector{1.0, 0.0, 0.0},
		},
		{
			v:    Vector{1.0, 2.0, 3.0},
			want: Vector{1.0 / math.Sqrt(14), 2.0 / math.Sqrt(14), 3.0 / math.Sqrt(14)},
		},
	}
	for _, test := range tests {
		t.Run("Simple", func(t *testing.T) {
			got, want := test.v.norm(), test.want
			if got != want {
				t.Errorf("%+v.magnitude() = %+v; want %+v", test.v, got, want)
			}
		})
	}
}

func TestVectorDot(t *testing.T) {
	v1 := Vector{1.0, 2.0, 3.0}
	v2 := Vector{2.0, 3.0, 4.0}
	got, want := v1.dot(v2), 20.0
	if got != want {
		t.Errorf("%+v.dot(%+v) = %f; want %f", v1, v2, got, want)
	}
}

func TestVectorCross(t *testing.T) {
	tests := []struct {
		v1   Vector
		v2   Vector
		want Vector
	}{
		{
			v1:   Vector{1.0, 2.0, 3.0},
			v2:   Vector{2.0, 3.0, 4.0},
			want: Vector{-1.0, 2.0, -1.0},
		},
		{
			v1:   Vector{2.0, 3.0, 4.0},
			v2:   Vector{1.0, 2.0, 3.0},
			want: Vector{1.0, -2.0, 1.0},
		},
	}
	for _, test := range tests {
		t.Run("Simple", func(t *testing.T) {
			got, want := test.v1.cross(test.v2), test.want
			if got != want {
				t.Errorf("%+v.cross(%+v) = %+v; want %+v", test.v1, test.v2, got, want)
			}
		})
	}
}

/*
func TestConstructTuplePoint(t *testing.T) {
  a := createTuple(4.3, -4.2, 3.1, 1.0)
  tr := Truth(t)
  tr.assertThat(a.x).isEqualTo(4.3)
  tr.assertThat(a.y).isEqualTo(-4.2)
  tr.assertThat(a.z).isEqualTo(3.1)
  tr.assertThat(a.w).isEqualTo(1.0)
  tr.assertThat(a.isPoint()).isTrue()
  tr.assertThat(a.isVector()).isFalse()
}
*/
