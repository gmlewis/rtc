package rtc

import (
	"testing"
)

func TestTranslation(t *testing.T) {
	transform := Translation(5, -3, 2)
	p := Point(-3, 4, 5)
	if got, want := transform.MultTuple(p), Point(2, 1, 7); got != want {
		t.Errorf("transform.MultTuple(p) = %v, want %v", got, want)
	}

	inv := transform.Inverse()
	if got, want := inv.MultTuple(p), Point(-8, 7, 3); got != want {
		t.Errorf("inv.MultTuple(p) = %v, want %v", got, want)
	}

	v := Vector(-3, 4, 5)
	if got := transform.MultTuple(v); !got.Equal(v) {
		t.Errorf("transform.MultTuple(v) = %v, want %v", got, v)
	}
}
