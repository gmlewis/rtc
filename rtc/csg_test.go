package rtc

import (
	"fmt"
	"testing"
)

func TestCSG(t *testing.T) {
	s1 := Sphere()
	s2 := Cube()
	c := CSG(CSGUnion, s1, s2)

	if got, want := c.Operation, CSGUnion; got != want {
		t.Errorf("c.Operation = %v, want %v", got, want)
	}

	if got, want := c.Left, s1; got != want {
		t.Errorf("c.Left = %v, want %v", got, want)
	}

	if got, want := c.Right, s2; got != want {
		t.Errorf("c.Right = %v, want %v", got, want)
	}

	if got, want := s1.GetParent(), c; got != want {
		t.Errorf("s1.GetParent() = %v, want %v", got, want)
	}
}

func Test_intersectionAllowed(t *testing.T) {
	tests := []struct {
		op      CSGOperation
		leftHit bool
		inLeft  bool
		inRight bool
		want    bool
	}{
		{op: CSGUnion, leftHit: true, inLeft: true, inRight: true, want: false},
		{op: CSGUnion, leftHit: true, inLeft: true, inRight: false, want: true},
		{op: CSGUnion, leftHit: true, inLeft: false, inRight: true, want: false},
		{op: CSGUnion, leftHit: true, inLeft: false, inRight: false, want: true},
		{op: CSGUnion, leftHit: false, inLeft: true, inRight: true, want: false},
		{op: CSGUnion, leftHit: false, inLeft: true, inRight: false, want: false},
		{op: CSGUnion, leftHit: false, inLeft: false, inRight: true, want: true},
		{op: CSGUnion, leftHit: false, inLeft: false, inRight: false, want: true},
		{op: CSGIntersection, leftHit: true, inLeft: true, inRight: true, want: true},
		{op: CSGIntersection, leftHit: true, inLeft: true, inRight: false, want: false},
		{op: CSGIntersection, leftHit: true, inLeft: false, inRight: true, want: true},
		{op: CSGIntersection, leftHit: true, inLeft: false, inRight: false, want: false},
		{op: CSGIntersection, leftHit: false, inLeft: true, inRight: true, want: true},
		{op: CSGIntersection, leftHit: false, inLeft: true, inRight: false, want: true},
		{op: CSGIntersection, leftHit: false, inLeft: false, inRight: true, want: false},
		{op: CSGIntersection, leftHit: false, inLeft: false, inRight: false, want: false},
		{op: CSGDifference, leftHit: true, inLeft: true, inRight: true, want: false},
		{op: CSGDifference, leftHit: true, inLeft: true, inRight: false, want: true},
		{op: CSGDifference, leftHit: true, inLeft: false, inRight: true, want: false},
		{op: CSGDifference, leftHit: true, inLeft: false, inRight: false, want: true},
		{op: CSGDifference, leftHit: false, inLeft: true, inRight: true, want: true},
		{op: CSGDifference, leftHit: false, inLeft: true, inRight: false, want: true},
		{op: CSGDifference, leftHit: false, inLeft: false, inRight: true, want: false},
		{op: CSGDifference, leftHit: false, inLeft: false, inRight: false, want: false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i+1), func(t *testing.T) {
			if got := intersectionAllowed(tt.op, tt.leftHit, tt.inLeft, tt.inRight); got != tt.want {
				t.Errorf("intersectionAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSG_FilterIntersections(t *testing.T) {
	s1 := Sphere()
	s2 := Cube()
	xs := Intersections(Intersection(1, s1), Intersection(2, s2), Intersection(3, s1), Intersection(4, s2))

	tests := []struct {
		name string
		op   CSGOperation
		x0   int
		x1   int
	}{
		{
			name: "union",
			op:   CSGUnion,
			x0:   0,
			x1:   3,
		},
		{
			name: "intersection",
			op:   CSGIntersection,
			x0:   1,
			x1:   2,
		},
		{
			name: "difference",
			op:   CSGDifference,
			x0:   0,
			x1:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CSG(tt.op, s1, s2)

			result := c.FilterIntersections(xs)

			if got, want := len(result), 2; got != want {
				t.Fatalf("len(result) = %v, want %v", got, want)
			}

			if got, want := result[0].T, xs[tt.x0].T; got != want {
				t.Errorf("result[0].T = %v, want %v", got, want)
			}

			if got, want := result[1].T, xs[tt.x1].T; got != want {
				t.Errorf("result[1].T = %v, want %v", got, want)
			}
		})
	}
}

func TestCSG_RayMisses(t *testing.T) {
	c := CSG(CSGUnion, Sphere(), Cube())
	r := Ray(Point(0, 2, -5), Vector(0, 0, 1))
	xs := c.LocalIntersect(r)

	if got, want := len(xs), 0; got != want {
		t.Errorf("len(xs) = %v, want %v", got, want)
	}
}

func TestCSG_RayHits(t *testing.T) {
	s1 := Sphere()
	s2 := Sphere().SetTransform(Translation(0, 0, 0.5))
	c := CSG(CSGUnion, s1, s2)
	r := Ray(Point(0, 0, -5), Vector(0, 0, 1))
	xs := c.LocalIntersect(r)

	if got, want := len(xs), 2; got != want {
		t.Fatalf("len(xs) = %v, want %v", got, want)
	}

	if got, want := xs[0].T, 4.0; got != want {
		t.Errorf("xs[0].T = %v, want %v", got, want)
	}

	if got, want := xs[0].Object, s1; got != want {
		t.Errorf("xs[0].Object = %v, want %v", got, want)
	}
	if got, want := xs[1].T, 6.5; got != want {
		t.Errorf("xs[1].T = %v, want %v", got, want)
	}

	if got, want := xs[1].Object, s2; got != want {
		t.Errorf("xs[1].Object = %v, want %v", got, want)
	}
}
