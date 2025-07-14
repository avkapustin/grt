package primitives

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// matrix - column vs row  ordered in terms of performance
// raylib - "column major" - what does it mean?
// M<row><col> naming
// 16 elems x 4 byte = 64 byte - still valid to pass as value?
type Matrix4 struct {
	M00, M01, M02, M03 float32
	M10, M11, M12, M13 float32
	M20, M21, M22, M23 float32
	M30, M31, M32, M33 float32
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
			sb.WriteString(fmt.Sprintf("| %.4f ", matrix.FieldByName(fieldName).Float()))
		}
		sb.WriteString("|\n")
	}

	return sb.String()
}

func (m Matrix4) Mul(o Matrix4) Matrix4 {
	return Matrix4{
		m.M00*o.M00 + m.M01*o.M10 + m.M02*o.M20 + m.M03*o.M30,
		m.M00*o.M01 + m.M01*o.M11 + m.M02*o.M21 + m.M03*o.M31,
		m.M00*o.M02 + m.M01*o.M12 + m.M02*o.M22 + m.M03*o.M32,
		m.M00*o.M03 + m.M01*o.M13 + m.M02*o.M23 + m.M03*o.M33,

		m.M10*o.M00 + m.M11*o.M10 + m.M12*o.M20 + m.M13*o.M30,
		m.M10*o.M01 + m.M11*o.M11 + m.M12*o.M21 + m.M13*o.M31,
		m.M10*o.M02 + m.M11*o.M12 + m.M12*o.M22 + m.M13*o.M32,
		m.M10*o.M03 + m.M11*o.M13 + m.M12*o.M23 + m.M13*o.M33,

		m.M20*o.M00 + m.M21*o.M10 + m.M22*o.M20 + m.M23*o.M30,
		m.M20*o.M01 + m.M21*o.M11 + m.M22*o.M21 + m.M23*o.M31,
		m.M20*o.M02 + m.M21*o.M12 + m.M22*o.M22 + m.M23*o.M32,
		m.M20*o.M03 + m.M21*o.M13 + m.M22*o.M23 + m.M23*o.M33,

		m.M30*o.M00 + m.M31*o.M10 + m.M32*o.M20 + m.M33*o.M30,
		m.M30*o.M01 + m.M31*o.M11 + m.M32*o.M21 + m.M33*o.M31,
		m.M30*o.M02 + m.M31*o.M12 + m.M32*o.M22 + m.M33*o.M32,
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
