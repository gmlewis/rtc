package rtc

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	s := Sphere()
	i := Intersection(3.5, s)

	if got, want := i.T, 3.5; got != want {
		t.Errorf("i.T = %v, want %v", got, want)
	}

	if got, want := i.Object, s; got != want {
		t.Errorf("i.Object = %v, want %v", got, want)
	}
}

func TestIntersections(t *testing.T) {
	s := Sphere()
	i1 := Intersection(1, s)
	i2 := Intersection(2, s)

	xs := Intersections(i1, i2)

	if len(xs) != 2 {
		t.Fatalf("len(xs) = %v, want 2", len(xs))
	}

	if got, want := xs[0].T, 1.0; got != want {
		t.Errorf("xs[0].T = %v, want %v", got, want)
	}

	if got, want := xs[1].T, 2.0; got != want {
		t.Errorf("xs[1].T = %v, want %v", got, want)
	}
}

func TestHit(t *testing.T) {
	s := Sphere()

	tests := []struct {
		name string
		xs   []IntersectionT
		want *IntersectionT
	}{
		{
			name: "The hit, when all intersections have positive t",
			xs:   Intersections(Intersection(2, s), Intersection(1, s)),
			want: &IntersectionT{T: 1, Object: s},
		},
		{
			name: "The hit, when some intersections have negative t",
			xs:   Intersections(Intersection(1, s), Intersection(-1, s)),
			want: &IntersectionT{T: 1, Object: s},
		},
		{
			name: "The hit, when all intersections have negative t",
			xs:   Intersections(Intersection(-1, s), Intersection(-2, s)),
			want: nil,
		},
		{
			name: "The hit is always the lowest nonnegative intersection",
			xs:   Intersections(Intersection(5, s), Intersection(7, s), Intersection(-3, s), Intersection(2, s)),
			want: &IntersectionT{T: 2, Object: s},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hit(tt.xs)

			if (got == nil && tt.want != nil) || (got != nil && tt.want == nil) {
				t.Fatalf("Hit = %v, want %v", got, tt.want)
			}

			if got == nil && tt.want == nil {
				return
			}

			if got.T != tt.want.T || got.Object != tt.want.Object {
				t.Errorf("Hit = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionT_PrepareComputations(t *testing.T) {
	shape := Sphere()

	tests := []struct {
		name string
		r    RayT
		i    IntersectionT
		want *Comps
	}{
		{
			name: "The hit, when an intersection occurs on the outside",
			r:    Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			i:    Intersection(4, shape),
			want: &Comps{
				T:             4,
				Object:        shape,
				Point:         Point(0, 0, -1),
				EyeVector:     Vector(0, 0, -1),
				NormalVector:  Vector(0, 0, -1),
				ReflectVector: Vector(0, 0, -1),
				Inside:        false,
				OverPoint:     Point(0, 0, -1.0001),
				N1:            1,
				N2:            1,
			},
		},
		{
			name: "The hit, when an intersection occurs on the inside",
			r:    Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			i:    Intersection(1, shape),
			want: &Comps{
				T:             1,
				Object:        shape,
				Point:         Point(0, 0, 1),
				EyeVector:     Vector(0, 0, -1),
				NormalVector:  Vector(0, 0, -1),
				ReflectVector: Vector(0, 0, -1),
				Inside:        true,
				OverPoint:     Point(0, 0, 0.9999),
				N1:            1,
				N2:            1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.i.PrepareComputations(tt.r, []IntersectionT{tt.i}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectionT.PrepareComputations() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestIntersectionT_PrepareComputations_OverPoint(t *testing.T) {
	r := Ray(Point(0, 0, -5), Vector(0, 0, 1))

	shape := Sphere()
	shape.SetTransform(Translation(0, 0, 1))
	i := Intersection(5, shape)

	comps := i.PrepareComputations(r, []IntersectionT{i})
	if comps.OverPoint.Z() >= -epsilon/2 {
		t.Errorf("comps.OverPoint.Z = %v, want >= %v", comps.OverPoint.Z(), -epsilon/2)
	}

	if comps.Point.Z() <= comps.OverPoint.Z() {
		t.Errorf("comps.Point.Z = %v, want <= %v", comps.Point.Z(), comps.OverPoint.Z())
	}
}

func TestIntersectionT_PrepareComputations_ReflectVector(t *testing.T) {
	sq2 := math.Sqrt2 / 2

	r := Ray(Point(0, 1, -1), Vector(0, -sq2, sq2))

	shape := Plane()
	i := Intersection(sq2, shape)

	comps := i.PrepareComputations(r, []IntersectionT{i})
	if got, want := comps.ReflectVector, Vector(0, sq2, sq2); !got.Equal(want) {
		t.Errorf("comps.ReflectVector = %v, want %v", got, want)
	}
}

func TestIntersectionT_PrepareComputations_N1N2(t *testing.T) {
	a := GlassSphere()
	a.SetTransform(Scaling(2, 2, 2))
	a.Material().RefractiveIndex = 1.5
	b := GlassSphere()
	b.SetTransform(Translation(0, 0, -0.25))
	b.Material().RefractiveIndex = 2.0
	c := GlassSphere()
	c.SetTransform(Translation(0, 0, 0.25))
	c.Material().RefractiveIndex = 2.5
	r := Ray(Point(0, 0, -4), Vector(0, 0, 1))

	xs := Intersections(Intersection(2, a), Intersection(2.75, b), Intersection(3.25, c), Intersection(4.75, b), Intersection(5.25, c), Intersection(6, a))

	tests := []struct {
		n1 float64
		n2 float64
	}{
		{n1: 1.0, n2: 1.5},
		{n1: 1.5, n2: 2.0},
		{n1: 2.0, n2: 2.5},
		{n1: 2.5, n2: 2.5},
		{n1: 2.5, n2: 1.5},
		{n1: 1.5, n2: 1.0},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("xs[%v]", i), func(t *testing.T) {
			comps := xs[i].PrepareComputations(r, xs)

			if got, want := comps.N1, tt.n1; got != want {
				t.Errorf("n1 = %#v, want %#v", got, want)
			}

			if got, want := comps.N2, tt.n2; got != want {
				t.Errorf("n2 = %#v, want %#v", got, want)
			}
		})
	}
}
