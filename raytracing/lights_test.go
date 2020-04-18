package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
	"reflect"
	"testing"
)

func TestLight(t *testing.T) {

	assertColorsEqual := func(t *testing.T, got Color, want Color) {
		t.Helper()

		allClose := datatypes.IsClose(got.Red, want.Red) && datatypes.IsClose(got.Green, want.Green) && datatypes.IsClose(got.Blue, want.Blue)

		if !allClose {
			t.Error("wanted equal colors are not equal")
		}
	}

	t.Run("A point light has a position and intensity", func(t *testing.T) {
		intensity := Color{Red: 1, Green: 1, Blue: 1}
		position := datatypes.Point(0, 0, 0)

		light := PointLight{Position: position, Intensity: intensity}

		if !reflect.DeepEqual(intensity, light.Intensity) {
			t.Error("Did not get expected intensity")
		}

		if !reflect.DeepEqual(position, light.Position) {
			t.Error("Did not get expected position")
		}
	})

	t.Run("Lighting with the eye between the light and the surface", func(t *testing.T) {
		m := GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, 0, -1)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 0, -10), Intensity: Color{Red: 1, Green: 1, Blue: 1}}

		result := Lighting(m, light, p, eyev, normalv)

		assertColorsEqual(t, result, Color{Red: 1.9, Green: 1.9, Blue: 1.9})
	})

	t.Run("Lighting with the eye between the light and the surface, eye offset 45 deg", func(t *testing.T) {
		m := GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 0, -10), Intensity: Color{Red: 1, Green: 1, Blue: 1}}

		result := Lighting(m, light, p, eyev, normalv)

		assertColorsEqual(t, result, Color{Red: 1.0, Green: 1.0, Blue: 1.0})
	})

	t.Run("Lighting with eye opposite surface, light offset 45 deg", func(t *testing.T) {
		m := GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, 0, -1)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 10, -10), Intensity: Color{Red: 1, Green: 1, Blue: 1}}

		result := Lighting(m, light, p, eyev, normalv)

		assertColorsEqual(t, result, Color{Red: 0.7364, Green: 0.7364, Blue: 0.7364})
	})

	t.Run("Lighting with eye in the path of the reflection vector", func(t *testing.T) {
		m := GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 10, -10), Intensity: Color{Red: 1, Green: 1, Blue: 1}}

		result := Lighting(m, light, p, eyev, normalv)

		assertColorsEqual(t, result, Color{Red: 1.6364, Green: 1.6364, Blue: 1.6364})
	})

	t.Run("Lighting with the light behind the surface", func(t *testing.T) {
		m := GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, 0, -1)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 0, 10), Intensity: Color{Red: 1, Green: 1, Blue: 1}}

		result := Lighting(m, light, p, eyev, normalv)

		assertColorsEqual(t, result, Color{Red: 0.1, Green: 0.1, Blue: 0.1})
	})

}
