package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
)

type Plane struct {
	Transform datatypes.Matrix
	raytracing.Material
	Parent Shape
}

func GetPlane() *Plane {
	s := Plane{}
	s.Transform = datatypes.GetIdentity()
	s.Material = raytracing.GetMaterial()

	return &s
}

func (p *Plane) GetParent() Shape {
	return p.Parent
}

func (p *Plane) SetParent(s Shape) {
	p.Parent = s
}

func (p *Plane) GetMaterial() raytracing.Material {
	return p.Material
}

func (p *Plane) SetMaterial(m raytracing.Material) {
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
