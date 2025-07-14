package primitives_test

import (
	"testing"

	"github.com/avkapustin/grt/internal/primitives"
)

func TestMatrixRepr(t *testing.T) {
	matrixS := `| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 10 | 11 | 12 |
| 13 | 14 | 15 | 16 |`
	actual, err := primitives.MatrixFromString(matrixS)

	if err != nil {
		t.Fatal(err)
	}

	if actual.M33 != 16 {
		t.Error(actual)
	}
	if actual.M30 != 13 {
		t.Error(actual)
	}
}

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
