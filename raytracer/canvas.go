package raytracer

import (
	"image"
	"image/png"
	"os"
)

type Canvas struct {
	Width, Height int
	pixel         []Color
}

func MakeCanvas(width, height int) *Canvas {
	if width < 1 || height < 1 {
		panic("invalid canvas size")
	}
	c := &Canvas{width, height, make([]Color, width*height)}
	return c
}

func (c *Canvas) Get(x, y int) Color {
	i := c.index(x, y)
	return c.pixel[i]
}

func (c *Canvas) Set(x, y int, color Color) {
	i := c.index(x, y)
	c.pixel[i] = color
}

func (c *Canvas) index(x, y int) int {
	if x < 0 || x >= c.Width {
		panic("x out of bounds")
	}
	if y < 0 || y >= c.Height {
		panic("y out of bounds")
	}
	return y*c.Width + x
}

func (c *Canvas) WritePng(filename string) error {
	img := image.NewNRGBA(image.Rect(0, 0, c.Width, c.Height))
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			img.Set(x, y, c.Get(x, y).asNRGBA())
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
