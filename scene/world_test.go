package scene

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
	"reflect"
	"testing"
)

func TestWorld(t *testing.T) {

	t.Run("the default world", func(t *testing.T) {
		light := raytracing.PointLight{Intensity: raytracing.Color{Red: 1, Green: 1, Blue: 1}, Position: datatypes.Point(-10, 10, -10)}
		s1 := raytracing.GetSphere()

		mat := s1.GetMaterial()
		mat.Color = raytracing.Color{Red: 0.8, Green: 1.0, Blue: 0.6}
		mat.Diffuse = 0.7
		mat.Specular = 0.2
		s1.SetMaterial(mat)

		s2 := raytracing.GetSphere()
		s2.SetTransform(datatypes.GetScaling(0.5, 0.5, 0.5))

		w := GetWorld()

		if !reflect.DeepEqual(w.Light, light) {
			t.Error("expected lights are not equal")
		}

		if !reflect.DeepEqual(w.Shapes[0], s1) {
			t.Error("expected shapes are not equal")
		}

		if !reflect.DeepEqual(w.Shapes[1], s2) {
			t.Error("expected shapes are not equal")
		}

	})

	t.Run("intersect a world with a ray", func(t *testing.T) {
		w := GetWorld()
		r := raytracing.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}

		xs := w.Intersect(r)

		datatypes.AssertVal(t, float64(len(xs)), 4)
		datatypes.AssertVal(t, xs[0].T, 4)
		datatypes.AssertVal(t, xs[1].T, 4.5)
		datatypes.AssertVal(t, xs[2].T, 5.5)
		datatypes.AssertVal(t, xs[3].T, 6)
	})

	t.Run("shading an intersection", func(t *testing.T) {
		w := GetWorld()
		r := raytracing.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}

		sphere := w.Shapes[0]
		i := raytracing.Intersection{T: 4, Object: sphere}

		comps := i.PrepareComputations(r)

		c := w.ShadeHit(comps, 5)

		raytracing.AssertColorsEqual(t, c, raytracing.Color{Red: 0.38066, Green: 0.47583, Blue: 0.2855})
	})

	t.Run("shading an intersection from the inside", func(t *testing.T) {
		w := GetWorld()
		w.Light = raytracing.PointLight{Intensity: raytracing.Color{Red: 1, Green: 1, Blue: 1}, Position: datatypes.Point(0, 0.25, 0)}
		r := raytracing.Ray{Origin: datatypes.Point(0, 0, 0), Direction: datatypes.Vector(0, 0, 1)}

		sphere := w.Shapes[1]
		i := raytracing.Intersection{T: 0.5, Object: sphere}

		comps := i.PrepareComputations(r)
		c := w.ShadeHit(comps, 5)

		raytracing.AssertColorsEqual(t, c, raytracing.Color{Red: 0.90498, Green: 0.90498, Blue: 0.90498})
	})

	t.Run("The color when a ray misses", func(t *testing.T) {
		w := GetWorld()
		r := raytracing.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 1, 0)}

		c := w.ColorAt(r, 5)
		raytracing.AssertColorsEqual(t, c, raytracing.Color{Red: 0, Green: 0, Blue: 0})
	})

	t.Run("The color when a ray hits", func(t *testing.T) {
		w := GetWorld()
		r := raytracing.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}

		c := w.ColorAt(r, 5)
		raytracing.AssertColorsEqual(t, c, raytracing.Color{Red: 0.38066, Green: 0.47583, Blue: 0.2855})
	})

	t.Run("The color with an intersection behind the ray", func(t *testing.T) {
		w := GetWorld()
		mat := w.Shapes[0].GetMaterial()
		mat.Ambient = 1
		w.Shapes[0].SetMaterial(mat)
		w.Shapes[1].SetMaterial(mat)

		r := raytracing.Ray{Origin: datatypes.Point(0, 0, 0.75), Direction: datatypes.Vector(0, 0, -1)}

		c := w.ColorAt(r, 5)
		raytracing.AssertColorsEqual(t, c, w.Shapes[1].GetMaterial().Color)
	})

	t.Run("There is no shadow when nothing is colinear with point and light", func(t *testing.T) {
		w := GetWorld()
		p := datatypes.Point(0, 10, 0)

		if w.IsShadowed(p) {
			t.Error("expected IsShadowed to return false")
		}
	})

	t.Run("The shadow when an object is between the point and the light", func(t *testing.T) {
		w := GetWorld()
		p := datatypes.Point(10, -10, 10)

		if !w.IsShadowed(p) {
			t.Error("expected IsShadowed to return true")
		}
	})

	t.Run("There is no shadow when an object is behind the light", func(t *testing.T) {
		w := GetWorld()
		p := datatypes.Point(-20, 20, -20)

		if w.IsShadowed(p) {
			t.Error("expected IsShadowed to return false")
		}
	})

	t.Run("There is no shadow when an object is behind the light", func(t *testing.T) {
		w := GetWorld()
		p := datatypes.Point(-2, 2, -2)

		if w.IsShadowed(p) {
			t.Error("expected IsShadowed to return false")
		}
	})

	t.Run("Shade hit correctly shades shadows", func(t *testing.T) {
		w := GetWorld()
		w.Light = raytracing.PointLight{Intensity: raytracing.Color{Red: 1, Green: 1, Blue: 1}, Position: datatypes.Point(0, 0, -10)}

		s1 := raytracing.GetSphere()
		s2 := raytracing.GetSphere()
		s2.SetTransform(datatypes.GetTranslation(0, 0, 10))

		w.Shapes = []raytracing.Shape{s1, s2}

		r := raytracing.Ray{Origin: datatypes.Point(0, 0, 5), Direction: datatypes.Vector(0, 0, 1)}
		i := raytracing.Intersection{T: 4, Object: s2}

		comps := i.PrepareComputations(r)
		c := w.ShadeHit(comps, 5)

		raytracing.AssertColorsEqual(t, c, raytracing.Color{Red: 0.1, Green: 0.1, Blue: 0.1})
	})

	t.Run("The reflected color for a nonreflective material", func(t *testing.T) {
		w := GetWorld()
		r := raytracing.Ray{Origin: datatypes.Point(0, 0, 0), Direction: datatypes.Vector(0, 0, 1)}
		shape := w.Shapes[1]
		material := shape.GetMaterial()
		material.Ambient = 1
		shape.SetMaterial(material)

		i := raytracing.Intersection{T: 1, Object: shape}
		comps := i.PrepareComputations(r)

		color := w.ReflectedColor(comps, 5)
		raytracing.AssertColorsEqual(t, color, raytracing.Color{Red: 0, Green: 0, Blue: 0})
	})

	t.Run("The reflected color for a reflective material", func(t *testing.T) {
		w := GetWorld()

		shape := raytracing.GetPlane()
		mat := shape.GetMaterial()
		mat.Reflective = 0.5
		shape.SetMaterial(mat)
		shape.SetTransform(datatypes.GetTranslation(0, -1, 0))

		w.Shapes = append(w.Shapes, shape)

		r := raytracing.Ray{Origin: datatypes.Point(0, 0, -3), Direction: datatypes.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
		i := raytracing.Intersection{T: math.Sqrt(2), Object: shape}

		comps := i.PrepareComputations(r)
		color := w.ReflectedColor(comps, 5)

		raytracing.AssertColorsEqual(t, color, raytracing.Color{Red: 0.19033, Green: 0.23791, Blue: 0.14274})
	})

	t.Run("ShadeHit with a reflective material", func(t *testing.T) {
		w := GetWorld()

		shape := raytracing.GetPlane()
		mat := shape.GetMaterial()
		mat.Reflective = 0.5
		shape.SetMaterial(mat)
		shape.SetTransform(datatypes.GetTranslation(0, -1, 0))

		w.Shapes = append(w.Shapes, shape)

		r := raytracing.Ray{Origin: datatypes.Point(0, 0, -3), Direction: datatypes.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
		i := raytracing.Intersection{T: math.Sqrt(2), Object: shape}

		comps := i.PrepareComputations(r)
		color := w.ShadeHit(comps, 5)

		raytracing.AssertColorsEqual(t, color, raytracing.Color{Red: 0.87675, Green: 0.92434, Blue: 0.82917})
	})

	t.Run("ColorAt with mutually reflective surfaces", func(t *testing.T) {
		w := GetWorld()
		w.Light = raytracing.PointLight{Intensity: raytracing.Color{Red: 1, Green: 1, Blue: 1}, Position: datatypes.Point(0, 0, 0)}

		lower := raytracing.GetPlane()
		mat := lower.GetMaterial()
		mat.Reflective = 1
		lower.SetMaterial(mat)
		//lower.SetTransform(datatypes.GetTranslation(0, -1, 0))

		upper := raytracing.GetPlane()
		mat = upper.GetMaterial()
		mat.Reflective = 1
		upper.SetMaterial(mat)
		upper.SetTransform(datatypes.GetTranslation(0, 1, 0))

		w.Shapes = []raytracing.Shape{upper, lower}

		r := raytracing.Ray{Origin: datatypes.Point(0, 0, 0), Direction: datatypes.Vector(0, 1, 0)}
		w.ColorAt(r, 5)
	})

	t.Run("The reflected color at the maximum recursive depth", func(t *testing.T) {
		w := GetWorld()

		shape := raytracing.GetPlane()
		mat := shape.GetMaterial()
		mat.Reflective = 0.5
		shape.SetMaterial(mat)
		shape.SetTransform(datatypes.GetTranslation(0, -1, 0))

		w.Shapes = append(w.Shapes, shape)

		r := raytracing.Ray{Origin: datatypes.Point(0, 0, -3), Direction: datatypes.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
		i := raytracing.Intersection{T: math.Sqrt(2), Object: shape}

		comps := i.PrepareComputations(r)
		color := w.ReflectedColor(comps, 0)

		raytracing.AssertColorsEqual(t, color, raytracing.Color{Red: 0, Green: 0, Blue: 0})
	})

}
