package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/tuples"
	"math"
)

type Sphere struct {
}

func Intersect(s Sphere, r Ray) []Intersection {
	sphereToRay := tuples.Subtract(r.Origin, tuples.Point(0, 0, 0))

	a := tuples.Dot(r.Direction, r.Direction)
	b := 2 * tuples.Dot(r.Direction, sphereToRay)
	c := tuples.Dot(sphereToRay, sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return make([]Intersection, 0)
	}

	var ret []Intersection

	return append(
		ret,
		Intersection{(-b - math.Sqrt(discriminant)) / (2 * a), &s},
		Intersection{(-b + math.Sqrt(discriminant)) / (2 * a), &s})
}

type Intersection struct {
	t      float64
	object *Sphere
}
