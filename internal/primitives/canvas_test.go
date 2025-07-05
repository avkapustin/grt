package primitives_test

import (
	"testing"

	"github.com/avkapustin/grt/internal/primitives"
)

func TestCanvas(t *testing.T) {
	c, err := primitives.MakeCanvas(512, 512)
	if err != nil {
		t.Fatalf("Cannot create canvas %s\n", err.Error())
	}

	expected := primitives.Color{100, 100, 100, 0}

	c.SetScreenPixel(511, 511, expected)
	actual := c.GetScreenPixel(511, 511)

	if !isColorsEqual(expected, actual) {
		t.Errorf("Incorrect color for pixel %#v, expected %#v\n", actual, expected)
	}
}

type coordCases struct {
	name             string
	setC             primitives.Color
	expectedC        primitives.Color
	viewX, viewY     int
	screenX, screenY int
}

func TestCoordTranslation(t *testing.T) {
	c, err := primitives.MakeCanvas(512, 512)
	if err != nil {
		t.Fatalf("Cannot create canvas %s\n", err.Error())
	}

	tests := []coordCases{
		{"regular", primitives.Color{122, 122, 122, 0}, primitives.Color{122, 122, 122, 0}, -100, 100, 156, 156},
		{"out of screen", primitives.Color{122, 122, 122, 0}, primitives.Color{0, 0, 0, 0}, -257, 100, 0, 156},
		{"zero/zero", primitives.Color{122, 122, 122, 0}, primitives.Color{122, 122, 122, 0}, -256, 256, 0, 0},
		{"far right down", primitives.Color{122, 122, 122, 0}, primitives.Color{122, 122, 122, 0}, 255, -255, 511, 511},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c.SetViewportPixel(tc.viewX, tc.viewY, tc.setC)
			actual := c.GetScreenPixel(tc.screenX, tc.screenY)
			if !isColorsEqual(actual, tc.expectedC) {
				t.Errorf("Falied case: %s. Incorrect color for pixel %#v, expected %#v\n", tc.name, actual, tc.expectedC)
			}
		})
	}
}
