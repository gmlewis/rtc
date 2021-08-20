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

	if got, want := s1.Parent(), c; got != want {
		t.Errorf("s1.Parent() = %v, want %v", got, want)
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
