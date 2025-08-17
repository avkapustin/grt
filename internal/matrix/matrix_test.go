package matrix_test

import (
	"testing"

	m "github.com/avkapustin/grt/internal/matrix"
	p "github.com/avkapustin/grt/internal/primitives"
)

func isMatrixEqual(a m.Matrix4, b m.Matrix4) bool {
	return p.EqualWithEps(a.M00, b.M00) &&
		p.EqualWithEps(a.M10, b.M10) &&
		p.EqualWithEps(a.M20, b.M20) &&
		p.EqualWithEps(a.M30, b.M30) &&
		p.EqualWithEps(a.M01, b.M01) &&
		p.EqualWithEps(a.M11, b.M11) &&
		p.EqualWithEps(a.M21, b.M21) &&
		p.EqualWithEps(a.M31, b.M31) &&
		p.EqualWithEps(a.M02, b.M02) &&
		p.EqualWithEps(a.M12, b.M12) &&
		p.EqualWithEps(a.M22, b.M22) &&
		p.EqualWithEps(a.M32, b.M32) &&
		p.EqualWithEps(a.M03, b.M03) &&
		p.EqualWithEps(a.M13, b.M13) &&
		p.EqualWithEps(a.M23, b.M23) &&
		p.EqualWithEps(a.M33, b.M33)
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
	if !p.EqualWithEps(actual.M10, 31) {
		t.Error(ma, mb, actual)
	}

	point := p.MakePoint(1, 2, 3)
	actualP := ma.MulTuple(point)
	if !p.EqualWithEps(actualP.X, 18) {
		t.Error(ma, point, actualP)
	}
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

	if !isMatrixEqual(inv, mInvExpected) {
		t.Errorf("Error in inv matrix, source \n%s\n, got \n%s\n, expected \n%s\n", ma, inv, mInvExpected)
	}

	// A * Inv(A) = Identity matrix
	mb := ma.Mul(inv)
	if !isMatrixEqual(mb, m.IdentityMatrix()) {
		t.Errorf("a * inv(a) == identity matrix, got %s", mb)
	}
}

func TestFastSRTInverseMatrix(t *testing.T) {
	ma := m.RotateXMatrix(3.14 / 4)
	mi, err := ma.FastInverseSRTMatrix()
	if err != nil {
		t.Fatal(err)
	}

	mb := ma.Mul(mi)
	if !isMatrixEqual(mb, m.IdentityMatrix()) {
		t.Errorf("a * inv(a) == identity matrix, got %s", mb)
	}

	miClassic, err := ma.InverseMatrix()
	if err != nil {
		t.Fatal(err)
	}

	if !isMatrixEqual(mi, miClassic) {
		t.Errorf("Inverse matrix for classic and fast SRT should be equal, fast \n%s\n, classic \n%s\n", mi, miClassic)
	}
}
