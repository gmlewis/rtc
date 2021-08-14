// Package rtc implements the code for the Ray Tracer Challenge book.
package rtc

import (
	"github.com/gmlewis/go3d/float64/vec4"
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
