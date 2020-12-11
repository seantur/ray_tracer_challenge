package scene

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"github.com/seantur/ray_tracer_challenge/shapes"
	"math"
)

type PointLight struct {
	Intensity raytracing.RGB
	Position  datatypes.Tuple
}

func Lighting(material raytracing.Material, shape shapes.Shape, light PointLight, point datatypes.Tuple, eyev datatypes.Tuple, normalv datatypes.Tuple, is_shadow bool) raytracing.RGB {

	var materialColor raytracing.RGB

	if material.Pattern != nil {
		materialColor = shapes.AtObj(material.Pattern, shape, point)
	} else {
		materialColor = material.RGB
	}

	diffuse := raytracing.RGB{}
	specular := raytracing.RGB{}

	effective_color := raytracing.Hadamard(materialColor, light.Intensity)

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
			specular = raytracing.RGB{}
		} else {

			factor := math.Pow(reflect_dot_eye, material.Shininess)

			specular = light.Intensity.Multiply(material.Specular)
			specular = specular.Multiply(factor)
		}
	}

	output := raytracing.Add(ambient, diffuse)
	output = raytracing.Add(output, specular)

	return output

}
