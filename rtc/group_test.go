package rtc

import (
	"testing"
)

func TestGroup(t *testing.T) {
	g := Group()

	if got, want := g.Transform(), M4Identity(); !got.Equal(want) {
		t.Errorf("Group.Transform = %v, want %v", got, want)
	}

	if got, want := len(g.Children), 0; got != want {
		t.Errorf("Group.Shapes len = %v, want %v", got, want)
	}
}

func TestGroupT_AddChild(t *testing.T) {
	g := Group()
	s := testShape()
	g.AddChild(s.shape)

	if got, want := len(g.Children), 1; got != want {
		t.Fatalf("g.Children len = %v, want %v", got, want)
	}

	if got, want := g.Children[0], Object(s.shape); got != want {
		t.Errorf("g.Children[0] = %v, want %v", got, want)
	}

	if got, want := s.shape.Parent(), g; got != want {
		t.Errorf("s.shape.Parent = %v, want %v", got, want)
	}
}

func TestGroupT_LocalIntersect_EmptyGroup(t *testing.T) {
	g := Group()
	r := Ray(Point(0, 0, 0), Vector(0, 0, 1))
	xs := g.LocalIntersect(r)

	if got, want := len(xs), 0; got != want {
		t.Errorf("len(xs) = %v, want %v", got, want)
	}
}

func TestGroupT_LocalIntersect_NonEmptyGroup(t *testing.T) {
	s1 := Sphere()
	s2 := Sphere().SetTransform(Translation(0, 0, -3))
	s3 := Sphere().SetTransform(Translation(5, 0, 0))
	g := Group(s1, s2, s3)
	r := Ray(Point(0, 0, -5), Vector(0, 0, 1))
	xs := g.LocalIntersect(r)

	if got, want := len(xs), 4; got != want {
		t.Fatalf("len(xs) = %v, want %v\nxs = %#v", got, want, xs)
	}

	if got, want := xs[0].Object, s2; got != want {
		t.Errorf("xs[0].Object = %v, want %v", got, want)
	}

	if got, want := xs[1].Object, s2; got != want {
		t.Errorf("xs[1].Object = %v, want %v", got, want)
	}

	if got, want := xs[2].Object, s1; got != want {
		t.Errorf("xs[2].Object = %v, want %v", got, want)
	}

	if got, want := xs[3].Object, s1; got != want {
		t.Errorf("xs[3].Object = %v, want %v", got, want)
	}
}
