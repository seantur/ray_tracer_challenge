package scene

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
	"sort"
)

type World struct {
	Light  raytracing.PointLight
	Shapes []raytracing.Shape
}

func GetWorld() World {
	w := World{Light: raytracing.PointLight{Position: datatypes.Point(-10, 10, -10), Intensity: raytracing.RGB{Red: 1, Green: 1, Blue: 1}}}

	s1 := raytracing.GetSphere()

	mat := s1.GetMaterial()
	mat.RGB = raytracing.RGB{Red: 0.8, Green: 1.0, Blue: 0.6}
	mat.Diffuse = 0.7
	mat.Specular = 0.2
	s1.SetMaterial(mat)

	s2 := raytracing.GetSphere()
	s2.SetTransform(datatypes.GetScaling(0.5, 0.5, 0.5))

	w.Shapes = []raytracing.Shape{s1, s2}

	return w
}

func (w *World) Intersect(r datatypes.Ray) []raytracing.Intersection {

	intersections := []raytracing.Intersection{}

	for i, _ := range w.Shapes {
		intersection := raytracing.Intersect(w.Shapes[i], r)
		intersections = append(intersections, intersection...)
	}
	sort.Sort(raytracing.ByT(intersections))

	return intersections

}

func (w *World) ShadeHit(c raytracing.Computation, remaining int) raytracing.RGB {
	shadowed := w.IsShadowed(c.OverPoint)

	surfaceColor := raytracing.Lighting(c.Object.GetMaterial(), c.Object, w.Light, c.OverPoint, c.Eyev, c.Normalv, shadowed)
	reflectedColor := w.ReflectedColor(c, remaining)
	refractedColor := w.RefractedColor(c, remaining)

	return raytracing.Add(surfaceColor, reflectedColor, refractedColor)
}

func (w *World) ColorAt(r datatypes.Ray, remaining int) raytracing.RGB {
	intersections := w.Intersect(r)

	hit, err := raytracing.Hit(intersections)

	if err != nil {
		return raytracing.RGB{}
	}

	comp := hit.PrepareComputations(r, intersections)

	c := w.ShadeHit(comp, remaining)

	return c
}

func (w *World) IsShadowed(p datatypes.Tuple) bool {
	v := datatypes.Subtract(w.Light.Position, p)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := datatypes.Ray{Origin: p, Direction: direction}
	intersections := w.Intersect(r)

	h, err := raytracing.Hit(intersections)
	if (err == nil) && (h.T < distance) {
		return true
	}

	return false
}

func (w *World) ReflectedColor(c raytracing.Computation, remaining int) raytracing.RGB {
	// Avoid infinite recursion
	if remaining < 1 {
		return raytracing.RGB{Red: 0, Green: 0, Blue: 0}
	}
	mat := c.Object.GetMaterial()
	if mat.Reflective == 0 {
		return raytracing.RGB{Red: 0, Green: 0, Blue: 0}
	}

	reflectRay := datatypes.Ray{Origin: c.OverPoint, Direction: c.Reflectv}
	remaining--
	color := w.ColorAt(reflectRay, remaining-1)

	return color.Multiply(mat.Reflective)
}

func (w *World) RefractedColor(c raytracing.Computation, remaining int) raytracing.RGB {
	material := c.Object.GetMaterial()
	if material.Transparency == 0 || remaining == 0 {
		return raytracing.RGB{Red: 0, Green: 0, Blue: 0}
	}

	// Check for total internal reflection
	nRatio := c.N1 / c.N2
	cosI := datatypes.Dot(c.Eyev, c.Normalv)
	sin2T := math.Pow(nRatio, 2) * (1 - math.Pow(cosI, 2))

	if sin2T > 1 {
		return raytracing.RGB{Red: 0, Green: 0, Blue: 0}
	}

	cosT := math.Sqrt(1.0 - sin2T)
	direction := datatypes.Subtract(c.Normalv.Multiply((nRatio*cosI)-cosT), c.Eyev.Multiply(nRatio))
	refractRay := datatypes.Ray{Origin: c.UnderPoint, Direction: direction}

	color := w.ColorAt(refractRay, remaining-1)

	return color.Multiply(material.Transparency)
}
