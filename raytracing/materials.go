package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/canvas"
)

type Material struct {
	Color     canvas.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func GetMaterial() Material {
	return Material{Color: canvas.Color{Red: 1, Green: 1, Blue: 1},
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0}
}
