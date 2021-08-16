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

func TestScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	p := Point(-4, 6, 8)
	if got, want := transform.MultTuple(p), Point(-8, 18, 32); got != want {
		t.Errorf("transform.MultTuple(p) = %v, want %v", got, want)
	}

	v := Vector(-4, 6, 8)
	if got, want := transform.MultTuple(v), Vector(-8, 18, 32); !got.Equal(want) {
		t.Errorf("transform.MultTuple(v) = %v, want %v", got, want)
	}

	inv := transform.Inverse()
	if got, want := inv.MultTuple(v), Vector(-2, 2, 2); got != want {
		t.Errorf("inv.MultTuple(v) = %v, want %v", got, want)
	}
}

func TestScaling_Reflection(t *testing.T) {
	transform := Scaling(-1, 1, 1)
	p := Point(2, 3, 4)
	if got, want := transform.MultTuple(p), Point(-2, 3, 4); got != want {
		t.Errorf("transform.MultTuple(p) = %v, want %v", got, want)
	}
}
