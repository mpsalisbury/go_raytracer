package main

import (
	"log"
	rt "raytracer/raytracer"
)

type projectile struct {
	loc rt.Point
	vel rt.Vector
}

type environment struct {
	gravity rt.Vector
	wind    rt.Vector
}

func main() {
	c := rt.MakeCanvas(900, 550)

	p := projectile{rt.Point{0, 1, 0}, rt.Vector{1, 1.8, 0}.Norm().Scale(11.25)}
	e := environment{rt.Vector{0, -0.1, 0}, rt.Vector{-0.01, 0, 0}}
	for p.loc.Y > 0 {
		c.Set(int(p.loc.X), c.Height-int(p.loc.Y), rt.White())
		p.loc = p.loc.PlusVector(p.vel)
		p.vel = p.vel.Plus(e.gravity).Plus(e.wind)
	}

	err := c.WritePng("rocket.png")
	if err != nil {
		log.Print(err)
	}
}
