package raytracer

import "testing"

func TestLighting(t *testing.T) {
	tests := []struct {
		name    string
		eyev    Vector
		normalv Vector
		light   Light
		want    Color
	}{
		{
			name:    "Lighting with the eye between the light and the surface",
			eyev:    Vector{0, 0, -1},
			normalv: Vector{0, 0, -1},
			light:   NewPointLight(Point{0, 0, -10}, White()),
			want:    Color{1.9, 1.9, 1.9},
		},
		{
			name:    "Lighting with the eye between the light and the surface, eye offset 45",
			eyev:    Vector{0, 1, -1}.Norm(),
			normalv: Vector{0, 0, -1},
			light:   NewPointLight(Point{0, 0, -10}, White()),
			want:    Color{1.0, 1.0, 1.0},
		},
		{
			name:    "Lighting with the eye opposite surface, light offset 45",
			eyev:    Vector{0, 0, -1},
			normalv: Vector{0, 0, -1},
			light:   NewPointLight(Point{0, 10, -10}, White()),
			want:    Color{0.7364, 0.7364, 0.7364},
		},
		{
			name:    "Lighting with the eye in the path of the reflection vector",
			eyev:    Vector{0, -1, -1}.Norm(),
			normalv: Vector{0, 0, -1},
			light:   NewPointLight(Point{0, 10, -10}, White()),
			want:    Color{1.6364, 1.6364, 1.6364},
		},
		{
			name:    "Lighting with the light behind the surface",
			eyev:    Vector{0, 0, -1},
			normalv: Vector{0, 0, -1},
			light:   NewPointLight(Point{0, 0, 10}, White()),
			want:    Color{0.1, 0.1, 0.1},
		},
	}

	material := NewMaterial()
	position := Point{0, 0, 0}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, want := Lighting(material, test.light, position, test.eyev, test.normalv), test.want
			if !approxEq(got, want) {
				t.Error(approxError(got, want))
			}
		})
	}
}
