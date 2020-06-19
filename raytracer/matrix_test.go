package raytracer

import (
	"fmt"
	"testing"
)

func TestMatrixEntries2x2(t *testing.T) {
	m := MakeMatrix([][]float64{
		{-3, 5},
		{1, -2},
	})

	tests := []struct {
		r, c int
		want float64
	}{{
		r:    0,
		c:    0,
		want: -3,
	},
		{
			r:    0,
			c:    1,
			want: 5,
		},
		{
			r:    1,
			c:    0,
			want: 1,
		},
		{
			r:    1,
			c:    1,
			want: -2,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("[%d,%d]", test.r, test.c), func(t *testing.T) {
			got, want := m.get(test.r, test.c), test.want
			if !approxEq(got, want) {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}

func TestMatrixEntries3x3(t *testing.T) {
	m := MakeMatrix([][]float64{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	})

	tests := []struct {
		r, c int
		want float64
	}{{
		r:    0,
		c:    0,
		want: -3,
	},
		{
			r:    1,
			c:    1,
			want: -2,
		},
		{
			r:    2,
			c:    2,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("[%d,%d]", test.r, test.c), func(t *testing.T) {
			got, want := m.get(test.r, test.c), test.want
			if !approxEq(got, want) {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}

func TestMatrixEntries4x4(t *testing.T) {
	m := MakeMatrix([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})

	tests := []struct {
		r, c int
		want float64
	}{{
		r:    0,
		c:    0,
		want: 1,
	},
		{
			r:    0,
			c:    3,
			want: 4,
		},
		{
			r:    1,
			c:    0,
			want: 5.5,
		},
		{
			r:    1,
			c:    2,
			want: 7.5,
		},
		{
			r:    2,
			c:    2,
			want: 11,
		},
		{
			r:    3,
			c:    0,
			want: 13.5,
		},
		{
			r:    3,
			c:    2,
			want: 15.5,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("[%d,%d]", test.r, test.c), func(t *testing.T) {
			got := m.get(test.r, test.c)
			want := test.want
			if !approxEq(got, want) {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}

func TestNonrectangularMatrix(t *testing.T) {
	assertPanic(t,
		func() {
			MakeMatrix([][]float64{
				{1, 2, 3, 4},
				{5, 6},
			})
		},
		"expected report of invalid number of columns")
}

func TestOutOfBoundsAccess(t *testing.T) {
	tests := []struct {
		r, c      int
		wantPanic bool
	}{
		{
			r:         0,
			c:         0,
			wantPanic: false,
		},
		{
			r:         1,
			c:         1,
			wantPanic: false,
		},
		{
			r:         -1,
			c:         0,
			wantPanic: true,
		},
		{
			r:         0,
			c:         -1,
			wantPanic: true,
		},
		{
			r:         2,
			c:         0,
			wantPanic: true,
		},
		{
			r:         0,
			c:         2,
			wantPanic: true,
		},
	}
	m := MakeMatrixWithSize(2, 2)
	for _, test := range tests {
		t.Run(fmt.Sprintf("[%d,%d]", test.r, test.c), func(t *testing.T) {
			accessor := func() { m.get(test.r, test.c) }
			if test.wantPanic {
				assertPanic(t, accessor, "expected out of bounds panic")
			} else {
				assertNoPanic(t, accessor, "expected no out of bounds panic")
			}
		})
	}
}

func TestMatrixEquality(t *testing.T) {
	m1 := MakeMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	m2 := MakeMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	m3 := MakeMatrix([][]float64{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	})

	if !approxEq(m1, m2) {
		t.Error(approxError(m1, m2))
	}
	if approxEq(m1, m3) {
		t.Errorf("should be different: %+v, %+v", m1, m2)
	}
}

func TestMatrixTimes(t *testing.T) {
	m1 := MakeMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	m2 := MakeMatrix([][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})
	want := MakeMatrix([][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})

	got := m1.times(m2)
	if !approxEq(got, want) {
		t.Error(approxError(got, want))
	}
}

func TestMatrixTimesPoint(t *testing.T) {
	m := MakeMatrix([][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	p := Point{1, 2, 3}
	want := Point{18, 24, 33}

	got := m.timesPoint(p)
	if !approxEq(got, want) {
		t.Error(approxError(got, want))
	}
}

func TestMatrixTimesVector(t *testing.T) {
	m := MakeMatrix([][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	v := Vector{1, 2, 3}
	want := Vector{14, 22, 32}

	got := m.timesVector(v)
	if !approxEq(got, want) {
		t.Error(approxError(got, want))
	}
}

func TestMatrixIdentity(t *testing.T) {
	m := MakeMatrix([][]float64{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	})
	got, want := m.times(MakeMatrixIdentity(4)), m
	if !approxEq(got, want) {
		t.Error(approxError(got, want))
	}
}

func TestMatrixTranspose(t *testing.T) {
	m := MakeMatrix([][]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	})
	want := MakeMatrix([][]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	})
	got := m.transpose()
	if !approxEq(got, want) {
		t.Error(approxError(got, want))
	}
}

func TestMatrixDeterminant(t *testing.T) {
	tests := []struct {
		m    *Matrix
		want float64
	}{
		{
			m: MakeMatrix([][]float64{
				{1, 5},
				{-3, 2},
			}),
			want: 17,
		},
		{
			m: MakeMatrix([][]float64{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			}),
			want: -196,
		},
		{
			m: MakeMatrix([][]float64{
				{-2, -8, 3, 5},
				{-3, 1, 7, 3},
				{1, 2, -9, 6},
				{-6, 7, 7, -9},
			}),
			want: -4071,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Mtrx%d", i), func(t *testing.T) {
			got, want := test.m.determinant(), test.want
			if !approxEq(got, want) {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}

func TestMatrixIsInvertible(t *testing.T) {
	tests := []struct {
		m    *Matrix
		want bool
	}{
		{
			m: MakeMatrix([][]float64{
				{6, 4, 4, 4},
				{5, 5, 7, 6},
				{4, -9, 3, -7},
				{9, 1, 7, -6},
			}),
			want: true,
		},
		{
			m: MakeMatrix([][]float64{
				{-4, 2, -2, -3},
				{9, 6, 2, 6},
				{0, -5, 1, -5},
				{0, 0, 0, 0},
			}),
			want: false,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Mtrx%d", i), func(t *testing.T) {
			got, want := test.m.isInvertible(), test.want
			if got != want {
				t.Errorf("got %t; want %t", got, want)
			}
		})
	}
}
func TestMatrixInverse(t *testing.T) {
	tests := []struct {
		m    *Matrix
		want *Matrix
	}{
		{
			m: MakeMatrix([][]float64{
				{-5, 2, 6, -8},
				{1, -5, 1, 8},
				{7, 7, -6, -7},
				{1, -3, 7, 4},
			}),
			want: MakeMatrix([][]float64{
				{0.21805, 0.45113, 0.24060, -0.04511},
				{-0.80827, -1.45677, -0.44361, 0.52068},
				{-0.07895, -0.22368, -0.05263, 0.19737},
				{-0.52256, -0.81391, -0.30075, 0.30639},
			}),
		},
		{
			m: MakeMatrix([][]float64{
				{8, -5, 9, 2},
				{7, 5, 6, 1},
				{-6, 0, 9, 6},
				{-3, 0, -9, -4},
			}),
			want: MakeMatrix([][]float64{
				{-0.15385, -0.15385, -0.28205, -0.53846},
				{-0.07692, 0.12308, 0.02564, 0.03077},
				{0.35897, 0.35897, 0.43590, 0.92308},
				{-0.69231, -0.69231, -0.76923, -1.92308},
			}),
		},
		{
			m: MakeMatrix([][]float64{
				{9, 3, 0, 9},
				{-5, -2, -6, -3},
				{-4, 9, 6, 4},
				{-7, 6, 6, 2},
			}),
			want: MakeMatrix([][]float64{
				{-0.04074, -0.07778, 0.14444, -0.22222},
				{-0.07778, 0.03333, 0.36667, -0.33333},
				{-0.02901, -0.14630, -0.10926, 0.12963},
				{0.17778, 0.06667, -0.26667, 0.33333},
			}),
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Mtrx%d", i), func(t *testing.T) {
			got, want := test.m.inverse(), test.want
			if !approxEq(got, want) {
				t.Errorf(approxError(got, want))
			}
		})
	}
}

func TestMatrixMultiplyByInverse(t *testing.T) {
	tests := []struct {
		m *Matrix
	}{
		{
			m: MakeMatrix([][]float64{
				{3, -9, 7, 3},
				{3, -8, 2, -9},
				{-4, 4, 4, 1},
				{-6, 5, -1, 1},
			}),
		},
		{
			m: MakeMatrix([][]float64{
				{8, 2, 2, 2},
				{3, -1, 7, 0},
				{7, 0, 5, 4},
				{6, -2, 0, 5},
			}),
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Mtrx%d", i), func(t *testing.T) {
			m := test.m
			got, want := m.times(m.inverse()), MakeMatrixIdentity(m.numRows)
			if !approxEq(got, want) {
				t.Errorf(approxError(got, want))
			}
		})
	}
}

func TestMatrixSubmatrix(t *testing.T) {
	tests := []struct {
		m    *Matrix
		r, c int
		want *Matrix
	}{{
		m: MakeMatrix([][]float64{
			{-6, 1, 1, 6},
			{-8, 5, 8, 6},
			{-1, 0, 8, 2},
			{-7, 1, -1, 1},
		}),
		r: 2,
		c: 1,
		want: MakeMatrix([][]float64{
			{-6, 1, 6},
			{-8, 8, 6},
			{-7, -1, 1},
		}),
	},
		{
			m: MakeMatrix([][]float64{
				{1, 5, 0},
				{-3, 2, 7},
				{0, 6, -3},
			}),
			r: 0,
			c: 2,
			want: MakeMatrix([][]float64{
				{-3, 2},
				{0, 6},
			}),
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v.submatrix(%d,%d)", test.m, test.r, test.c), func(t *testing.T) {
			got, want := test.m.submatrix(test.r, test.c), test.want
			if !approxEq(got, want) {
				t.Errorf(approxError(got, want))
			}
		})
	}
}

func TestMatrixMinor(t *testing.T) {
	m := MakeMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	got, want := m.minor(1, 0), 25.0
	if !approxEq(got, want) {
		t.Errorf("got %f; want %f", got, want)
	}
}

func TestMatrixCofactor(t *testing.T) {
	m := MakeMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	tests := []struct {
		r, c int
		want float64
	}{{
		r:    0,
		c:    0,
		want: -12,
	},
		{
			r:    1,
			c:    0,
			want: -25,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v.cofactor(%d,%d)", m, test.r, test.c), func(t *testing.T) {
			got, want := m.cofactor(test.r, test.c), test.want
			if !approxEq(got, want) {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}
