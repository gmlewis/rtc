package rtc

import (
	"math"
	"testing"
)

func TestNormalAt_OnChildObject(t *testing.T) {
	s := Sphere().SetTransform(Translation(5, 0, 0))
	g2 := Group(s).SetTransform(Scaling(1, 2, 3))
	Group(g2).SetTransform(RotationY(math.Pi / 2))

	xs := Intersection(0, s)
	if got, want := xs.NormalAt(Point(1.7321, 1.1547, -5.5774)), Vector(0.2857, 0.4286, -0.8571); !got.Equal(want) {
		t.Errorf("NormalAt(s, Point(1.7321, 1.1547, -5.5774)) = %v, want %v", got, want)
	}
}
