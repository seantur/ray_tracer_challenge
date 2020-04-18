package scene

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"reflect"
	"testing"
)

func TestWorld(t *testing.T) {

	t.Run("the default world", func(t *testing.T) {
		light := raytracing.PointLight{Intensity: raytracing.Color{Red: 1, Green: 1, Blue: 1}, Position: datatypes.Point(-10, 10, -10)}
		s1 := raytracing.GetSphere()

		s1.Color = raytracing.Color{Red: 0.8, Green: 1.0, Blue: 0.6}
		s1.Diffuse = 0.7
		s1.Specular = 0.2

		s2 := raytracing.GetSphere()
		s2.Transform = datatypes.GetScaling(0.5, 0.5, 0.5)

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

}
