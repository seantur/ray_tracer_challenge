package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
)

type Cylinder struct {
	Transform datatypes.Matrix
	raytracing.Material
	Min, Max float64
	Closed   bool
	Parent   Shape
}

func GetCylinder() *Cylinder {
	c := Cylinder{}
	c.Transform = datatypes.GetIdentity()
	c.Material = raytracing.GetMaterial()
	c.Min = -datatypes.INFINITY
	c.Max = datatypes.INFINITY

	return &c
}

func (c *Cylinder) GetParent() Shape {
	return c.Parent
}

func (c *Cylinder) SetParent(s Shape) {
	c.Parent = s
}

func (c *Cylinder) GetMaterial() raytracing.Material {
	return c.Material
}

func (c *Cylinder) SetMaterial(m raytracing.Material) {
	c.Material = m
}

func (c *Cylinder) GetTransform() datatypes.Matrix {
	return c.Transform
}

func (c *Cylinder) SetTransform(m datatypes.Matrix) {
	c.Transform = m
}

func (cyl *Cylinder) Intersect(r datatypes.Ray) (xs []Intersection) {
	r.Direction = r.Direction.Normalize()

	a := math.Pow(r.Direction.X, 2) + math.Pow(r.Direction.Z, 2)

	if math.Abs(a) > datatypes.EPSILON {

		b := 2*r.Origin.X*r.Direction.X + 2*r.Origin.Z*r.Direction.Z
		c := math.Pow(r.Origin.X, 2) + math.Pow(r.Origin.Z, 2) - 1

		discriminant := math.Pow(b, 2) - 4*a*c

		if discriminant < 0 {
			return
		}

		t0 := (-b - math.Sqrt(discriminant)) / (2 * a)
		t1 := (-b + math.Sqrt(discriminant)) / (2 * a)

		if t0 > t1 {
			tmp := t0
			t0 = t1
			t1 = tmp
		}

		y0 := r.Origin.Y + t0*r.Direction.Y
		if cyl.Min < y0 && y0 < cyl.Max {
			xs = append(xs, Intersection{t0, cyl})
		}

		y1 := r.Origin.Y + t1*r.Direction.Y
		if cyl.Min < y1 && y1 < cyl.Max {
			xs = append(xs, Intersection{t1, cyl})
		}
	}

	xs = cyl.intersectCap(r, xs)

	return
}

func (c *Cylinder) Normal(obj_p datatypes.Tuple) datatypes.Tuple {
	dist := math.Pow(obj_p.X, 2) + math.Pow(obj_p.Z, 2)

	if dist < 1 && obj_p.Y >= c.Max-datatypes.EPSILON {
		return datatypes.Vector(0, 1, 0)
	} else if dist < 1 && obj_p.Y <= c.Min+datatypes.EPSILON {
		return datatypes.Vector(0, -1, 0)
	}

	return datatypes.Vector(obj_p.X, 0, obj_p.Z)
}

func checkCapCylinder(r datatypes.Ray, t float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	return (math.Pow(x, 2) + math.Pow(z, 2)) <= 1
}

func (c *Cylinder) intersectCap(r datatypes.Ray, xs []Intersection) (xsRet []Intersection) {
	if !c.Closed || (math.Abs(r.Direction.Y) < datatypes.EPSILON) {
		return xs
	}

	t := (c.Min - r.Origin.Y) / r.Direction.Y
	if checkCapCylinder(r, t) {
		xs = append(xs, Intersection{t, c})
	}

	t = (c.Max - r.Origin.Y) / r.Direction.Y
	if checkCapCylinder(r, t) {
		xs = append(xs, Intersection{t, c})
	}

	return xs
}

func GetGlassCylinder() Shape {
	s := GetCylinder()

	material := s.GetMaterial()
	material.Transparency = 1.0
	material.RefractiveIndex = 1.5
	s.SetMaterial(material)

	return s
}
