package rtc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestRayT_Transform(t *testing.T) {
	tests := []struct {
		name string
		ray  RayT
		m    M4
		want RayT
	}{
		{
			name: "Translating a ray",
			ray:  Ray(Point(1, 2, 3), Vector(0, 1, 0)),
			m:    Translation(3, 4, 5),
			want: Ray(Point(4, 6, 8), Vector(0, 1, 0)),
		},
		{
			name: "Scaling a ray",
			ray:  Ray(Point(1, 2, 3), Vector(0, 1, 0)),
			m:    Scaling(2, 3, 4),
			want: Ray(Point(2, 6, 12), Vector(0, 3, 0)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ray.Transform(tt.m); !cmp.Equal(got, tt.want) {
				t.Errorf("RayT.GetTransform() = %v, want %v", got, tt.want)
			}

			if cmp.Equal(tt.ray, tt.want) {
				t.Errorf("RayT.Transform modified original ray but should not")
			}
		})
	}
}
