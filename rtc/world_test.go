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

func TestWorldT_ShadeHit(t *testing.T) {
	w := DefaultWorld()
	r := Ray(Point(0, 0, -5), Vector(0, 0, 1))
	shape := w.Objects[0]
	i := Intersection(4, shape)

	comps := i.PrepareComputations(r)

	if got, want := w.ShadeHit(comps), Color(0.38066, 0.47583, 0.2855); !got.Equal(want) {
		t.Errorf("Shading an intersection from the outside: WorldT.ShadeHit() = %v, want %v", got, want)
	}

	w.Lights = []*PointLightT{PointLight(Point(0, 0.25, 0), Color(1, 1, 1))}
	r = Ray(Point(0, 0, 0), Vector(0, 0, 1))
	shape = w.Objects[1]

	i = Intersection(0.5, shape)

	comps = i.PrepareComputations(r)

	if got, want := w.ShadeHit(comps), Color(0.90498, 0.90498, 0.90498); !got.Equal(want) {
		t.Errorf("Shading an intersection from the inside: WorldT.ShadeHit() = %v, want %v", got, want)
	}
}

func TestWorldT_ColorAt(t *testing.T) {
	w := DefaultWorld()

	tests := []struct {
		name         string
		ray          RayT
		outerAmbient float64
		innerAmbient float64
		want         Tuple
	}{
		{
			name:         "The color when a ray misses",
			ray:          Ray(Point(0, 0, -5), Vector(0, 1, 0)),
			outerAmbient: 0.1,
			innerAmbient: 0.1,
			want:         Color(0, 0, 0),
		},
		{
			name:         "The color when a ray hits",
			ray:          Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			outerAmbient: 0.1,
			innerAmbient: 0.1,
			want:         Color(0.38066, 0.47583, 0.2855),
		},
		{
			name:         "The color with an intersection behind the ray",
			ray:          Ray(Point(0, 0, 0.75), Vector(0, 0, -1)),
			outerAmbient: 1,
			innerAmbient: 1,
			want:         w.Objects[1].Material().Color,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w.Objects[0].Material().Ambient = tt.outerAmbient
			w.Objects[1].Material().Ambient = tt.innerAmbient

			if got := w.ColorAt(tt.ray); !got.Equal(tt.want) {
				t.Errorf("WorldT.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
