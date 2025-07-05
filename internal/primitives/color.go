package primitives

import "math"

// color representation choices
//
// r, g, b, a float32 - between 0..1, map to bytes then displays to screen/file
// contra - double representation, 3x memory (color can be represented as 24/32bit)
// pro - math without rounding errors, no overflow, capping only then go to screen
//
// r, g, b, a uint8 - between 0..255
// contra - extra round errors in intensive math (scale, blur, ...) + round on every operation
// contra - extra logic to prevent overflow on add/scale operation
// pro - less memory, simple representation to screen/file

// color representation in 0..255 interval
type Color struct {
	R, G, B, A uint8
}

func clamp(v uint32) uint8 {
	if v > 255 {
		return 255
	}

	return uint8(v & 0xFF)
}

func clampFloat(v float64) uint8 {
	if math.IsNaN(v) {
		return 0
	}
	if v > 255.0 {
		return 255
	}
	if v < 0.0 {
		return 0
	}

	return uint8(v)
}

func (c Color) Add(o Color) Color {
	// we need to perform math on extended values
	return Color{
		clamp(uint32(c.R) + uint32(o.R)),
		clamp(uint32(c.G) + uint32(o.G)),
		clamp(uint32(c.B) + uint32(o.B)),
		0,
	}
}

func (c Color) Scale(s float64) Color {
	return Color{
		clampFloat(math.Round(float64(c.R) * s)),
		clampFloat(math.Round(float64(c.G) * s)),
		clampFloat(math.Round(float64(c.B) * s)),
		0,
	}
}
