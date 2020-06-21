package main

import (
	"log"
	"math"
	rt "raytracer/raytracer"
)

func main() {
	c := rt.MakeCanvas(200, 200)

	center := rt.Point{100, 100, 0}
	arm := rt.Vector{0, 80, 0}
	rot := rt.MakeRotationZ(math.Pi / 6)
	for i := 0; i < 12; i++ {
		p := center.PlusV(arm)
		c.Set(int(p.X), c.Height-int(p.Y), rt.White())
		arm = rot.TimesV(arm)
	}

	err := c.WritePng("clock.png")
	if err != nil {
		log.Print(err)
	}
}
