package raytracer

import (
	"fmt"
	"testing"
)

func TestMatrixEntries2x2(t *testing.T) {
	m := MakeMatrix([][]float64{
		{-3.0, 5.0},
		{1.0, -2.0},
	})

	tests := []struct {
		r, c int
		want float64
	}{{
		r:    0,
		c:    0,
		want: -3.0,
	},
		{
			r:    0,
			c:    1,
			want: 5.0,
		},
		{
			r:    1,
			c:    0,
			want: 1.0,
		},
		{
			r:    1,
			c:    1,
			want: -2.0,
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
		{-3.0, 5.0, 0.0},
		{1.0, -2.0, -7.0},
		{0.0, 1.0, 1.0},
	})

	tests := []struct {
		r, c int
		want float64
	}{{
		r:    0,
		c:    0,
		want: -3.0,
	},
		{
			r:    1,
			c:    1,
			want: -2.0,
		},
		{
			r:    2,
			c:    2,
			want: 1.0,
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
		{1.0, 2.0, 3.0, 4.0},
		{5.5, 6.5, 7.5, 8.5},
		{9.0, 10.0, 11.0, 12.0},
		{13.5, 14.5, 15.5, 16.5},
	})

	tests := []struct {
		r, c int
		want float64
	}{{
		r:    0,
		c:    0,
		want: 1.0,
	},
		{
			r:    0,
			c:    3,
			want: 4.0,
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
			want: 11.0,
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
				{1.0, 2.0, 3.0, 4.0},
				{5.0, 6.0},
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
		{1.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0, 8.0},
		{9.0, 8.0, 7.0, 6.0},
		{5.0, 4.0, 3.0, 2.0},
	})
	m2 := MakeMatrix([][]float64{
		{1.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0, 8.0},
		{9.0, 8.0, 7.0, 6.0},
		{5.0, 4.0, 3.0, 2.0},
	})
	m3 := MakeMatrix([][]float64{
		{2.0, 3.0, 4.0, 5.0},
		{6.0, 7.0, 8.0, 9.0},
		{8.0, 7.0, 6.0, 5.0},
		{4.0, 3.0, 2.0, 1.0},
	})

	if !approxEq(m1, m2) {
		t.Error(approxError(m1, m2))
	}
	if approxEq(m1, m3) {
		t.Errorf("should be different: %+v, %+v", m1, m2)
	}
}

func TestMatrixMultiply(t *testing.T) {
	m1 := MakeMatrix([][]float64{
		{1.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0, 8.0},
		{9.0, 8.0, 7.0, 6.0},
		{5.0, 4.0, 3.0, 2.0},
	})
	m2 := MakeMatrix([][]float64{
		{-2.0, 1.0, 2.0, 3.0},
		{3.0, 2.0, 1.0, -1.0},
		{4.0, 3.0, 6.0, 5.0},
		{1.0, 2.0, 7.0, 8.0},
	})
	want := MakeMatrix([][]float64{
		{20.0, 22.0, 50.0, 48.0},
		{44.0, 54.0, 114.0, 108.0},
		{40.0, 58.0, 110.0, 102.0},
		{16.0, 26.0, 46.0, 42.0},
	})

	got := m1.times(m2)
	if !approxEq(got, want) {
		t.Error(approxError(got, want))
	}
}

func TestMatrixIdentity(t *testing.T) {
	m := MakeMatrix([][]float64{
		{0.0, 1.0, 2.0, 4.0},
		{1.0, 2.0, 4.0, 8.0},
		{2.0, 4.0, 8.0, 16.0},
		{4.0, 8.0, 16.0, 32.0},
	})
	got, want := m.times(MakeMatrixIdentity(4)), m
	if !approxEq(got, want) {
		t.Error(approxError(got, want))
	}
}

func TestMatrixTranspose(t *testing.T) {
	m := MakeMatrix([][]float64{
		{0.0, 9.0, 3.0, 0.0},
		{9.0, 8.0, 0.0, 8.0},
		{1.0, 8.0, 5.0, 3.0},
		{0.0, 0.0, 5.0, 8.0},
	})
	want := MakeMatrix([][]float64{
		{0.0, 9.0, 1.0, 0.0},
		{9.0, 8.0, 8.0, 0.0},
		{3.0, 0.0, 5.0, 5.0},
		{0.0, 8.0, 3.0, 8.0},
	})
  got := m.transpose()
	if !approxEq(got, want) {
		t.Error(approxError(got, want))
	}
}

