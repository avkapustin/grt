package matrix_test

import (
	"testing"

	m "github.com/avkapustin/grt/internal/matrix"
	p "github.com/avkapustin/grt/internal/primitives"

	"github.com/stretchr/testify/assert"
)

func checkMatrixEquality(t *testing.T, expected m.Matrix4, actual m.Matrix4) {
	assert.InDeltaf(t, expected.M00, actual.M00, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M00, expected %.6f, actual %.6f",
		expected, actual, expected.M00, actual.M00)

	assert.InDeltaf(t, expected.M10, actual.M10, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M10, expected %.6f, actual %.6f",
		expected, actual, expected.M10, actual.M10)

	assert.InDeltaf(t, expected.M20, actual.M20, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M20, expected %.6f, actual %.6f",
		expected, actual, expected.M20, actual.M20)

	assert.InDeltaf(t, expected.M30, actual.M30, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M30, expected %.6f, actual %.6f",
		expected, actual, expected.M30, actual.M30)

	assert.InDeltaf(t, expected.M01, actual.M01, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M01, expected %.6f, actual %.6f",
		expected, actual, expected.M01, actual.M01)

	assert.InDeltaf(t, expected.M11, actual.M11, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M11, expected %.6f, actual %.6f",
		expected, actual, expected.M11, actual.M11)

	assert.InDeltaf(t, expected.M21, actual.M21, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M21, expected %.6f, actual %.6f",
		expected, actual, expected.M21, actual.M21)

	assert.InDeltaf(t, expected.M31, actual.M31, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M31, expected %.6f, actual %.6f",
		expected, actual, expected.M31, actual.M31)

	assert.InDeltaf(t, expected.M02, actual.M02, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M02, expected %.6f, actual %.6f",
		expected, actual, expected.M02, actual.M02)

	assert.InDeltaf(t, expected.M12, actual.M12, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M12, expected %.6f, actual %.6f",
		expected, actual, expected.M12, actual.M12)

	assert.InDeltaf(t, expected.M22, actual.M22, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M22, expected %.6f, actual %.6f",
		expected, actual, expected.M22, actual.M22)

	assert.InDeltaf(t, expected.M32, actual.M32, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M32, expected %.6f, actual %.6f",
		expected, actual, expected.M32, actual.M32)

	assert.InDeltaf(t, expected.M03, actual.M03, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M03, expected %.6f, actual %.6f",
		expected, actual, expected.M03, actual.M03)

	assert.InDeltaf(t, expected.M13, actual.M13, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M13, expected %.6f, actual %.6f",
		expected, actual, expected.M13, actual.M13)

	assert.InDeltaf(t, expected.M23, actual.M23, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M23, expected %.6f, actual %.6f",
		expected, actual, expected.M23, actual.M23)

	assert.InDeltaf(t, expected.M33, actual.M33, 0.00001,
		"expected matrix %s\nactual matrix %s\nfield M33, expected %.6f, actual %.6f",
		expected, actual, expected.M33, actual.M33)
}

func TestMatrixMul(t *testing.T) {
	ma, err := m.MatrixFromString(`| 1 | 2 | 3 | 4 |
| 2 | 3 | 4 | 5 |
| 3 | 4 | 5 | 6 |
| 4 | 5 | 6 | 7 |`)
	if err != nil {
		t.Fatal(err)
	}

	mb, err := m.MatrixFromString(`| 0 | 1 | 2 | 4 |
| 1 | 2 | 4 | 8 |
| 2 | 4 | 8 | 16 |
| 4 | 8 | 16 | 32 |`)
	if err != nil {
		t.Fatal(err)
	}

	actual := ma.Mul(mb)
	assert.InDelta(t, 31, actual.M10, 0.00001)

	point := p.MakePoint(1, 2, 3)
	actualP := ma.MulTuple(point)
	assert.InDelta(t, 18, actualP.X, 0.00001)
}

func TestInverseMatrix(t *testing.T) {
	ma, err := m.MatrixFromString(`| 8 | -5 | 9 | 2 |
| 7 | 5 | 6 | 1 |
| -6 | 0 | 9 | 6 |
| -3 | 0 | -9 | -4 |`)
	if err != nil {
		t.Fatal(err)
	}

	inv, err := ma.InverseMatrix()
	if err != nil {
		t.Fatal(err)
	}

	// inv matrix should looks like
	mInvExpected, err := m.MatrixFromString(`| -0.15385 | -0.15385 | -0.28205 | -0.53846 |
| -0.07692 | 0.12308 | 0.02564 | 0.03077 |
| 0.35897 | 0.35897 | 0.43590 | 0.92308 |
| -0.69231 | -0.69231 | -0.76923 | -1.92308 |`)
	if err != nil {
		t.Fatal(err)
	}

	checkMatrixEquality(t, mInvExpected, inv)
	// A * Inv(A) = Identity matrix
	mb := ma.Mul(inv)

	checkMatrixEquality(t, m.IdentityMatrix(), mb)
}

func TestFastSRTInverseMatrix(t *testing.T) {
	ma := m.RotateXMatrix(3.14 / 4)
	mi, err := ma.FastInverseSRTMatrix()
	if err != nil {
		t.Fatal(err)
	}

	mb := ma.Mul(mi)

	checkMatrixEquality(t, m.IdentityMatrix(), mb)

	miClassic, err := ma.InverseMatrix()
	if err != nil {
		t.Fatal(err)
	}

	checkMatrixEquality(t, mi, miClassic)
}
