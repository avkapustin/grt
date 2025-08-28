package primitives_test

import (
	"testing"

	p "github.com/avkapustin/grt/internal/primitives"
)

func TestCanvas(t *testing.T) {
	c, err := p.MakeCanvas(512, 512)
	if err != nil {
		t.Fatalf("Cannot create canvas %s\n", err.Error())
	}

	expected := p.Color{100, 100, 100, 0}

	c.SetScreenPixel(511, 511, expected)
	actual := c.GetScreenPixel(511, 511)

	checkColorEquality(t, expected, actual)
}

type coordCases struct {
	name             string
	setC             p.Color
	expectedC        p.Color
	viewX, viewY     int
	screenX, screenY int
}

func TestCoordTranslation(t *testing.T) {
	c, err := p.MakeCanvas(512, 512)
	if err != nil {
		t.Fatalf("Cannot create canvas %s\n", err.Error())
	}

	tests := []coordCases{
		{"regular", p.Color{122, 122, 122, 0}, p.Color{122, 122, 122, 0}, -100, 100, 156, 156},
		{"out of screen", p.Color{122, 122, 122, 0}, p.Color{0, 0, 0, 0}, -257, 100, 0, 156},
		{"zero/zero", p.Color{122, 122, 122, 0}, p.Color{122, 122, 122, 0}, -256, 256, 0, 0},
		{"far right down", p.Color{122, 122, 122, 0}, p.Color{122, 122, 122, 0}, 255, -255, 511, 511},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c.SetViewportPixel(tc.viewX, tc.viewY, tc.setC)
			actual := c.GetScreenPixel(tc.screenX, tc.screenY)
			checkColorEquality(t, tc.expectedC, actual)
		})
	}
}
