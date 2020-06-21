package raytracer

import (
	"fmt"
	"testing"
)

func TestColorPlus(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	t.Run(fmt.Sprintf("%+v.plus(%+v)", c1, c2), func(t *testing.T) {
		got, want := c1.plus(c2), Color{1.6, 0.7, 1.0}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestColorMinus(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	t.Run(fmt.Sprintf("%+v.minus(%+v)", c1, c2), func(t *testing.T) {
		got, want := c1.minus(c2), Color{0.2, 0.5, 0.5}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestColorTimes(t *testing.T) {
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}
	t.Run(fmt.Sprintf("%+v.times(%+v)", c1, c2), func(t *testing.T) {
		got, want := c1.times(c2), Color{0.9, 0.2, 0.04}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestColorScale(t *testing.T) {
	c := Color{0.2, 0.3, 0.4}
	s := 2.0
	t.Run(fmt.Sprintf("%+v.scale(%f)", c, s), func(t *testing.T) {
		got, want := c.scale(s), Color{0.4, 0.6, 0.8}
		if !approxEq(got, want) {
			t.Error(approxError(got, want))
		}
	})
}

func TestAs255(t *testing.T) {
	tests := []struct {
		c    float64
		want uint8
	}{
		{
			c:    0.0,
			want: 0,
		},
		{
			c:    1.0,
			want: 255,
		},
		{
			c:    -0.1,
			want: 0,
		},
		{
			c:    1.1,
			want: 255,
		},
		{
			c:    0.5,
			want: 127,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Color %f", test.c), func(t *testing.T) {
			got, want := to255(test.c), test.want
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
