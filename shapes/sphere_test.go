package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
	"reflect"
	"testing"
)

func TestSpheres(t *testing.T) {

	t.Run("Sphere's default transform is identity matrix", func(t *testing.T) {
		s := GetSphere()

		datatypes.AssertMatrixEqual(t, s.GetTransform(), datatypes.GetIdentity())
	})

	t.Run("Changing a sphere's transformation", func(t *testing.T) {
		s := GetSphere()
		s.SetTransform(datatypes.GetTranslation(2, 3, 4))

		datatypes.AssertMatrixEqual(t, s.GetTransform(), datatypes.GetTranslation(2, 3, 4))
	})

	t.Run("Intersecting a scaled sphere with a ray", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}

		s := GetSphere()

		s.SetTransform(datatypes.GetScaling(2, 2, 2))
		xs := Intersect(s, r)

		datatypes.AssertVal(t, float64(len(xs)), 2)
		datatypes.AssertVal(t, xs[0].T, 3)
		datatypes.AssertVal(t, xs[1].T, 7)
	})

	t.Run("Intersecting a translated sphere with a ray", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}

		s := GetSphere()

		s.SetTransform(datatypes.GetTranslation(5, 0, 0))
		xs := Intersect(s, r)

		datatypes.AssertVal(t, float64(len(xs)), 0)
	})

	t.Run("The normal on a sphere at a point on the x axis", func(t *testing.T) {
		s := GetSphere()

		n := s.Normal(datatypes.Point(1, 0, 0))
		datatypes.AssertTupleEqual(t, n, datatypes.Vector(1, 0, 0))
	})

	t.Run("The normal on a sphere at a point on the y axis", func(t *testing.T) {
		s := GetSphere()

		n := s.Normal(datatypes.Point(0, 1, 0))
		datatypes.AssertTupleEqual(t, n, datatypes.Vector(0, 1, 0))
	})

	t.Run("The normal on a sphere at a point on the z axis", func(t *testing.T) {
		s := GetSphere()

		n := s.Normal(datatypes.Point(0, 0, 1))
		datatypes.AssertTupleEqual(t, n, datatypes.Vector(0, 0, 1))
	})

	t.Run("The normal on a sphere at a noaxial point", func(t *testing.T) {
		s := GetSphere()

		n := s.Normal(datatypes.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
		datatypes.AssertTupleEqual(t, n, datatypes.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	})

	t.Run("The normal is a normalized vector", func(t *testing.T) {
		s := GetSphere()

		n := s.Normal(datatypes.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
		datatypes.AssertTupleEqual(t, n.Normalize(), n)
	})

	t.Run("Computing the normal on a translated sphere", func(t *testing.T) {
		s := GetSphere()
		s.SetTransform(datatypes.GetTranslation(0, 1, 0))

		n := Normal(s, datatypes.Point(0, 1.70711, -.70711))
		datatypes.AssertTupleEqual(t, n, datatypes.Vector(0, 0.70711, -.70711))
	})

	t.Run("Computing the normal on a transformed sphere", func(t *testing.T) {
		s := GetSphere()
		s.SetTransform(datatypes.Multiply(datatypes.GetScaling(1, 0.5, 1), datatypes.GetRotationZ(math.Pi/5)))

		n := Normal(s, datatypes.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
		datatypes.AssertTupleEqual(t, n, datatypes.Vector(0, 0.97014, -0.24254))
	})

	t.Run("A sphere has a default material", func(t *testing.T) {
		s := GetSphere()
		m := raytracing.GetMaterial()

		if !reflect.DeepEqual(s.GetMaterial(), m) {
			t.Error("expected sphere material is not the default")
		}
	})

	t.Run("A sphere may be assigned a materials", func(t *testing.T) {
		s := GetSphere()
		m := raytracing.GetMaterial()

		m.Ambient = 1
		s.SetMaterial(m)

		if !reflect.DeepEqual(s.GetMaterial(), m) {
			t.Error("sphere material is not expected")
		}
	})

	t.Run("Precompute the state of an intersection", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}

		sphere := GetSphere()
		i := Intersection{T: 4, Object: sphere}

		comps := i.PrepareComputations(r, []Intersection{i})

		datatypes.AssertVal(t, comps.T, i.T)
		datatypes.AssertTupleEqual(t, comps.Point, datatypes.Point(0, 0, -1))
		datatypes.AssertTupleEqual(t, comps.Eyev, datatypes.Vector(0, 0, -1))
		datatypes.AssertTupleEqual(t, comps.Normalv, datatypes.Vector(0, 0, -1))

		if !reflect.DeepEqual(comps.Object, sphere) {
			t.Error("spheres are not equal")
		}

		if comps.IsInside {
			t.Error("comp.IsInside is true, should be false")
		}

	})

	t.Run("The hit, when an intersection occurs on the outside", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, 0), Direction: datatypes.Vector(0, 0, 1)}

		sphere := GetSphere()
		i := Intersection{T: 1, Object: sphere}

		comps := i.PrepareComputations(r, []Intersection{i})

		if !comps.IsInside {
			t.Error("comp.IsInside is false, should be true")
		}

		datatypes.AssertTupleEqual(t, comps.Point, datatypes.Point(0, 0, 1))
		datatypes.AssertTupleEqual(t, comps.Eyev, datatypes.Vector(0, 0, -1))
		datatypes.AssertTupleEqual(t, comps.Normalv, datatypes.Vector(0, 0, -1))

		if !reflect.DeepEqual(comps.Object, sphere) {
			t.Error("spheres are not equal")
		}

	})

	t.Run("The hit should offset the point", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}

		sphere := GetSphere()
		sphere.SetTransform(datatypes.GetTranslation(0, 0, 1))

		i := Intersection{T: 5, Object: sphere}

		comps := i.PrepareComputations(r, []Intersection{i})

		if comps.OverPoint.Z >= -datatypes.EPSILON/2 {
			t.Error("over point is larger than expected")
		}

		if comps.Point.Z <= comps.OverPoint.Z {
			t.Error("over point is not larger than original point")
		}
	})

	t.Run("A helper for producing a sphere with a glassy material", func(t *testing.T) {
		s := GetGlassSphere()
		datatypes.AssertMatrixEqual(t, s.GetTransform(), datatypes.GetIdentity())

		material := s.GetMaterial()
		datatypes.AssertVal(t, material.Transparency, 1.0)
		datatypes.AssertVal(t, material.RefractiveIndex, 1.5)
	})

	t.Run("Finding n1 and n2 at various intersections", func(t *testing.T) {
		A := GetGlassSphere()
		A.SetTransform(datatypes.GetScaling(2, 2, 2))

		B := GetGlassSphere()
		B.SetTransform(datatypes.GetTranslation(0, 0, -0.25))
		material := B.GetMaterial()
		material.RefractiveIndex = 2.0
		B.SetMaterial(material)

		C := GetGlassSphere()
		C.SetTransform(datatypes.GetTranslation(0, 0, 0.25))
		material = C.GetMaterial()
		material.RefractiveIndex = 2.5
		C.SetMaterial(material)

		r := datatypes.Ray{Origin: datatypes.Point(0, 0, -4), Direction: datatypes.Vector(0, 0, 1)}

		xs := []Intersection{Intersection{2, A},
			Intersection{2.75, B},
			Intersection{3.25, C},
			Intersection{4.75, B},
			Intersection{5.25, C},
			Intersection{6, A}}

		n1 := []float64{1.0, 1.5, 2.0, 2.5, 2.5, 1.5}
		n2 := []float64{1.5, 2.0, 2.5, 2.5, 1.5, 1.0}

		for index := range n1 {
			comps := xs[index].PrepareComputations(r, xs)
			datatypes.AssertVal(t, comps.N1, n1[index])
			datatypes.AssertVal(t, comps.N2, n2[index])
		}
	})

	t.Run("The under point is offset below the surface", func(t *testing.T) {
		sphere := GetGlassSphere()
		sphere.SetTransform(datatypes.GetTranslation(0, 0, 1))
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}

		i := Intersection{5, sphere}

		comps := i.PrepareComputations(r, []Intersection{i})

		if comps.UnderPoint.Z <= datatypes.EPSILON/2 {
			t.Error("Expected under point z to be > EPSILON/2")
		}
		if comps.Point.Z >= comps.UnderPoint.Z {
			t.Error("Expected point z to be < under point z")
		}

	})
}
