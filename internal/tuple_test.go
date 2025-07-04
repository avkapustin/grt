package tuple_test

import (
	"log"
	"math"
	"testing"

	"github.com/avkapustin/grt/internal"
)

// This is valid only for relative small float values
func equalWithEps(a, b float32) bool {
	eps := 0.0000001
	af := float64(a)
	bf := float64(b)

	if math.IsNaN(af) || math.IsNaN(bf) {
		return false
	}

	return math.Abs(af-bf) < eps
}

func isTuplesEqual(t1 tuple.Tuple4, t2 tuple.Tuple4) bool {
	return equalWithEps(t1.X, t2.X) || equalWithEps(t1.Y, t2.Y) || equalWithEps(t1.Z, t2.Z) || equalWithEps(t1.W, t2.W)
}

func TestTuplesMath(t *testing.T) {
	t1 := tuple.MakePoint(1, 1, 1)
	t2 := tuple.MakeVector(2, 2, 2)

	actual := t1.Add(t2)
	expected := tuple.MakePoint(3, 3, 3)

	if !isTuplesEqual(actual, expected) {
		t.Errorf("tuples add failed, given %#v, expected %#v\n", actual, expected)
	}

	m := t2.Magnitude()
	scaled := t2.Norm().Scale(m)

	if !isTuplesEqual(t2, scaled) {
		t.Errorf("tuples add failed, given %#v, expected %#v\n", t2, scaled)
	}
}

// for small structures copy is preferred vs pointer
// stack vs heap + direct vs indirect access to fields
// apple m4 max 0.5 ns/op vs 7 ns/op
func BenchmarkCopy(t *testing.B) {
	t1 := tuple.MakePoint(1, 1, 1)
	t2 := tuple.MakeVector(2, 2, 2)

	for i := 0; i < t.N ; i++ {
		t1 = t1.Add(t2)
	}
	log.Printf("Final value %#v\n", t1)
}

func tAdd(a *tuple.Tuple4, b *tuple.Tuple4) *tuple.Tuple4 {
	return &tuple.Tuple4{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
		a.W + b.W,
	}
}

func BenchmarkPointer(t *testing.B) {
	t1 := tuple.MakePoint(1, 1, 1)
	t2 := tuple.MakeVector(2, 2, 2)
	var tr *tuple.Tuple4
	tr = &t1

	for i := 0; i < t.N ; i++ {
		tr = tAdd(tr, &t2)
	}

	log.Printf("Final value %#v\n", tr)
}
