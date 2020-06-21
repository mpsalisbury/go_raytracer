package raytracer

type Light struct {
	position  Point
	intensity Color
}

func NewPointLight(position Point, intensity Color) Light {
	return Light{position, intensity}
}
