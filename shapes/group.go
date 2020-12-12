package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"log"
	"sort"
)

type Group struct {
	Transform datatypes.Matrix
	raytracing.Material
	Shapes []Shape
	Parent Shape
}

func GetGroup() *Group {
	g := Group{}

	g.Transform = datatypes.GetIdentity()
	g.Material = raytracing.GetMaterial()

	return &g
}

func (g *Group) GetMaterial() raytracing.Material {
	return g.Material
}

func (g *Group) SetMaterial(m raytracing.Material) {
	g.Material = m
}

func (g *Group) GetTransform() datatypes.Matrix {
	return g.Transform
}

func (g *Group) SetTransform(m datatypes.Matrix) {
	g.Transform = m
}

func (g *Group) Normal(datatypes.Tuple) datatypes.Tuple {
	log.Fatal("groups should not call Normal")
	return datatypes.Tuple{} // needed to satisfy the Shape interface
}

func (g *Group) Intersect(r datatypes.Ray) (intersections []Intersection) {

	for _, shape := range g.Shapes {
		shapeIntersection := Intersect(shape, r)
		intersections = append(intersections, shapeIntersection...)
	}

	sort.Sort(ByT(intersections))

	return
}

func (g *Group) GetParent() Shape {
	return g.Parent
}

func (g *Group) SetParent(parent Shape) {
	g.Parent = parent
}

func (g *Group) AddChild(s Shape) {
	g.Shapes = append(g.Shapes, s)
	s.SetParent(g)
}
