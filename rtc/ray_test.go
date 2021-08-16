package rtc

import (
	"testing"
)

func TestRay_Create_Query(t *testing.T) {
	origin := Point(1, 2, 3)
	direction := Vector(4, 5, 6)
	r := Ray(origin, direction)

	if r.Origin != origin {
		t.Errorf("r.Origin = %v, want %v", r.Origin, origin)
	}

	if r.Direction != direction {
		t.Errorf("r.Direction = %v, want %v", r.Direction, direction)
	}
}

func TestRayT_Position(t *testing.T) {
	r := Ray(Point(2, 3, 4), Vector(1, 0, 0))

	if got, want := r.Position(0), Point(2, 3, 4); got != want {
		t.Errorf("r.Position(0) = %v, want %v", got, want)
	}

	if got, want := r.Position(1), Point(3, 3, 4); got != want {
		t.Errorf("r.Position(1) = %v, want %v", got, want)
	}

	if got, want := r.Position(-1), Point(1, 3, 4); got != want {
		t.Errorf("r.Position(-1) = %v, want %v", got, want)
	}

	if got, want := r.Position(2.5), Point(4.5, 3, 4); got != want {
		t.Errorf("r.Position(2.5) = %v, want %v", got, want)
	}
}
