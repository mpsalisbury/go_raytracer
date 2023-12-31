package raytracer

import (
	"fmt"
	"math"
	"testing"
)

func TestPointAsTuple(t *testing.T) {
	x := 4.0
	y := -4.0
	z := 3.0
	w := 1.0

	p := Point{x, y, z}

	testValue := func(name string, got float64, want float64) {
		if !approxEq(got, want) {
			t.Errorf("%s = %f; want %f", name, got, want)
		}
	}
	testValue("tx()", p.tx(), x)
	testValue("ty()", p.ty(), y)
	testValue("tz()", p.tz(), z)
	testValue("tw()", p.tw(), w)
}

func TestVectorAsTuple(t *testing.T) {
	x := 4.0
	y := -4.0
	z := 3.0
	w := 0.0

	v := Vector{x, y, z}

	testValue := func(name string, got float64, want float64) {
		if !approxEq(got, want) {
			t.Errorf("%s = %f; want %f", name, got, want)
		}
	}
	testValue("tx()", v.tx(), x)
	testValue("ty()", v.ty(), y)
	testValue("tz()", v.tz(), z)
	testValue("tw()", v.tw(), w)
}

func TestVectorPlusP(t *testing.T) {
	v := Vector{3.0, -2.0, 5.0}
	p := Point{-2.0, 3.0, 1.0}
	t.Run(fmt.Sprintf("%+v.PlusP(%+v)", v, p), func(t *testing.T) {
		got, want := v.PlusP(p), Point{1.0, 1.0, 6.0}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestVectorPlus(t *testing.T) {
	v1 := Vector{3.0, -2.0, 5.0}
	v2 := Vector{-2.0, 3.0, 1.0}
	t.Run(fmt.Sprintf("%+v.Plus(%+v)", v1, v2), func(t *testing.T) {
		got, want := v1.Plus(v2), Vector{1.0, 1.0, 6.0}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestPointMinusPoint(t *testing.T) {
	p1 := Point{3.0, 2.0, 1.0}
	p2 := Point{5.0, 6.0, 7.0}
	t.Run(fmt.Sprintf("%+v.Minus(%+v)", p1, p2), func(t *testing.T) {
		got, want := p1.Minus(p2), Vector{-2.0, -4.0, -6.0}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestPointMinusV(t *testing.T) {
	p := Point{3.0, 2.0, 1.0}
	v := Vector{5.0, 6.0, 7.0}
	t.Run(fmt.Sprintf("%+v.MinusV(%+v)", p, v), func(t *testing.T) {
		got, want := p.MinusV(v), Point{-2.0, -4.0, -6.0}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestVectorMinusVector(t *testing.T) {
	v1 := Vector{3.0, 2.0, 1.0}
	v2 := Vector{5.0, 6.0, 7.0}
	t.Run(fmt.Sprintf("%+v.Minus(%+v)", v1, v2), func(t *testing.T) {
		got, want := v1.Minus(v2), Vector{-2.0, -4.0, -6.0}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestVectorNegate(t *testing.T) {
	v := Vector{1.0, -2.0, 3.0}
	t.Run(fmt.Sprintf("%+v.Negate()", v), func(t *testing.T) {
		got, want := v.Negate(), Vector{-1.0, 2.0, -3.0}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestVectorScale(t *testing.T) {
	v := Vector{1.0, -2.0, 3.0}
	t.Run(fmt.Sprintf("%+v.Scale(3.5)", v), func(t *testing.T) {
		got, want := v.Scale(3.5), Vector{3.5, -7.0, 10.5}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
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
		t.Run(fmt.Sprintf("%+v.Magnitude()", test.v), func(t *testing.T) {
			got, want := test.v.Magnitude(), test.want
			if !approxEq(got, want) {
				t.Error(approxError(got, want))
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
		t.Run(fmt.Sprintf("%+v.Norm()", test.v), func(t *testing.T) {
			got, want := test.v.Norm(), test.want
			if !approxEq(got, want) {
				t.Error(approxError(got, want))
			}
		})
	}
}

func TestVectorDot(t *testing.T) {
	v1 := Vector{1.0, 2.0, 3.0}
	v2 := Vector{2.0, 3.0, 4.0}
	t.Run(fmt.Sprintf("%+v.Dot(%+v)", v1, v2), func(t *testing.T) {
		got, want := v1.Dot(v2), 20.0
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
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
		t.Run(fmt.Sprintf("%+v.Cross(%+v)", test.v1, test.v2), func(t *testing.T) {
			got, want := test.v1.Cross(test.v2), test.want
			if !approxEq(got, want) {
				t.Error(approxError(got, want))
			}
		})
	}
}

func TestVectorReflect(t *testing.T) {
	tests := []struct {
		name   string
		v      Vector
		normal Vector
		want   Vector
	}{
		{
			name:   "approaching at 45",
			v:      Vector{1, -1, 0},
			normal: Vector{0, 1, 0},
			want:   Vector{1, 1, 0},
		},
		{
			name:   "off slanted surface",
			v:      Vector{0, -1, 0},
			normal: Vector{0.70711, 0.70711, 0},
			want:   Vector{1, 0, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, want := test.v.Reflect(test.normal), test.want
			if !approxEq(got, want) {
				t.Error(approxError(got, want))
			}
		})
	}
}
