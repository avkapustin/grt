package primitives

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// matrix - column vs row  ordered in terms of performance
// store in column-major mode (OpenGL, raylib, ...)
// M<row><col> naming
// 16 elems x 4 byte = 64 byte - still valid to pass as value?
type Matrix4 struct {
	M00, M10, M20, M30 float32
	M01, M11, M21, M31 float32
	M02, M12, M22, M32 float32
	M03, M13, M23, M33 float32
}

// for test purposes
// from / to string representation
// support only 4x4 matricies
// format | 0 | 1 | 2 | 3 |
// will panic if anything goes wrong
func MatrixFromString(s string) (Matrix4, error) {
	rows := strings.Split(s, "\n")
	if len(rows) != 4 {
		return Matrix4{}, fmt.Errorf("Matrix4 from string - invalid number of rows, input: %s\n", s)
	}
	matrix := Matrix4{}
	val := reflect.ValueOf(&matrix).Elem()
	for ri, row := range rows {
		cols := strings.Split(row, "|")
		if len(cols) != 6 {
			return Matrix4{}, fmt.Errorf("Matrix4 from string - invalid number of cols in row, input: %s\n", s)
		}
		for ci, col := range cols[1:5] {
			fieldName := fmt.Sprintf("M%d%d", ri, ci)
			col, _ := strconv.ParseFloat(strings.TrimSpace(col), 64)
			val.FieldByName(fieldName).SetFloat(col)
		}
	}
	return matrix, nil
}

func (m Matrix4) String() string {
	var sb strings.Builder
	matrix := reflect.ValueOf(&m).Elem()

	for ri := range 4 {
		for ci := range 4 {
			fieldName := fmt.Sprintf("M%d%d", ri, ci)
			sb.WriteString(fmt.Sprintf("| %.5f ", matrix.FieldByName(fieldName).Float()))
		}
		sb.WriteString("|\n")
	}

	return sb.String()
}

func (m Matrix4) Mul(o Matrix4) Matrix4 {
	return Matrix4{
		m.M00*o.M00 + m.M01*o.M10 + m.M02*o.M20 + m.M03*o.M30,
		m.M10*o.M00 + m.M11*o.M10 + m.M12*o.M20 + m.M13*o.M30,
		m.M20*o.M00 + m.M21*o.M10 + m.M22*o.M20 + m.M23*o.M30,
		m.M30*o.M00 + m.M31*o.M10 + m.M32*o.M20 + m.M33*o.M30,

		m.M00*o.M01 + m.M01*o.M11 + m.M02*o.M21 + m.M03*o.M31,
		m.M10*o.M01 + m.M11*o.M11 + m.M12*o.M21 + m.M13*o.M31,
		m.M20*o.M01 + m.M21*o.M11 + m.M22*o.M21 + m.M23*o.M31,
		m.M30*o.M01 + m.M31*o.M11 + m.M32*o.M21 + m.M33*o.M31,

		m.M00*o.M02 + m.M01*o.M12 + m.M02*o.M22 + m.M03*o.M32,
		m.M10*o.M02 + m.M11*o.M12 + m.M12*o.M22 + m.M13*o.M32,
		m.M20*o.M02 + m.M21*o.M12 + m.M22*o.M22 + m.M23*o.M32,
		m.M30*o.M02 + m.M31*o.M12 + m.M32*o.M22 + m.M33*o.M32,

		m.M00*o.M03 + m.M01*o.M13 + m.M02*o.M23 + m.M03*o.M33,
		m.M10*o.M03 + m.M11*o.M13 + m.M12*o.M23 + m.M13*o.M33,
		m.M20*o.M03 + m.M21*o.M13 + m.M22*o.M23 + m.M23*o.M33,
		m.M30*o.M03 + m.M31*o.M13 + m.M32*o.M23 + m.M33*o.M33,
	}
}

func (m Matrix4) MulTuple(o Tuple4) Tuple4 {
	return Tuple4{
		m.M00*o.X + m.M01*o.Y + m.M02*o.Z + m.M03*o.W,
		m.M10*o.X + m.M11*o.Y + m.M12*o.Z + m.M13*o.W,
		m.M20*o.X + m.M21*o.Y + m.M22*o.Z + m.M23*o.W,
		m.M30*o.X + m.M31*o.Y + m.M32*o.Z + m.M33*o.W,
	}
}

func IdentityMatrix() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (m Matrix4) TransponseMatrix() Matrix4 {
	return Matrix4{
		m.M00, m.M01, m.M02, m.M03,
		m.M10, m.M11, m.M12, m.M13,
		m.M20, m.M21, m.M22, m.M23,
		m.M30, m.M31, m.M32, m.M33,
	}
}

// matrix inversion
// pre-calc 36 2x2 minors
// ??? why 36 (source - deepseek) - probably 18 is enough?
// represents as map[string]float32
// where key - upper-left and top-down indices from column-based 4x4 matrix
// i.e. 0-5, 0-9, 10-15, ...
func matrixMinors22(m Matrix4) map[string]float32 {
	minors := make(map[string]float32)
	minors["10-15"] = m.M22*m.M33 - m.M23*m.M32
	minors["9-15"] = m.M12*m.M33 - m.M13*m.M32
	minors["9-14"] = m.M12*m.M23 - m.M13*m.M22
	minors["8-15"] = m.M02*m.M33 - m.M03*m.M32
	minors["8-14"] = m.M02*m.M23 - m.M03*m.M22
	minors["8-13"] = m.M02*m.M13 - m.M03*m.M12

	minors["6-15"] = m.M21*m.M33 - m.M23*m.M31
	minors["5-15"] = m.M11*m.M33 - m.M13*m.M31
	minors["5-14"] = m.M11*m.M23 - m.M13*m.M21
	minors["4-15"] = m.M01*m.M33 - m.M03*m.M31
	minors["4-14"] = m.M01*m.M23 - m.M03*m.M21
	minors["4-13"] = m.M01*m.M13 - m.M03*m.M11

	minors["4-9"] = m.M01*m.M12 - m.M02*m.M11
	minors["4-10"] = m.M01*m.M22 - m.M02*m.M21
	minors["5-10"] = m.M11*m.M22 - m.M12*m.M21
	minors["4-11"] = m.M01*m.M32 - m.M02*m.M31
	minors["5-11"] = m.M11*m.M32 - m.M12*m.M31
	minors["6-11"] = m.M21*m.M32 - m.M22*m.M31

	return minors
}

// return column-based minors
// 0 4 8 12
// 1 5 9 13
// 2 6 10 14
// 3 7 11 15
func matrixMinors33(m Matrix4, minors22 map[string]float32) []float32 {
	minors := make([]float32, 16)
	minors[0] = m.M11*minors22["10-15"] - m.M21*minors22["9-15"] + m.M31*minors22["9-14"]
	minors[1] = m.M01*minors22["10-15"] - m.M21*minors22["8-15"] + m.M31*minors22["8-14"]
	minors[2] = m.M01*minors22["9-15"] - m.M11*minors22["8-15"] + m.M31*minors22["8-13"]
	minors[3] = m.M01*minors22["9-14"] - m.M11*minors22["8-14"] + m.M21*minors22["8-13"]
	minors[4] = m.M10*minors22["10-15"] - m.M20*minors22["9-15"] + m.M30*minors22["9-14"]
	minors[5] = m.M00*minors22["10-15"] - m.M20*minors22["8-15"] + m.M30*minors22["8-14"]
	minors[6] = m.M00*minors22["9-15"] - m.M10*minors22["8-15"] + m.M30*minors22["8-13"]
	minors[7] = m.M00*minors22["9-14"] - m.M10*minors22["8-14"] + m.M20*minors22["8-13"]
	minors[8] = m.M10*minors22["6-15"] - m.M20*minors22["5-15"] + m.M30*minors22["5-14"]
	minors[9] = m.M00*minors22["6-15"] - m.M20*minors22["4-15"] + m.M30*minors22["4-14"]
	minors[10] = m.M00*minors22["5-15"] - m.M10*minors22["4-15"] + m.M30*minors22["4-13"]
	minors[11] = m.M00*minors22["5-14"] - m.M10*minors22["4-14"] + m.M20*minors22["4-13"]
	minors[12] = m.M10*minors22["6-11"] - m.M20*minors22["5-11"] + m.M30*minors22["5-10"]
	minors[13] = m.M00*minors22["6-11"] - m.M20*minors22["4-11"] + m.M30*minors22["4-9"]
	minors[14] = m.M00*minors22["5-11"] - m.M10*minors22["4-11"] + m.M30*minors22["4-9"]
	minors[15] = m.M00*minors22["5-10"] - m.M10*minors22["4-10"] + m.M20*minors22["4-9"]

	return minors
}

func det(m Matrix4, minors33 []float32) float32 {
	return m.M00*minors33[0] - m.M10*minors33[1] + m.M20*minors33[2] - m.M30*minors33[3]
}

func (m Matrix4) InverseMatrix() (Matrix4, error) {
	minors22 := matrixMinors22(m)
	minors33 := matrixMinors33(m, minors22)
	det := det(m, minors33)
	if EqualWithEps(det, 0) {
		return Matrix4{}, fmt.Errorf("Cannot inverse matrix %s, determinant 0\n", m)
	}

	// inverse cofactor matrix * 1/det
	// 0 1 2 3
	// 4 5 6 7
	// 8 9 10 11
	// 12 13 14 15
	return Matrix4{
		minors33[0] / det, -minors33[4] / det, minors33[8] / det, -minors33[12] / det,
		-minors33[1] / det, minors33[5] / det, -minors33[9] / det, minors33[13] / det,
		minors33[2] / det, -minors33[6] / det, minors33[10] / det, -minors33[14] / det,
		-minors33[3] / det, minors33[7] / det, -minors33[11] / det, minors33[15] / det,
	}, nil
}
