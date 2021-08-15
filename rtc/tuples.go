package rtc

import (
	"math"

	"github.com/gmlewis/go3d/float64/vec4"
)

const (
	epsilon = 1e-4
)

// Tuple is a vec4.
type Tuple vec4.T

// X returns the X value of the Tuple.
func (t *Tuple) X() float64 {
	if t == nil {
		return 0
	}
	return t[0]
}

// Y returns the Y value of the Tuple.
func (t *Tuple) Y() float64 {
	if t == nil {
		return 0
	}
	return t[1]
}

// Z returns the Z value of the Tuple.
func (t *Tuple) Z() float64 {
	if t == nil {
		return 0
	}
	return t[2]
}

// W returns the W value of the Tuple.
func (t *Tuple) W() float64 {
	if t == nil {
		return 0
	}
	return t[3]
}

// IsPoint identifies the Tuple as a Point.
func (t *Tuple) IsPoint() bool {
	return t != nil && t[3] == 1.0
}

// IsVector identifies the Tuple as a Vector.
func (t *Tuple) IsVector() bool {
	return t == nil || t[3] == 0.0
}

// Point returns a new Tuple as a Point.
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

// Vector returns a new Tuple as a Vector.
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

// Equal tests if two Tuples are equal.
func (t *Tuple) Equal(other *Tuple) bool {
	return math.Abs(t.X()-other.X()) < epsilon &&
		math.Abs(t.Y()-other.Y()) < epsilon &&
		math.Abs(t.Z()-other.Z()) < epsilon &&
		math.Abs(t.W()-other.W()) < epsilon
}

// Add adds two Tuples and returns a new one.
func (t *Tuple) Add(other *Tuple) Tuple {
	return Tuple{
		t.X() + other.X(),
		t.Y() + other.Y(),
		t.Z() + other.Z(),
		t.W() + other.W(),
	}
}

// Sub subtracts two Tuples and returns a new one.
func (t *Tuple) Sub(other *Tuple) Tuple {
	return Tuple{
		t.X() - other.X(),
		t.Y() - other.Y(),
		t.Z() - other.Z(),
		t.W() - other.W(),
	}
}

// Negate negates a Tuple.
func (t *Tuple) Negate() Tuple {
	return Tuple{
		-t.X(),
		-t.Y(),
		-t.Z(),
		-t.W(),
	}
}

// MulScalar multiplies a tuple by a scalar.
func (t *Tuple) MulScalar(f float64) Tuple {
	return Tuple{
		f * t.X(),
		f * t.Y(),
		f * t.Z(),
		f * t.W(),
	}
}

// DivScalar divides a tuple by a scalar.
func (t *Tuple) DivScalar(f float64) Tuple {
	return t.MulScalar(1 / f)
}

// Magnitude computes the magnitude or length of a vector (Tuple).
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(
		t.X()*t.X() +
			t.Y()*t.Y() +
			t.Z()*t.Z() +
			t.W()*t.W())
}

// Normalize normalizes a vector to a unit vector (of length 1).
func (t *Tuple) Normalize() Tuple {
	return t.DivScalar(t.Magnitude())
}

// Dot computes the dot product (aka "scalar product" or "inner product")
// of two vectors (Tuples). The dot product is the cosine of the angle
// between two unit vectors.
func (t *Tuple) Dot(other *Tuple) float64 {
	return t.X()*other.X() +
		t.Y()*other.Y() +
		t.Z()*other.Z() +
		t.W()*other.W()
}
