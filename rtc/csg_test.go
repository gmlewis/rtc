package rtc

import (
	"testing"
)

func TestCSG(t *testing.T) {
	s1 := Sphere()
	s2 := Cube()
	c := CSG(CSGUnion, s1, s2)

	if got, want := c.Operation, CSGUnion; got != want {
		t.Errorf("c.Operation = %v, want %v", got, want)
	}

	if got, want := c.Left, s1; got != want {
		t.Errorf("c.Left = %v, want %v", got, want)
	}

	if got, want := c.Right, s2; got != want {
		t.Errorf("c.Right = %v, want %v", got, want)
	}

	if got, want := s1.Parent(), c; got != want {
		t.Errorf("s1.Parent() = %v, want %v", got, want)
	}
}
