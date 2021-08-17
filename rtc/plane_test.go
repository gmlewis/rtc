package rtc

import (
	"reflect"
	"testing"
)

func TestPlane(t *testing.T) {
	tests := []struct {
		name string
		want *PlaneT
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Plane(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plane() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaneT_LocalIntersect(t *testing.T) {
	p := Plane()

	tests := []struct {
		name string
		ray  RayT
		want []IntersectionT
	}{
		{
			name: "Intersect with a ray parallel to the plane",
			ray:  Ray(Point(0, 10, 0), Vector(0, 0, 1)),
			want: nil,
		},
		{
			name: "Intersect with a coplanar ray",
			ray:  Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			want: nil,
		},
		{
			name: "A ray intersecting a plane from above",
			ray:  Ray(Point(0, 1, 0), Vector(0, -1, 0)),
			want: Intersections(Intersection(1, p)),
		},
		{
			name: "A ray intersecting a plane from below",
			ray:  Ray(Point(0, -1, 0), Vector(0, 1, 0)),
			want: Intersections(Intersection(1, p)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.LocalIntersect(tt.ray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaneT.LocalIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaneT_LocalNormalAt(t *testing.T) {
	p := Plane()

	tests := []struct {
		name        string
		objectPoint Tuple
		want        Tuple
	}{
		{
			name:        "0,0,0",
			objectPoint: Point(0, 0, 0),
			want:        Vector(0, 1, 0),
		},
		{
			name:        "10,0,-10",
			objectPoint: Point(10, 0, -10),
			want:        Vector(0, 1, 0),
		},
		{
			name:        "-5,0,150",
			objectPoint: Point(-5, 0, 150),
			want:        Vector(0, 1, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.LocalNormalAt(tt.objectPoint); !got.Equal(tt.want) {
				t.Errorf("PlaneT.LocalNormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
