package main

import (
	"log"
	"math"
	rt "raytracer/raytracer"
)

func main() {
	cam := rt.Camera{200, 200}
	s := rt.NewSphere(rt.MakeScaling(0.5, 1, 1).RotateZ(math.Pi / 4.0))
	canvas := cam.Render(s, rt.PrimitiveShader{})

	err := canvas.WritePng("silhouette.png")
	if err != nil {
		log.Print(err)
	}
}
