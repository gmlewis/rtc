package rtc

import (
	"math"
	"testing"
)

func TestCylinderT_LocalIntersect(t *testing.T) {
	c := Cylinder()

	tests := []struct {
		name string
		ray  RayT
		want []IntersectionT
	}{
		{
			name: "miss 1",
			ray:  Ray(Point(1, 0, 0), Vector(0, 1, 0)),
			want: nil,
		},
		{
			name: "miss 2",
			ray:  Ray(Point(0, 0, 0), Vector(0, 1, 0)),
			want: nil,
		},
		{
			name: "miss 3",
			ray:  Ray(Point(0, 0, -5), Vector(1, 1, 1)),
			want: nil,
		},
		{
			name: "hit 1",
			ray:  Ray(Point(1, 0, -5), Vector(0, 0, 1)),
			want: Intersections(Intersection(5, c), Intersection(5, c)),
		},
		{
			name: "hit 2",
			ray:  Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			want: Intersections(Intersection(4, c), Intersection(6, c)),
		},
		{
			name: "hit 3",
			ray:  Ray(Point(0.5, 0, -5), Vector(0.1, 1, 1)),
			want: Intersections(Intersection(6.80798, c), Intersection(7.08872, c)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ray.Direction = tt.ray.Direction.Normalize()
			got := c.LocalIntersect(tt.ray)

			if len(got) != len(tt.want) {
				t.Fatalf("got = %v hits, want %v", len(got), len(tt.want))
			}

			if len(got) == 0 {
				return
			}

			if math.Abs(got[0].T-tt.want[0].T) > epsilon {
				t.Errorf("LocalIntersect[0] = %v, want %v", got, tt.want)
			}

			if math.Abs(got[1].T-tt.want[1].T) > epsilon {
				t.Errorf("LocalIntersect[1] = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCylinderT_LocalNormalAt(t *testing.T) {
	p := Cylinder()

	tests := []struct {
		name        string
		objectPoint Tuple
		want        Tuple
	}{
		{
			name:        "1, 0.5, -0.8",
			objectPoint: Point(1, 0.5, -0.8),
			want:        Vector(1, 0, 0),
		},
		{
			name:        "-1, -0.2, 0.9",
			objectPoint: Point(-1, -0.2, 0.9),
			want:        Vector(-1, 0, 0),
		},
		{
			name:        "-0.4, 1, -0.1",
			objectPoint: Point(-0.4, 1, -0.1),
			want:        Vector(0, 1, 0),
		},
		{
			name:        "0.3, -1, -0.7",
			objectPoint: Point(0.3, -1, -0.7),
			want:        Vector(0, -1, 0),
		},
		{
			name:        "-0.6, 0.3, 1",
			objectPoint: Point(-0.6, 0.3, 1),
			want:        Vector(0, 0, 1),
		},
		{
			name:        "0.4, 0.4, -1",
			objectPoint: Point(0.4, 0.4, -1),
			want:        Vector(0, 0, -1),
		},
		{
			name:        "1, 1, 1",
			objectPoint: Point(1, 1, 1),
			want:        Vector(1, 0, 0),
		},
		{
			name:        "-1, -1, -1",
			objectPoint: Point(-1, -1, -1),
			want:        Vector(-1, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.LocalNormalAt(tt.objectPoint); !got.Equal(tt.want) {
				t.Errorf("CylinderT.LocalNormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
