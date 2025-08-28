package primitives_test

import (
	"log"
	"testing"

	p "github.com/avkapustin/grt/internal/primitives"
	"github.com/stretchr/testify/assert"
)

func checkTuplesEquality(t *testing.T, expected p.Tuple4, actual p.Tuple4) {
	assert.InDeltaf(t, expected.X, actual.X, 0.00001,
		"expected tuple %#v, actual tuple %#v, field X, expected %.6f, actual %.6f", expected, actual, expected.X, actual.X)
	assert.InDeltaf(t, expected.Y, actual.Y, 0.00001,
		"expected tuple %#v, actual tuple %#v, field Y, expected %.6f, actual %.6f", expected, actual, expected.Y, actual.Y)
	assert.InDeltaf(t, expected.Z, actual.Z, 0.00001,
		"expected tuple %#v, actual tuple %#v, field Z, expected %.6f, actual %.6f", expected, actual, expected.Z, actual.Z)
	assert.InDeltaf(t, expected.W, actual.W, 0.00001,
		"expected tuple %#v, actual tuple %#v, field W, expected %.6f, actual %.6f", expected, actual, expected.W, actual.W)
}

func TestTuplesMath(t *testing.T) {
	t1 := p.MakePoint(1, 1, 1)
	t2 := p.MakeVector(2, 2, 2)

	actual := t1.Add(t2)
	expected := p.MakePoint(3, 3, 3)

	checkTuplesEquality(t, expected, actual)

	m := t2.Magnitude()
	scaled := t2.Norm().Scale(m)

	checkTuplesEquality(t, t2, scaled)
}

// for small structures copy is preferred vs pointer
// stack vs heap + direct vs indirect access to fields
// apple m4 max 0.5 ns/op vs 7 ns/op
func BenchmarkCopy(t *testing.B) {
	t1 := p.MakePoint(1, 1, 1)
	t2 := p.MakeVector(2, 2, 2)

	for i := 0; i < t.N; i++ {
		t1 = t1.Add(t2)
	}
	log.Printf("Final value %#v\n", t1)
}

func tAdd(a *p.Tuple4, b *p.Tuple4) *p.Tuple4 {
	return &p.Tuple4{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
		a.W + b.W,
	}
}

func BenchmarkPointer(t *testing.B) {
	t1 := p.MakePoint(1, 1, 1)
	t2 := p.MakeVector(2, 2, 2)
	var tr *p.Tuple4
	tr = &t1

	for i := 0; i < t.N; i++ {
		tr = tAdd(tr, &t2)
	}

	log.Printf("Final value %#v\n", tr)
}
