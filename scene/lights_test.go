package scene

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"github.com/seantur/ray_tracer_challenge/shapes"
	"math"
	"reflect"
	"testing"
)

func TestLight(t *testing.T) {

	sphere := shapes.GetSphere()

	t.Run("A point light has a position and intensity", func(t *testing.T) {
		intensity := raytracing.RGB{Red: 1, Green: 1, Blue: 1}
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
		m := raytracing.GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, 0, -1)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 0, -10), Intensity: raytracing.RGB{Red: 1, Green: 1, Blue: 1}}
		in_shadow := false

		result := Lighting(m, sphere, light, p, eyev, normalv, in_shadow)

		raytracing.AssertColorsEqual(t, result, raytracing.RGB{Red: 1.9, Green: 1.9, Blue: 1.9})
	})

	t.Run("Lighting with the eye between the light and the surface, eye offset 45 deg", func(t *testing.T) {
		m := raytracing.GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 0, -10), Intensity: raytracing.RGB{Red: 1, Green: 1, Blue: 1}}
		in_shadow := false

		result := Lighting(m, sphere, light, p, eyev, normalv, in_shadow)

		raytracing.AssertColorsEqual(t, result, raytracing.RGB{Red: 1.0, Green: 1.0, Blue: 1.0})
	})

	t.Run("Lighting with eye opposite surface, light offset 45 deg", func(t *testing.T) {
		m := raytracing.GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, 0, -1)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 10, -10), Intensity: raytracing.RGB{Red: 1, Green: 1, Blue: 1}}
		in_shadow := false

		result := Lighting(m, sphere, light, p, eyev, normalv, in_shadow)

		raytracing.AssertColorsEqual(t, result, raytracing.RGB{Red: 0.7364, Green: 0.7364, Blue: 0.7364})
	})

	t.Run("Lighting with eye in the path of the reflection vector", func(t *testing.T) {
		m := raytracing.GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 10, -10), Intensity: raytracing.RGB{Red: 1, Green: 1, Blue: 1}}
		in_shadow := false

		result := Lighting(m, sphere, light, p, eyev, normalv, in_shadow)

		raytracing.AssertColorsEqual(t, result, raytracing.RGB{Red: 1.6364, Green: 1.6364, Blue: 1.6364})
	})

	t.Run("Lighting with the light behind the surface", func(t *testing.T) {
		m := raytracing.GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, 0, -1)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 0, 10), Intensity: raytracing.RGB{Red: 1, Green: 1, Blue: 1}}
		in_shadow := false

		result := Lighting(m, sphere, light, p, eyev, normalv, in_shadow)

		raytracing.AssertColorsEqual(t, result, raytracing.RGB{Red: 0.1, Green: 0.1, Blue: 0.1})
	})

	t.Run("Lighting with the surface in shadow", func(t *testing.T) {
		m := raytracing.GetMaterial()
		p := datatypes.Point(0, 0, 0)

		eyev := datatypes.Vector(0, 0, -1)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 0, -1), Intensity: raytracing.RGB{Red: 1, Green: 1, Blue: 1}}
		in_shadow := true

		result := Lighting(m, sphere, light, p, eyev, normalv, in_shadow)
		raytracing.AssertColorsEqual(t, result, raytracing.RGB{Red: 0.1, Green: 0.1, Blue: 0.1})

	})

	t.Run("Lighting with a pattern applied", func(t *testing.T) {
		m := raytracing.GetMaterial()
		m.Pattern = raytracing.GetStripe(raytracing.RGB{Red: 1, Green: 1, Blue: 1}, raytracing.RGB{Red: 0, Green: 0, Blue: 0})
		m.Ambient = 1
		m.Diffuse = 0
		m.Specular = 0

		eyev := datatypes.Vector(0, 0, -1)
		normalv := datatypes.Vector(0, 0, -1)
		light := PointLight{Position: datatypes.Point(0, 0, -10), Intensity: raytracing.RGB{Red: 1, Green: 1, Blue: 1}}

		c1 := Lighting(m, sphere, light, datatypes.Point(0.9, 0, 0), eyev, normalv, false)
		c2 := Lighting(m, sphere, light, datatypes.Point(1.1, 0, 0), eyev, normalv, false)

		raytracing.AssertColorsEqual(t, c1, raytracing.RGB{Red: 1, Green: 1, Blue: 1})
		raytracing.AssertColorsEqual(t, c2, raytracing.RGB{Red: 0, Green: 0, Blue: 0})
	})

}
