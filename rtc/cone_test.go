package rtc

import (
	"fmt"
	"math"
	"testing"
)

func TestCone(t *testing.T) {
	c := Cone()

	if got, want := c.Minimum, math.Inf(-1); got != want {
		t.Errorf("Cone.Minimum = %v, want %v", got, want)
	}

	if got, want := c.Maximum, math.Inf(1); got != want {
		t.Errorf("Cone.Maximum = %v, want %v", got, want)
	}

	if got, want := c.Closed, false; got != want {
		t.Errorf("Cone.Closed = %v, want %v", got, want)
	}
}

func TestConeT_LocalIntersect(t *testing.T) {
	c := Cone()

	tests := []struct {
		ray  RayT
		want []IntersectionT
	}{
		{
			ray:  Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			want: Intersections(Intersection(5, c), Intersection(5, c)),
		},
		{
			ray:  Ray(Point(0, 0, -5), Vector(1, 1, 1)),
			want: Intersections(Intersection(8.66025, c), Intersection(8.66025, c)),
		},
		{
			ray:  Ray(Point(1, 1, -5), Vector(-0.5, -1, 1)),
			want: Intersections(Intersection(4.55006, c), Intersection(49.44994, c)),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i+1), func(t *testing.T) {
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

func TestCone_IntersectWithParallelRayToSide(t *testing.T) {
	c := Cone()
	dir := Vector(0, 1, 1).Normalize()
	r := Ray(Point(0, 0, -1), dir)
	xs := c.LocalIntersect(r)

	if got, want := len(xs), 1; got != want {
		t.Fatalf("c.LocalIntersect = %v hits, want %v", got, want)
	}

	if got, want := xs[0].T, 0.35355; math.Abs(got-want) > epsilon {
		t.Errorf("c.LocalIntersect = %v, want %v", got, want)
	}
}

func TestCone_IntersectContrainedCone(t *testing.T) {
	c := Cone()
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
				t.Errorf("Cone.intersections = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCone_IntersectCapsOfClosedCone(t *testing.T) {
	c := Cone()
	c.Minimum = -0.5
	c.Maximum = 0.5
	c.Closed = true

	tests := []struct {
		point Tuple
		dir   Tuple
		want  int
	}{
		{point: Point(0, 0, -5), dir: Vector(0, 1, 0), want: 0},
		{point: Point(0, 0, -0.25), dir: Vector(0, 1, 1), want: 2},
		{point: Point(0, 0, -0.25), dir: Vector(0, 1, 0), want: 4},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i+1), func(t *testing.T) {
			dir := tt.dir.Normalize()
			r := Ray(tt.point, dir)
			xs := c.LocalIntersect(r)

			if got := len(xs); got != tt.want {
				t.Errorf("Cone.intersections = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConeT_LocalNormalAt(t *testing.T) {
	c := Cone()

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
			if got := c.LocalNormalAt(tt.objectPoint); !got.Equal(tt.want) {
				t.Errorf("ConeT.LocalNormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConeT_LocalNormalAtWithEndCaps(t *testing.T) {
	c := Cone()
	c.Minimum = 1
	c.Maximum = 2
	c.Closed = true

	tests := []struct {
		objectPoint Tuple
		want        Tuple
	}{
		{
			objectPoint: Point(0, 1, 0),
			want:        Vector(0, -1, 0),
		},
		{
			objectPoint: Point(0.5, 1, 0),
			want:        Vector(0, -1, 0),
		},
		{
			objectPoint: Point(0, 1, 0.5),
			want:        Vector(0, -1, 0),
		},
		{
			objectPoint: Point(0, 2, 0),
			want:        Vector(0, 1, 0),
		},
		{
			objectPoint: Point(0.5, 2, 0),
			want:        Vector(0, 1, 0),
		},
		{
			objectPoint: Point(0, 2, 0.5),
			want:        Vector(0, 1, 0),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if got := c.LocalNormalAt(tt.objectPoint); !got.Equal(tt.want) {
				t.Errorf("ConeT.LocalNormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
