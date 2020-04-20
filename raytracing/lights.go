package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
)

type PointLight struct {
	Intensity Color
	Position  datatypes.Tuple
}

func Lighting(material Material, light PointLight, point datatypes.Tuple, eyev datatypes.Tuple, normalv datatypes.Tuple, is_shadow bool) Color {

	diffuse := Color{}
	specular := Color{}

	effective_color := Hadamard(material.Color, light.Intensity)

	lightv := datatypes.Subtract(light.Position, point)
	lightv = lightv.Normalize()

	ambient := effective_color.Multiply(material.Ambient)

	if is_shadow {
		return ambient
	}

	light_dot_normal := datatypes.Dot(lightv, normalv)

	if light_dot_normal >= 0 {
		diffuse = effective_color.Multiply(material.Diffuse)
		diffuse = diffuse.Multiply(light_dot_normal)

		reflectv := lightv.Negate()
		reflectv = reflectv.Reflect(normalv)

		reflect_dot_eye := datatypes.Dot(reflectv, eyev)

		if reflect_dot_eye <= 0 {
			specular = Color{}
		} else {

			factor := math.Pow(reflect_dot_eye, material.Shininess)

			specular = light.Intensity.Multiply(material.Specular)
			specular = specular.Multiply(factor)
		}
	}

	output := Add(ambient, diffuse)
	output = Add(output, specular)

	return output

}
