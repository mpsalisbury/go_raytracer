package main

import (
	"log"
	rt "raytracer/raytracer"
)

func main() {
	cam := rt.Camera{200, 200}
	s := rt.NewSphere(rt.MakeIdentity())
	s.Material().Color = rt.Color{1, 0.2, 1}
	canvas := cam.Render(s, rt.NewLightShader())

	err := canvas.WritePng("blueball.png")
	if err != nil {
		log.Print(err)
	}
}
