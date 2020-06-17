package raytracer

import "math"

type Tuple interface {
	tx() float64
	ty() float64
	tz() float64
	tw() float64
}

type Point struct {
	x, y, z float64
}

func (p Point) tx() float64 {
	return p.x
}

func (p Point) ty() float64 {
	return p.y
}

func (p Point) tz() float64 {
	return p.z
}

func (p Point) tw() float64 {
	return 1.0
}

func (p Point) plusVector(v Vector) Point {
	return Point{p.x + v.x, p.y + v.y, p.z + v.z}
}

func (p Point) minusVector(v Vector) Point {
	return Point{p.x - v.x, p.y - v.y, p.z - v.z}
}

func (p Point) minus(p2 Point) Vector {
	return Vector{p.x - p2.x, p.y - p2.y, p.z - p2.z}
}

type Vector struct {
	x, y, z float64
}

func (v Vector) tx() float64 {
	return v.x
}

func (v Vector) ty() float64 {
	return v.y
}

func (v Vector) tz() float64 {
	return v.z
}

func (v Vector) tw() float64 {
	return 0.0
}

func (v Vector) magnitude() float64 {
	return math.Sqrt(v.dot(v))
}

func (v Vector) norm() Vector {
	return v.scale(1.0 / v.magnitude())
}

func (v Vector) scale(s float64) Vector {
	return Vector{s * v.x, s * v.y, s * v.z}
}

func (v Vector) negate() Vector {
	return v.scale(-1.0)
}

func (v Vector) plusPoint(p Point) Point {
	return Point{p.x + v.x, p.y + v.y, p.z + v.z}
}

func (v Vector) plus(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y, v.z + v2.z}
}

func (v Vector) minus(v2 Vector) Vector {
	return Vector{v.x - v2.x, v.y - v2.y, v.z - v2.z}
}

func (v Vector) dot(v2 Vector) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v1 Vector) cross(v2 Vector) Vector {
	x := v1.y*v2.z - v1.z*v2.y
	y := v1.z*v2.x - v1.x*v2.z
	z := v1.x*v2.y - v2.x*v1.y
	return Vector{x, y, z}
}
