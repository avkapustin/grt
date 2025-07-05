package primitives_test

import (
	"testing"

	"github.com/avkapustin/grt/internal/primitives"
)

func isColorsEqual(a, b primitives.Color) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B
}

func TestColorMath(t *testing.T) {
	c1 := primitives.Color{200, 200, 200, 0}
	c2 := primitives.Color{200, 200, 200, 0}

	// check for cap max value instead of rotating
	actual := c1.Add(c2)
	expected := primitives.Color{255, 255, 255, 0}

	if !isColorsEqual(actual, expected) {
		t.Errorf("Error when comparing colors, actual %#v, expected %#v\n", actual, expected)
	}

	scaled := c1.Scale(10.0)

	if !isColorsEqual(scaled, expected) {
		t.Errorf("Error when scale color, actual %#v, expected %#v\n", scaled, expected)
	}
}
