package primitives_test

import (
	"testing"

	p "github.com/avkapustin/grt/internal/primitives"
	"github.com/stretchr/testify/assert"
)

func checkColorEquality(t *testing.T, expected, actual p.Color) {
	assert.Equalf(t, expected.R, actual.R,
		"expected tuple %#v, actual tuple %#v, field R, expected %d, actual %d", expected, actual, expected.R, actual.R)
	assert.Equalf(t, expected.G, actual.G,
		"expected tuple %#v, actual tuple %#v, field G, expected %d, actual %d", expected, actual, expected.G, actual.G)
	assert.Equalf(t, expected.B, actual.B,
		"expected tuple %#v, actual tuple %#v, field B, expected %d, actual %d", expected, actual, expected.B, actual.B)
	assert.Equalf(t, expected.A, actual.A,
		"expected tuple %#v, actual tuple %#v, field A, expected %d, actual %d", expected, actual, expected.A, actual.A)
}

func TestColorMath(t *testing.T) {
	c1 := p.Color{200, 200, 200, 0}
	c2 := p.Color{200, 200, 200, 0}

	// check for cap max value instead of rotating
	actual := c1.Add(c2)
	expected := p.Color{255, 255, 255, 0}

	checkColorEquality(t, expected, actual)

	scaled := c1.Scale(10.0)

	checkColorEquality(t, expected, scaled)
}
