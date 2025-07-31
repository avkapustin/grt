package primitives_test

import (
	"testing"

	"github.com/avkapustin/grt/internal/primitives"
)

func TestMatrixMul(t *testing.T) {
	maS := `| 1 | 2 | 3 | 4 |
| 2 | 3 | 4 | 5 |
| 3 | 4 | 5 | 6 |
| 4 | 5 | 6 | 7 |`

	mbS := `| 0 | 1 | 2 | 4 |
| 1 | 2 | 4 | 8 |
| 2 | 4 | 8 | 16 |
| 4 | 8 | 16 | 32 |`

	ma, err := primitives.MatrixFromString(maS)
	if err != nil {
		t.Fatal(err)
	}
	mb, err := primitives.MatrixFromString(mbS)
	if err != nil {
		t.Fatal(err)
	}

	actual := ma.Mul(mb)
	if actual.M10 != 31 {
		t.Error(ma, mb, actual)
	}

	p := primitives.MakePoint(1, 2, 3)
	actualP := ma.MulTuple(p)
	if actualP.X != 18 {
		t.Error(ma, p, actual)
	}
}

func TestInverseMatrix(t *testing.T) {
	mS := `| 8 | -5 | 9 | 2 |
| 7 | 5 | 6 | 1 |
| -6 | 0 | 9 | 6 |
| -3 | 0 | -9 | -4 |`

	ma, err := primitives.MatrixFromString(mS)
	if err != nil {
		t.Fatal(err)
	}

	inv, err := ma.InverseMatrix()
	if err != nil {
		t.Fatal(err)
	}

	if !primitives.EqualWithEps(inv.M00, -0.15384616) {
		t.Error(ma, inv, err, inv.M00)
	}

	mb := ma.Mul(inv)
	if !primitives.EqualWithEps(mb.M00, 1.0) {
		t.Error(mb)
	}
	if !primitives.EqualWithEps(mb.M10, 0.0) {
		t.Error(mb)
	}
}
