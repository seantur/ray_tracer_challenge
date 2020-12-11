package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
)

type Plane struct {
	Transform datatypes.Matrix
	Material
}

func GetPlane() *Plane {
	s := Plane{}
	s.Transform = datatypes.GetIdentity()
	s.Material = GetMaterial()

	return &s
}

func (p *Plane) GetMaterial() Material {
	return p.Material
}

func (p *Plane) SetMaterial(m Material) {
	p.Material = m
}

func (p *Plane) GetTransform() datatypes.Matrix {
	return p.Transform
}

func (p *Plane) SetTransform(m datatypes.Matrix) {
	p.Transform = m
}

func (p *Plane) Intersect(r datatypes.Ray) []Intersection {
	if math.Abs(r.Direction.Y) < datatypes.EPSILON {
		return []Intersection{}
	}

	t := -r.Origin.Y / r.Direction.Y

	return []Intersection{Intersection{T: t, Object: p}}
}

func (p *Plane) Normal(obj_p datatypes.Tuple) datatypes.Tuple {
	return datatypes.Vector(0, 1, 0)
}
