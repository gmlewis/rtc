package rtc

import (
	"math"
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

func TestRotationX(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := RotationX(math.Pi / 4)

	sq2 := math.Sqrt(2) / 2
	if got, want := halfQuarter.MultTuple(p), Point(0, sq2, sq2); !got.Equal(want) {
		t.Errorf("halfQuarter.MultTuple(p) = %v, want %v", got, want)
	}

	fullQuarter := RotationX(math.Pi / 2)
	if got, want := fullQuarter.MultTuple(p), Point(0, 0, 1); !got.Equal(want) {
		t.Errorf("fullQuarter.MultTuple(p) = %v, want %v", got, want)
	}

	inv := halfQuarter.Inverse()
	if got, want := inv.MultTuple(p), Point(0, sq2, -sq2); !got.Equal(want) {
		t.Errorf("inv.MultTuple(p) = %v, want %v", got, want)
	}
}

func TestRotationY(t *testing.T) {
	p := Point(0, 0, 1)
	halfQuarter := RotationY(math.Pi / 4)

	sq2 := math.Sqrt(2) / 2
	if got, want := halfQuarter.MultTuple(p), Point(sq2, 0, sq2); !got.Equal(want) {
		t.Errorf("halfQuarter.MultTuple(p) = %v, want %v", got, want)
	}

	fullQuarter := RotationY(math.Pi / 2)
	if got, want := fullQuarter.MultTuple(p), Point(1, 0, 0); !got.Equal(want) {
		t.Errorf("fullQuarter.MultTuple(p) = %v, want %v", got, want)
	}

	inv := halfQuarter.Inverse()
	if got, want := inv.MultTuple(p), Point(-sq2, 0, sq2); !got.Equal(want) {
		t.Errorf("inv.MultTuple(p) = %v, want %v", got, want)
	}
}

func TestRotationZ(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := RotationZ(math.Pi / 4)

	sq2 := math.Sqrt(2) / 2
	if got, want := halfQuarter.MultTuple(p), Point(-sq2, sq2, 0); !got.Equal(want) {
		t.Errorf("halfQuarter.MultTuple(p) = %v, want %v", got, want)
	}

	fullQuarter := RotationZ(math.Pi / 2)
	if got, want := fullQuarter.MultTuple(p), Point(-1, 0, 0); !got.Equal(want) {
		t.Errorf("fullQuarter.MultTuple(p) = %v, want %v", got, want)
	}

	inv := halfQuarter.Inverse()
	if got, want := inv.MultTuple(p), Point(sq2, sq2, 0); !got.Equal(want) {
		t.Errorf("inv.MultTuple(p) = %v, want %v", got, want)
	}
}

func TestShearing(t *testing.T) {
	transform := Shearing(0, 1, 0, 0, 0, 0)
	p := Point(2, 3, 4)
	if got, want := transform.MultTuple(p), Point(6, 3, 4); got != want {
		t.Errorf("transform.MultTuple(p) = %v, want %v", got, want)
	}

	transform = Shearing(0, 0, 1, 0, 0, 0)
	if got, want := transform.MultTuple(p), Point(2, 5, 4); got != want {
		t.Errorf("transform.MultTuple(p) = %v, want %v", got, want)
	}

	transform = Shearing(0, 0, 0, 1, 0, 0)
	if got, want := transform.MultTuple(p), Point(2, 7, 4); got != want {
		t.Errorf("transform.MultTuple(p) = %v, want %v", got, want)
	}

	transform = Shearing(0, 0, 0, 0, 1, 0)
	if got, want := transform.MultTuple(p), Point(2, 3, 6); got != want {
		t.Errorf("transform.MultTuple(p) = %v, want %v", got, want)
	}

	transform = Shearing(0, 0, 0, 0, 0, 1)
	if got, want := transform.MultTuple(p), Point(2, 3, 7); got != want {
		t.Errorf("transform.MultTuple(p) = %v, want %v", got, want)
	}
}

func TestTransformation_Sequence(t *testing.T) {
	p := Point(1, 0, 1)
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)

	p2 := a.MultTuple(p)
	if want := Point(1, -1, 0); !p2.Equal(want) {
		t.Errorf("a.MultTuple(p) = %v, want %v", p2, want)
	}

	p3 := b.MultTuple(p2)
	if want := Point(5, -5, 0); !p3.Equal(want) {
		t.Errorf("b.MultTuple(p2) = %v, want %v", p3, want)
	}

	p4 := c.MultTuple(p3)
	if want := Point(15, 0, 7); !p4.Equal(want) {
		t.Errorf("c.MultTuple(p3) = %v, want %v", p4, want)
	}

	tx := c.Mult(b.Mult(a))
	if got, want := tx.MultTuple(p), Point(15, 0, 7); !got.Equal(want) {
		t.Errorf("tx.MultTuple(p) = %v, want %v", got, want)
	}

	tx2 := M4Identity().RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7)
	if got, want := tx2.MultTuple(p), Point(15, 0, 7); !got.Equal(want) {
		t.Errorf("tx2.MultTuple(p) = %v, want %v", got, want)
	}
}
