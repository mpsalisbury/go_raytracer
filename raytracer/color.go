package raytracer

import "image/color"

type Color struct {
	r, g, b float64
}

func Black() Color {
	return Color{0, 0, 0}
}

func White() Color {
	return Color{1, 1, 1}
}

func Red() Color {
	return Color{1, 0, 0}
}

func (c Color) asNRGBA() color.NRGBA {
	return color.NRGBA{
		R: to255(c.r),
		G: to255(c.g),
		B: to255(c.b),
		A: 255,
	}
}

func to255(c float64) uint8 {
	if c <= 0.0 {
		return uint8(0)
	}
	if c >= 1.0 {
		return uint8(255)
	}
	return uint8(c * 255)
}

func (c1 Color) plus(c2 Color) Color {
	return Color{c1.r + c2.r, c1.g + c2.g, c1.b + c2.b}
}

func (c1 Color) minus(c2 Color) Color {
	return Color{c1.r - c2.r, c1.g - c2.g, c1.b - c2.b}
}

func (c1 Color) times(c2 Color) Color {
	return Color{c1.r * c2.r, c1.g * c2.g, c1.b * c2.b}
}

func (c Color) scale(s float64) Color {
	return Color{s * c.r, s * c.g, s * c.b}
}
