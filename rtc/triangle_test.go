package rtc

import (
	"math"
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

func TestTriangleT_LocalIntersect_ParallelRay(t *testing.T) {
	tri := Triangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := Ray(Point(0, -1, -2), Vector(0, 1, 0))
	xs := tri.LocalIntersect(r)

	if got, want := len(xs), 0; got != want {
		t.Errorf("len(xs) = %v, want %v", got, want)
	}
}

func TestTriangleT_LocalIntersect_RayMissOnE2(t *testing.T) {
	tri := Triangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := Ray(Point(1, 1, -2), Vector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	if got, want := len(xs), 0; got != want {
		t.Errorf("len(xs) = %v, want %v", got, want)
	}
}

func TestTriangleT_LocalIntersect_RayMissOnE1(t *testing.T) {
	tri := Triangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := Ray(Point(-1, 1, -2), Vector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	if got, want := len(xs), 0; got != want {
		t.Errorf("len(xs) = %v, want %v", got, want)
	}
}

func TestTriangleT_LocalIntersect_RayMissOnE3(t *testing.T) {
	tri := Triangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := Ray(Point(0, -1, -2), Vector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	if got, want := len(xs), 0; got != want {
		t.Errorf("len(xs) = %v, want %v", got, want)
	}
}

func TestTriangleT_LocalIntersect_RayHits(t *testing.T) {
	tri := Triangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := Ray(Point(0, 0.5, -2), Vector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	if got, want := len(xs), 1; got != want {
		t.Errorf("len(xs) = %v, want %v", got, want)
	}

	if got, want := xs[0].T, 2.0; math.Abs(got-want) > epsilon {
		t.Errorf("xs[0].T = %v, want %v", got, want)
	}
}

func testTri(t *testing.T) *TriangleT {
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

func TestSmoothTriangle(t *testing.T) {
	testTri(t)
}

func TestTriangle_WithUV(t *testing.T) {
	Triangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
}
