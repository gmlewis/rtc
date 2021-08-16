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
