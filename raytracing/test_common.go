package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"testing"
)

func AssertColorsEqual(t *testing.T, got Color, want Color) {
	allClose := datatypes.IsClose(got.Red, want.Red) && datatypes.IsClose(got.Green, want.Green) && datatypes.IsClose(got.Blue, want.Blue)

	if !allClose {
		t.Error("wanted equal colors are not equal")
	}
}
