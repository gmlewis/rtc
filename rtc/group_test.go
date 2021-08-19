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
