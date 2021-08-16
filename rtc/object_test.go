package rtc

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	s := Sphere{}
	i := Intersection(3.5, s)

	if got, want := i.T, 3.5; got != want {
		t.Errorf("i.T = %v, want %v", got, want)
	}

	if got, want := i.Object, s; got != want {
		t.Errorf("i.Object = %v, want %v", got, want)
	}
}

func TestIntersections(t *testing.T) {
	s := Sphere{}
	i1 := Intersection(1, s)
	i2 := Intersection(2, s)

	xs := Intersections(i1, i2)

	if len(xs) != 2 {
		t.Fatalf("len(xs) = %v, want 2", len(xs))
	}

	if got, want := xs[0].T, 1.0; got != want {
		t.Errorf("xs[0].T = %v, want %v", got, want)
	}

	if got, want := xs[1].T, 2.0; got != want {
		t.Errorf("xs[1].T = %v, want %v", got, want)
	}
}
