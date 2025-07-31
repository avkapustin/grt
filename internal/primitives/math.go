package primitives

import "math"

// This is valid only for relative small float values
func EqualWithEps(a, b float32) bool {
	eps := 0.000001
	af := float64(a)
	bf := float64(b)

	if math.IsNaN(af) || math.IsNaN(bf) {
		return false
	}

	return math.Abs(af-bf) < eps
}
