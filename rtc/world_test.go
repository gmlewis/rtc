package rtc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWorld(t *testing.T) {
	w := World()

	if len(w.Objects) != 0 {
		t.Errorf("World should contain no objects, got %v", w.Objects)
	}

	if len(w.Lights) != 0 {
		t.Errorf("World should contain no lights, got %v", w.Lights)
	}
}

func TestDefaultWorld(t *testing.T) {
	w := DefaultWorld()

	if got, want := len(w.Objects), 2; got != want {
		t.Errorf("DefaultWorld got %v objects, want %v", got, want)
	}

	if got, want := len(w.Lights), 1; got != want {
		t.Errorf("DefaultWorld got %v lights, want %v", got, want)
	}
}

func TestWorldT_IntersectWorld(t *testing.T) {
	w := DefaultWorld()

	tests := []struct {
		name string
		ray  RayT
		want []float64
	}{
		{
			name: "Intersecting a world with a ray",
			ray:  Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			want: []float64{4, 4.5, 5.5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xs := w.IntersectWorld(tt.ray)

			if len(xs) != len(tt.want) {
				t.Fatalf("len(xs) = %v, want %v", len(xs), len(tt.want))
			}

			var got []float64
			for _, x := range xs {
				got = append(got, x.T)
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("xs = %v, want %v", got, tt.want)
			}
		})
	}
}
