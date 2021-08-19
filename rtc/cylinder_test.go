package rtc

import (
	"fmt"
	"math"
	"testing"
)

func TestCylinder(t *testing.T) {
	c := Cylinder()

	if got, want := c.Minimum, math.Inf(-1); got != want {
		t.Errorf("Cylinder.Minimum = %v, want %v", got, want)
	}

	if got, want := c.Maximum, math.Inf(1); got != want {
		t.Errorf("Cylinder.Maximum = %v, want %v", got, want)
	}

	if got, want := c.Closed, false; got != want {
		t.Errorf("Cylinder.Closed = %v, want %v", got, want)
	}
}

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
			name:        "+x",
			objectPoint: Point(1, 0, 0),
			want:        Vector(1, 0, 0),
		},
		{
			name:        "-x",
			objectPoint: Point(0, 5, -1),
			want:        Vector(0, 0, -1),
		},
		{
			name:        "+z",
			objectPoint: Point(0, -2, 1),
			want:        Vector(0, 0, 1),
		},
		{
			name:        "-z",
			objectPoint: Point(-1, 1, 0),
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

func TestCylinder_IntersectContrainedCylinder(t *testing.T) {
	c := Cylinder()
	c.Minimum = 1
	c.Maximum = 2

	tests := []struct {
		point Tuple
		dir   Tuple
		want  int
	}{
		{point: Point(0, 1.5, 0), dir: Vector(0.1, 1, 1), want: 0},
		{point: Point(0, 3, -5), dir: Vector(0, 0, 1), want: 0},
		{point: Point(0, 0, -5), dir: Vector(0, 0, 1), want: 0},
		{point: Point(0, 2, -5), dir: Vector(0, 0, 1), want: 0},
		{point: Point(0, 1, -5), dir: Vector(0, 0, 1), want: 0},
		{point: Point(0, 1.5, -2), dir: Vector(0, 0, 1), want: 2},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i+1), func(t *testing.T) {
			dir := tt.dir.Normalize()
			r := Ray(tt.point, dir)
			xs := c.LocalIntersect(r)

			if got := len(xs); got != tt.want {
				t.Errorf("Cylinder.intersections = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCylinder_IntersectCapsOfClosedCylinder(t *testing.T) {
	c := Cylinder()
	c.Minimum = 1
	c.Maximum = 2
	c.Closed = true

	tests := []struct {
		point Tuple
		dir   Tuple
		want  int
	}{
		{point: Point(0, 3, 0), dir: Vector(0, -1, 0), want: 2},
		{point: Point(0, 3, -2), dir: Vector(0, -1, 2), want: 2},
		{point: Point(0, 4, -2), dir: Vector(0, -1, 1), want: 2},
		{point: Point(0, 0, -2), dir: Vector(0, 1, 2), want: 2},
		{point: Point(0, -1, -2), dir: Vector(0, 1, 1), want: 2},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i+1), func(t *testing.T) {
			dir := tt.dir.Normalize()
			r := Ray(tt.point, dir)
			xs := c.LocalIntersect(r)

			if got := len(xs); got != tt.want {
				t.Errorf("Cylinder.intersections = %v, want %v", got, tt.want)
			}
		})
	}
}
