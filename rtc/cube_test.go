package rtc

import (
	"reflect"
	"testing"
)

func TestCubeT_LocalIntersect(t *testing.T) {
	c := Cube()

	tests := []struct {
		name string
		ray  RayT
		want []IntersectionT
	}{
		{
			name: "+x",
			ray:  Ray(Point(5, 0.5, 0), Vector(-1, 0, 0)),
			want: Intersections(Intersection(4, c), Intersection(6, c)),
		},
		{
			name: "-x",
			ray:  Ray(Point(-5, 0.5, 0), Vector(1, 0, 0)),
			want: Intersections(Intersection(4, c), Intersection(6, c)),
		},
		{
			name: "+y",
			ray:  Ray(Point(0.5, 5, 0), Vector(0, -1, 0)),
			want: Intersections(Intersection(4, c), Intersection(6, c)),
		},
		{
			name: "-y",
			ray:  Ray(Point(0.5, -5, 0), Vector(0, 1, 0)),
			want: Intersections(Intersection(4, c), Intersection(6, c)),
		},
		{
			name: "+z",
			ray:  Ray(Point(0.5, 0, 5), Vector(0, 0, -1)),
			want: Intersections(Intersection(4, c), Intersection(6, c)),
		},
		{
			name: "-z",
			ray:  Ray(Point(0.5, 0, -5), Vector(0, 0, 1)),
			want: Intersections(Intersection(4, c), Intersection(6, c)),
		},
		{
			name: "inside",
			ray:  Ray(Point(0, 0.5, 0), Vector(1, 0, 0)),
			want: Intersections(Intersection(-1, c), Intersection(1, c)),
		},
		{
			name: "miss 1",
			ray:  Ray(Point(-2, 0, 0), Vector(0.2673, 0.5345, 0.8018)),
			want: nil,
		},
		{
			name: "miss 2",
			ray:  Ray(Point(0, -2, 0), Vector(0.8018, 0.2673, 0.5345)),
			want: nil,
		},
		{
			name: "miss 3",
			ray:  Ray(Point(0, 0, -2), Vector(0.5345, 0.8018, 0.2673)),
			want: nil,
		},
		{
			name: "miss 4",
			ray:  Ray(Point(2, 0, 2), Vector(0, 0, -1)),
			want: nil,
		},
		{
			name: "miss 5",
			ray:  Ray(Point(0, 2, 2), Vector(0, -1, 0)),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.LocalIntersect(tt.ray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CubeT.LocalIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCubeT_LocalNormalAt(t *testing.T) {
	p := Cube()

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
				t.Errorf("CubeT.LocalNormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
