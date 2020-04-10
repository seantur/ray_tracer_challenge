package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/canvas"
	"reflect"
	"testing"
)

func TestMaterials(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	t.Run("Test the default material", func(t *testing.T) {
		m := GetMaterial()

		if !reflect.DeepEqual(m.Color, canvas.Color{Red: 1, Green: 1, Blue: 1}) {
			t.Error("Expected default material color did not match")
		}

		assertVal(t, m.Ambient, 0.1)
		assertVal(t, m.Diffuse, 0.9)
		assertVal(t, m.Specular, 0.9)
		assertVal(t, m.Shininess, 200.0)
	})

	t.Run("Can set the values of materials", func(t *testing.T) {
		m := GetMaterial()

		m.Color = canvas.Color{Red: 0.5, Green: 0.5, Blue: 0.5}
		m.Ambient = 0.2
		m.Diffuse = 0.2
		m.Specular = 0.2
		m.Shininess = 500.0

		if !reflect.DeepEqual(m.Color, canvas.Color{Red: 0.5, Green: 0.5, Blue: 0.5}) {
			t.Error("Expected default material color did not match")
		}

		assertVal(t, m.Ambient, 0.2)
		assertVal(t, m.Diffuse, 0.2)
		assertVal(t, m.Specular, 0.2)
		assertVal(t, m.Shininess, 500.0)
	})

}
