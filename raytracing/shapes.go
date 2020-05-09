package raytracing

import (
	"errors"
	"github.com/seantur/ray_tracer_challenge/datatypes"
)

type Shape interface {
	GetMaterial() Material
	SetMaterial(Material)
	GetTransform() datatypes.Matrix
	SetTransform(datatypes.Matrix)
	Normal(datatypes.Tuple) datatypes.Tuple
	Intersect(Ray) []Intersection
}

func Normal(s Shape, world_p datatypes.Tuple) datatypes.Tuple {
	transform := s.GetTransform()
	sTransformInv, _ := transform.Inverse()
	objP := datatypes.TupleMultiply(sTransformInv, world_p)

	objNormal := s.Normal(objP)

	worldNormal := datatypes.TupleMultiply(sTransformInv.Transpose(), objNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}

func Intersect(s Shape, r Ray) []Intersection {
	transform := s.GetTransform()
	tInv, _ := transform.Inverse()
	r = r.Transform(tInv)

	return s.Intersect(r)
}

type Intersection struct {
	T      float64
	Object Shape
}

type Computation struct {
	T, N1, N2                                             float64
	Object                                                Shape
	Point, UnderPoint, Eyev, Normalv, OverPoint, Reflectv datatypes.Tuple
	IsInside                                              bool
}

func (i *Intersection) PrepareComputations(r Ray, intersections []Intersection) Computation {
	c := Computation{}

	c.T = i.T
	c.Object = i.Object
	c.Point = r.Position(c.T)
	c.Eyev = r.Direction.Negate()
	c.Normalv = Normal(c.Object, c.Point)

	if datatypes.Dot(c.Normalv, c.Eyev) < 0 {
		c.IsInside = true
		c.Normalv = c.Normalv.Negate()
	} else {
		c.IsInside = false
	}

	c.OverPoint = datatypes.Add(c.Point, c.Normalv.Multiply(datatypes.EPSILON))
	c.UnderPoint = datatypes.Subtract(c.Point, c.Normalv.Multiply(datatypes.EPSILON))
	c.Reflectv = r.Direction.Reflect(c.Normalv)

	// Refraction calculation
	containers := []Shape{}

	for _, intersection := range intersections {

		if intersection == *i {
			if len(containers) == 0 {
				c.N1 = 1.0
			} else {
				material := containers[len(containers)-1].GetMaterial()
				c.N1 = material.RefractiveIndex
			}
		}

		included := false
		for index, item := range containers {
			if item == intersection.Object {
				containers = append(containers[:index], containers[index+1:]...)
				included = true
				break
			}
		}

		if !included {
			containers = append(containers, intersection.Object)
		}

		if intersection == *i {
			if len(containers) == 0 {
				c.N2 = 1.0
			} else {
				material := containers[len(containers)-1].GetMaterial()
				c.N2 = material.RefractiveIndex
			}
		}
	}

	return c
}

// ByT implements sort.Interface for []Intersection based on the T field
type ByT []Intersection

func (in ByT) Len() int           { return len(in) }
func (in ByT) Swap(i, j int)      { in[i], in[j] = in[j], in[i] }
func (in ByT) Less(i, j int) bool { return in[i].T < in[j].T }

func Hit(intersections []Intersection) (Intersection, error) {

	if len(intersections) == 0 {
		return Intersection{}, errors.New("did not find hit")
	}

	var hitVal float64
	hitIntersection := intersections[0]

	for _, intersection := range intersections {
		if intersection.T > 0 && (hitVal == 0 || intersection.T < hitVal) {
			hitIntersection = intersection
			hitVal = intersection.T
		}
	}

	if hitVal == 0 {
		return Intersection{}, errors.New("did not find hit")
	}

	return hitIntersection, nil
}
