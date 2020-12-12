package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
	"testing"
)

func TestGroups(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if !datatypes.IsClose(got, want) {
			t.Errorf("got %f want %f", got, want)
		}

	}

	assertShape := func(t *testing.T, got Shape, want Shape) {
		t.Helper()
		if got != want {
			t.Errorf("shapes did not match: got %v want %v", got, want)
		}
	}

	t.Run("Creating a new group", func(t *testing.T) {
		g := GetGroup()

		datatypes.AssertMatrixEqual(t, g.GetTransform(), datatypes.GetIdentity())
		assertVal(t, float64(len(g.Shapes)), 0)
	})

	t.Run("Adding a child to a group", func(t *testing.T) {
		g := GetGroup()
		s := GetSphere()

		g.AddChild(s)

		assertVal(t, float64(len(g.Shapes)), 1)

		if s.GetParent() != g {
			t.Error("Expected sphere parent to be group")
		}

		if g.Shapes[0] != s {
			t.Error("Expected shape to be in group")
		}

	})

	t.Run("Intersecting a ray with an empty group", func(t *testing.T) {
		g := GetGroup()
		r := datatypes.Ray{datatypes.Point(0, 0, 0), datatypes.Vector(0, 0, 1)}

		xs := g.Intersect(r)

		assertVal(t, float64(len(xs)), 0)
	})

	t.Run("Intersecting a ray with a nonempty group", func(t *testing.T) {
		g := GetGroup()

		s1 := GetSphere()
		s2 := GetSphere()
		s2.SetTransform(datatypes.GetTranslation(0, 0, -3))
		s3 := GetSphere()
		s3.SetTransform(datatypes.GetTranslation(5, 0, 0))

		g.AddChild(s1)
		g.AddChild(s2)
		g.AddChild(s3)

		r := datatypes.Ray{datatypes.Point(0, 0, -5), datatypes.Vector(0, 0, 1)}

		xs := g.Intersect(r)

		if len(xs) != 4 {
			t.Fatalf("expected 4 intersection got %v", len(xs))
		}

		assertShape(t, xs[0].Object, s2)
		assertShape(t, xs[1].Object, s2)
		assertShape(t, xs[2].Object, s1)
		assertShape(t, xs[3].Object, s1)
	})

	t.Run("Intersecting a transformed group", func(t *testing.T) {
		g := GetGroup()
		g.SetTransform(datatypes.GetScaling(2, 2, 2))

		s := GetSphere()
		s.SetTransform(datatypes.GetTranslation(5, 0, 0))
		g.AddChild(s)

		r := datatypes.Ray{datatypes.Point(10, 0, -10), datatypes.Vector(0, 0, 1)}
		xs := Intersect(g, r)

		if len(xs) != 2 {
			t.Errorf("Expected len(xs) == %v got %v", 2, len(xs))
		}
	})

	t.Run("Converting a point from world to object space", func(t *testing.T) {
		g1 := GetGroup()
		g1.SetTransform(datatypes.GetRotationY(math.Pi / 2))

		g2 := GetGroup()
		g2.SetTransform(datatypes.GetScaling(2, 2, 2))
		g1.AddChild(g2)

		s := GetSphere()
		s.SetTransform(datatypes.GetTranslation(5, 0, 0))
		g2.AddChild(s)

		p := WorldToObject(s, datatypes.Point(-2, 0, -10))

		datatypes.AssertTupleEqual(t, p, datatypes.Point(0, 0, -1))

	})

	t.Run("Converting an normal from object to world space", func(t *testing.T) {
		g1 := GetGroup()
		g1.SetTransform(datatypes.GetRotationY(math.Pi / 2))

		g2 := GetGroup()
		g2.SetTransform(datatypes.GetScaling(1, 2, 3))
		g1.AddChild(g2)

		s := GetSphere()
		s.SetTransform(datatypes.GetTranslation(5, 0, 0))
		g2.AddChild(s)

		n := NormalToWorld(s, datatypes.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

		datatypes.AssertTupleEqual(t, n, datatypes.Vector(0.28571, 0.42857, -0.85714))
	})

	t.Run("Finding the normal on a child object", func(t *testing.T) {
		g1 := GetGroup()
		g1.SetTransform(datatypes.GetRotationY(math.Pi / 2))

		g2 := GetGroup()
		g2.SetTransform(datatypes.GetScaling(1, 2, 3))
		g1.AddChild(g2)

		s := GetSphere()
		s.SetTransform(datatypes.GetTranslation(5, 0, 0))
		g2.AddChild(s)

		n := NormalAt(s, datatypes.Point(1.7321, 1.1547, -5.5774))
		datatypes.AssertTupleEqual(t, n, datatypes.Vector(0.28570, 0.42854, -0.85716))
	})

}
