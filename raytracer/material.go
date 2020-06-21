package raytracer

import "math"

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() *Material {
	return &Material{White(), 0.1, 0.9, 0.9, 200.0}
}

func Lighting(m *Material, light Light, position Point, eyeV Vector, normalV Vector) Color {
	effectiveColor := m.Color.times(light.intensity)
	lightV := light.position.Minus(position).Norm()

	ambient := effectiveColor.scale(m.Ambient)
	diffuse := Black()
	specular := Black()

	lightDotNormal := lightV.Dot(normalV)
	if lightDotNormal > 0 {
		diffuse = effectiveColor.scale(m.Diffuse * lightDotNormal)
		reflectV := lightV.Negate().Reflect(normalV)
		reflectDotEye := reflectV.Dot(eyeV)
		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.intensity.scale(m.Specular * factor)
		}
	}

	return ambient.plus(diffuse).plus(specular)
}
