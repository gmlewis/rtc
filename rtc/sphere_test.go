package rtc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSphere_Intersect(t *testing.T) {
	tests := []struct {
		name string
		ray  RayT
		want []float64
	}{
		{
			name: "A ray intersects a sphere at two points",
			ray:  Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			want: []float64{4.0, 6.0},
		},
		{
			name: "A ray intersects a sphere at a tangent",
			ray:  Ray(Point(0, 1, -5), Vector(0, 0, 1)),
			want: []float64{5.0, 5.0},
		},
		{
			name: "A ray intersects misses a sphere",
			ray:  Ray(Point(0, 2, -5), Vector(0, 0, 1)),
			want: nil,
		},
		{
			name: "A ray intersects originates inside a sphere",
			ray:  Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			want: []float64{-1.0, 1.0},
		},
		{
			name: "A sphere is behind a ray",
			ray:  Ray(Point(0, 0, 5), Vector(0, 0, 1)),
			want: []float64{-6.0, -4.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sphere{}
			if got := s.Intersect(tt.ray); !cmp.Equal(got, tt.want) {
				t.Errorf("Sphere.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
