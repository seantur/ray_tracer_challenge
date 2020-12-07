package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"reflect"
	"testing"
)

func TestMaterials(t *testing.T) {

	t.Run("Test the default material", func(t *testing.T) {
		m := GetMaterial()

		if !reflect.DeepEqual(m.RGB, RGB{Red: 1, Green: 1, Blue: 1}) {
			t.Error("Expected default material color did not match")
		}

		datatypes.AssertVal(t, m.Ambient, 0.1)
		datatypes.AssertVal(t, m.Diffuse, 0.9)
		datatypes.AssertVal(t, m.Specular, 0.9)
		datatypes.AssertVal(t, m.Shininess, 200.0)
	})

	t.Run("Can set the values of materials", func(t *testing.T) {
		m := GetMaterial()

		m.RGB = RGB{Red: 0.5, Green: 0.5, Blue: 0.5}
		m.Ambient = 0.2
		m.Diffuse = 0.2
		m.Specular = 0.2
		m.Shininess = 500.0

		if !reflect.DeepEqual(m.RGB, RGB{Red: 0.5, Green: 0.5, Blue: 0.5}) {
			t.Error("Expected default material color did not match")
		}

		datatypes.AssertVal(t, m.Ambient, 0.2)
		datatypes.AssertVal(t, m.Diffuse, 0.2)
		datatypes.AssertVal(t, m.Specular, 0.2)
		datatypes.AssertVal(t, m.Shininess, 500.0)
	})

	t.Run("Default material has no reflectivity", func(t *testing.T) {
		m := GetMaterial()
		datatypes.AssertVal(t, m.Reflective, 0.0)
	})

	t.Run("Transparency and Refractive Index for the default material", func(t *testing.T) {
		m := GetMaterial()
		datatypes.AssertVal(t, m.Transparency, 0.0)
		datatypes.AssertVal(t, m.RefractiveIndex, 1.0)
	})
}
