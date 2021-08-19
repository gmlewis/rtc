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
