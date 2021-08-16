package rtc

import (
	"testing"
)

func TestSphere_Intersect(t *testing.T) {
	s := Sphere()

	tests := []struct {
		name string
		ray  RayT
		want []IntersectionT
	}{
		{
			name: "A ray intersects a sphere at two points",
			ray:  Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			want: []IntersectionT{Intersection(4.0, s), Intersection(6.0, s)},
		},
		{
			name: "A ray intersects a sphere at a tangent",
			ray:  Ray(Point(0, 1, -5), Vector(0, 0, 1)),
			want: []IntersectionT{Intersection(5.0, s), Intersection(5.0, s)},
		},
		{
			name: "A ray intersects misses a sphere",
			ray:  Ray(Point(0, 2, -5), Vector(0, 0, 1)),
			want: nil,
		},
		{
			name: "A ray intersects originates inside a sphere",
			ray:  Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			want: []IntersectionT{Intersection(-1.0, s), Intersection(1.0, s)},
		},
		{
			name: "A sphere is behind a ray",
			ray:  Ray(Point(0, 0, 5), Vector(0, 0, 1)),
			want: []IntersectionT{Intersection(-6.0, s), Intersection(-4.0, s)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Intersect(tt.ray)
			if len(got) != len(tt.want) {
				t.Fatalf("Sphere.Intersect() = %v, want %v", got, tt.want)
			}

			for i, w := range tt.want {
				if got[i].T != w.T || got[i].Object != w.Object {
					t.Errorf("Sphere.Intersect[%v] = %v, want %v", i, got[i], w)
				}
			}
		})
	}
}
