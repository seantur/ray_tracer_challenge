package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/canvas"
	"github.com/seantur/ray_tracer_challenge/tuples"
	"math"
)

type PointLight struct {
	Intensity canvas.Color
	Position  tuples.Tuple
}

func Lighting(material Material, light PointLight, point tuples.Tuple, eyev tuples.Tuple, normalv tuples.Tuple) canvas.Color {

	diffuse := canvas.Color{}
	specular := canvas.Color{}

	effective_color := canvas.Hadamard(material.Color, light.Intensity)

	lightv := tuples.Subtract(light.Position, point)
	lightv = lightv.Normalize()

	ambient := effective_color.Multiply(material.Ambient)

	light_dot_normal := tuples.Dot(lightv, normalv)

	if light_dot_normal >= 0 {
		diffuse = effective_color.Multiply(material.Diffuse)
		diffuse = diffuse.Multiply(light_dot_normal)

		reflectv := lightv.Negate()
		reflectv = reflectv.Reflect(normalv)

		reflect_dot_eye := tuples.Dot(reflectv, eyev)

		if reflect_dot_eye <= 0 {
			specular = canvas.Color{}
		} else {

			factor := math.Pow(reflect_dot_eye, material.Shininess)

			specular = light.Intensity.Multiply(material.Specular)
			specular = specular.Multiply(factor)
		}
	}

	output := canvas.Add(ambient, diffuse)
	output = canvas.Add(output, specular)

	return output

}
