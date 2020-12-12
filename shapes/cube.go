package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
)

type Cube struct {
	Transform datatypes.Matrix
	raytracing.Material
	Parent Shape
}

func GetCube() *Cube {
	c := Cube{}
	c.Transform = datatypes.GetIdentity()
	c.Material = raytracing.GetMaterial()

	return &c
}

func (c *Cube) GetParent() Shape {
	return c.Parent
}

func (c *Cube) SetParent(s Shape) {
	c.Parent = s
}

func (c *Cube) GetMaterial() raytracing.Material {
	return c.Material
}

func (c *Cube) SetMaterial(m raytracing.Material) {
	c.Material = m
}

func (c *Cube) GetTransform() datatypes.Matrix {
	return c.Transform
}

func (c *Cube) SetTransform(m datatypes.Matrix) {
	c.Transform = m
}

func checkAxis(origin, direction float64) (tmin, tmax float64) {
	tminNumerator := -1 - origin
	tmaxNumerator := 1 - origin

	if math.Abs(direction) >= datatypes.EPSILON {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
	} else {
		tmin = tminNumerator * datatypes.INFINITY
		tmax = tmaxNumerator * datatypes.INFINITY
	}

	if tmin > tmax {
		tmp := tmin
		tmin = tmax
		tmax = tmp
	}

	return
}

func (c *Cube) Intersect(r datatypes.Ray) []Intersection {

	xtmin, xtmax := checkAxis(r.Origin.X, r.Direction.X)
	ytmin, ytmax := checkAxis(r.Origin.Y, r.Direction.Y)
	ztmin, ztmax := checkAxis(r.Origin.Z, r.Direction.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return []Intersection{}
	}

	xs := []Intersection{
		Intersection{tmin, c},
		Intersection{tmax, c}}

	return xs
}

func (c *Cube) Normal(obj_p datatypes.Tuple) datatypes.Tuple {

	absX := math.Abs(obj_p.X)
	absY := math.Abs(obj_p.Y)
	absZ := math.Abs(obj_p.Z)

	maxC := math.Max(math.Max(absX, absY), absZ)

	if maxC == absX {
		return datatypes.Vector(obj_p.X, 0, 0)
	} else if maxC == absY {
		return datatypes.Vector(0, obj_p.Y, 0)
	}
	return datatypes.Vector(0, 0, obj_p.Z)
}

func GetGlassCube() Shape {
	s := GetCube()

	material := s.GetMaterial()
	material.Transparency = 1.0
	material.RefractiveIndex = 1.5
	s.SetMaterial(material)

	return s
}
