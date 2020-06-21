package raytracer

import (
	"math"
)

type Tuple interface {
	tx() float64
	ty() float64
	tz() float64
	tw() float64
}

func toMatrix(t Tuple) *Matrix {
	return MakeMatrix([][]float64{{t.tx()}, {t.ty()}, {t.tz()}, {t.tw()}})
}

type Point struct {
	X, Y, Z float64
}

func (p Point) tx() float64 {
	return p.X
}

func (p Point) ty() float64 {
	return p.Y
}

func (p Point) tz() float64 {
	return p.Z
}

func (p Point) tw() float64 {
	return 1.0
}

func (p Point) PlusV(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func (p Point) MinusV(v Vector) Point {
	return Point{p.X - v.X, p.Y - v.Y, p.Z - v.Z}
}

func (p Point) Minus(p2 Point) Vector {
	return Vector{p.X - p2.X, p.Y - p2.Y, p.Z - p2.Z}
}

type Vector struct {
	X, Y, Z float64
}

func (v Vector) tx() float64 {
	return v.X
}

func (v Vector) ty() float64 {
	return v.Y
}

func (v Vector) tz() float64 {
	return v.Z
}

func (v Vector) tw() float64 {
	return 0.0
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vector) Norm() Vector {
	return v.Scale(1.0 / v.Magnitude())
}

func (v Vector) Scale(s float64) Vector {
	return Vector{s * v.X, s * v.Y, s * v.Z}
}

func (v Vector) Negate() Vector {
	return v.Scale(-1.0)
}

func (v Vector) PlusP(p Point) Point {
	return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func (v Vector) Plus(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vector) Minus(v2 Vector) Vector {
	return Vector{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v1 Vector) Cross(v2 Vector) Vector {
	x := v1.Y*v2.Z - v1.Z*v2.Y
	y := v1.Z*v2.X - v1.X*v2.Z
	z := v1.X*v2.Y - v2.X*v1.Y
	return Vector{x, y, z}
}

func (v Vector) Reflect(normal Vector) Vector {
	return v.Minus(normal.Scale(2.0 * v.Dot(normal)))
}
