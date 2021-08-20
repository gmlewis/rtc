package rtc

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)
	tri := Triangle(p1, p2, p3)

	if got, want := tri.P1, p1; !got.Equal(want) {
		t.Errorf("tri.P1 = %v, want %v", got, want)
	}

	if got, want := tri.P2, p2; !got.Equal(want) {
		t.Errorf("tri.P2 = %v, want %v", got, want)
	}

	if got, want := tri.P3, p3; !got.Equal(want) {
		t.Errorf("tri.P3 = %v, want %v", got, want)
	}

	if got, want := tri.E1, Vector(-1, -1, 0); !got.Equal(want) {
		t.Errorf("tri.E1 = %v, want %v", got, want)
	}

	if got, want := tri.E2, Vector(1, -1, 0); !got.Equal(want) {
		t.Errorf("tri.E2 = %v, want %v", got, want)
	}

	if got, want := tri.Normal, Vector(0, 0, -1); !got.Equal(want) {
		t.Errorf("tri.Normal = %v, want %v", got, want)
	}
}

func TestTriangleT_LocalNormalAt(t *testing.T) {
	tri := Triangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	n1 := tri.LocalNormalAt(Point(0, 0.5, 0))
	n2 := tri.LocalNormalAt(Point(-0.5, 0.75, 0))
	n3 := tri.LocalNormalAt(Point(0.5, 0.25, 0))

	if got, want := n1, tri.Normal; got != want {
		t.Errorf("n1 = %v, want %v", got, want)
	}
	if got, want := n2, tri.Normal; got != want {
		t.Errorf("n2 = %v, want %v", got, want)
	}
	if got, want := n3, tri.Normal; got != want {
		t.Errorf("n3 = %v, want %v", got, want)
	}
}
