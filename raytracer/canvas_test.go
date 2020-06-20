package raytracer

import (
	"testing"
)

func TestCanvasEntries(t *testing.T) {
	c := MakeCanvas(100, 100)
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			got, want := c.Get(x, y), Black()
			if !approxEq(got, want) {
				t.Errorf("got %f; want %f", got, want)
			}
		}
	}
}

func TestCanvasSetPixel(t *testing.T) {
	c := MakeCanvas(100, 100)
	c.Set(4, 10, Red())
	got, want := c.Get(4, 10), Red()
	if !approxEq(got, want) {
		t.Errorf("got %f; want %f", got, want)
	}
}
