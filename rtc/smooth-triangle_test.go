package rtc

import (
	"math"
	"testing"
)

func testTri(t *testing.T) *SmoothTriangleT {
	t.Helper()
	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)
	n1 := Vector(0, 1, 0)
	n2 := Vector(-1, 0, 0)
	n3 := Vector(1, 0, 0)
	tri := SmoothTriangle(p1, p2, p3, n1, n2, n3)

	if got, want := tri.P1, p1; !got.Equal(want) {
		t.Errorf("tri.P1 = %v, want %v", got, want)
	}
	if got, want := tri.P2, p2; !got.Equal(want) {
		t.Errorf("tri.P2 = %v, want %v", got, want)
	}
	if got, want := tri.P3, p3; !got.Equal(want) {
		t.Errorf("tri.P3 = %v, want %v", got, want)
	}
	if got, want := tri.N1, n1; !got.Equal(want) {
		t.Errorf("tri.N1 = %v, want %v", got, want)
	}
	if got, want := tri.N2, n2; !got.Equal(want) {
		t.Errorf("tri.N2 = %v, want %v", got, want)
	}
	if got, want := tri.N3, n3; !got.Equal(want) {
		t.Errorf("tri.N3 = %v, want %v", got, want)
	}
	return tri
}

func TestSmoothTriangle_LocalIntersect_WithUV(t *testing.T) {
	tri := testTri(t)
	r := Ray(Point(-0.2, 0.3, -2), Vector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	if got, want := len(xs), 1; got != want {
		t.Fatalf("len(xs) = %v, want %v", got, want)
	}

	if got, want := xs[0].U, 0.45; math.Abs(got-want) > epsilon {
		t.Errorf("xs[0].U = %v, want %v", got, want)
	}

	if got, want := xs[0].V, 0.25; math.Abs(got-want) > epsilon {
		t.Errorf("xs[0].V = %v, want %v", got, want)
	}
}

func TestSmoothTriangle_NormalAt(t *testing.T) {
	tri := testTri(t)
	xs := IntersectionWithUV(3.5, tri, 0.45, 0.25)

	n := xs.NormalAt(Point(0, 0, 0))

	if got, want := n, Vector(-0.5547, 0.83205, 0); !got.Equal(want) {
		t.Errorf("n = %v, want %v", got, want)
	}
}
