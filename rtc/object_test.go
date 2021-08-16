package rtc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntersection(t *testing.T) {
	s := Sphere{}
	i := Intersection(3.5, s)

	if got, want := i.T, 3.5; got != want {
		t.Errorf("i.T = %v, want %v", got, want)
	}

	if got, want := i.Object, s; got != want {
		t.Errorf("i.Object = %v, want %v", got, want)
	}
}

func TestIntersections(t *testing.T) {
	s := Sphere{}
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
	s := Sphere{}

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
			if got := Hit(tt.xs); !cmp.Equal(got, tt.want) {
				t.Errorf("Hit() = %v, want %v", got, tt.want)
			}
		})
	}
}
