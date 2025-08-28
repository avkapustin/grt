package primitives

import "math"

type Tuple4 struct {
	X, Y, Z, W float32
}

type Ray struct {
	Origin    Tuple4
	Direction Tuple4
}

func MakeVector(x, y, z float32) Tuple4 {
	return Tuple4{
		x,
		y,
		z,
		0.0,
	}
}

func MakePoint(x, y, z float32) Tuple4 {
	return Tuple4{
		x,
		y,
		z,
		1.0,
	}
}

func (t Tuple4) Add(o Tuple4) Tuple4 {
	return Tuple4{
		t.X + o.X,
		t.Y + o.Y,
		t.Z + o.Z,
		t.W + o.W,
	}
}

func (t Tuple4) Sub(o Tuple4) Tuple4 {
	return Tuple4{
		t.X - o.X,
		t.Y - o.Y,
		t.Z - o.Z,
		t.W - o.W,
	}
}

func (t Tuple4) Neg() Tuple4 {
	return Tuple4{
		-t.X,
		-t.Y,
		-t.Z,
		t.W,
	}
}

func (t Tuple4) Scale(s float32) Tuple4 {
	return Tuple4{
		s * t.X,
		s * t.Y,
		s * t.Z,
		t.W,
	}
}

func (t Tuple4) Magnitude() float32 {
	return float32(math.Sqrt(float64(t.X*t.X + t.Y*t.Y + t.Z*t.Z)))
}

func (t Tuple4) Norm() Tuple4 {
	m := t.Magnitude()

	return Tuple4{
		t.X / m,
		t.Y / m,
		t.Z / m,
		t.W,
	}
}

func (t Tuple4) Dot(o Tuple4) float32 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z
}

func (t Tuple4) Cross(o Tuple4) Tuple4 {
	return Tuple4{
		t.Y*o.Z - t.Z*o.Y,
		t.Z*o.X - t.X*o.Z,
		t.X*o.Y - t.Y*o.X,
		0.0,
	}
}

func (r Ray) Position(distance float32) Tuple4 {
	return r.Origin.Add(r.Direction.Scale(distance))
}
