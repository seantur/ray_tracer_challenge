package scene

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"sort"
)

type World struct {
	Light  raytracing.PointLight
	Shapes []raytracing.Sphere
}

func GetWorld() World {
	w := World{Light: raytracing.PointLight{Position: datatypes.Point(-10, 10, -10), Intensity: raytracing.Color{Red: 1, Green: 1, Blue: 1}}}

	s1 := raytracing.GetSphere()
	s1.Material.Color = raytracing.Color{Red: 0.8, Green: 1.0, Blue: 0.6}
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := raytracing.GetSphere()
	s2.Transform = datatypes.GetScaling(0.5, 0.5, 0.5)

	w.Shapes = []raytracing.Sphere{s1, s2}

	return w
}

func (w *World) Intersect(r raytracing.Ray) []raytracing.Intersection {

	intersections := []raytracing.Intersection{}

	for _, shape := range w.Shapes {
		intersections = append(intersections, shape.Intersect(r)...)
	}

	sort.Sort(raytracing.ByT(intersections))

	return intersections

}
