package rtc

import (
	"math"
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

func TestWorldT_ShadeHit_From_Outside(t *testing.T) {
	w := DefaultWorld()
	r := Ray(Point(0, 0, -5), Vector(0, 0, 1))
	shape := w.Objects[0]
	i := Intersection(4, shape)

	comps := i.PrepareComputations(r)

	if got, want := w.ShadeHit(comps, maxReflections), Color(0.38066, 0.47583, 0.2855); !got.Equal(want) {
		t.Errorf("Shading an intersection from the outside: WorldT.ShadeHit() = %v, want %v", got, want)
	}
}

func TestWorldT_ShadeHit_From_Inside(t *testing.T) {
	w := DefaultWorld()
	w.Lights = []*PointLightT{PointLight(Point(0, 0.25, 0), Color(1, 1, 1))}
	r := Ray(Point(0, 0, 0), Vector(0, 0, 1))
	shape := w.Objects[1]

	i := Intersection(0.5, shape)

	comps := i.PrepareComputations(r)

	if got, want := w.ShadeHit(comps, maxReflections), Color(0.90498, 0.90498, 0.90498); !got.Equal(want) {
		t.Errorf("Shading an intersection from the inside: WorldT.ShadeHit() = %v, want %v", got, want)
	}
}

func TestWorldT_ShadeHit_In_Shadow(t *testing.T) {
	w := DefaultWorld()
	w.Lights = []*PointLightT{PointLight(Point(0, 0, -10), Color(1, 1, 1))}
	s1 := Sphere()
	s2 := Sphere()
	s2.SetTransform(Translation(0, 0, 10))
	w.Objects = append(w.Objects, s1, s2)

	r := Ray(Point(0, 0, 5), Vector(0, 0, 1))

	i := Intersection(4, s2)

	comps := i.PrepareComputations(r)

	if got, want := w.ShadeHit(comps, maxReflections), Color(0.1, 0.1, 0.1); !got.Equal(want) {
		t.Errorf("Shading an intersection in shadow: WorldT.ShadeHit() = %v, want %v", got, want)
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

			if got := w.ColorAt(tt.ray, maxReflections); !got.Equal(tt.want) {
				t.Errorf("WorldT.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestViewTransform(t *testing.T) {
	tests := []struct {
		name string
		from Tuple
		to   Tuple
		up   Tuple
		want M4
	}{
		{
			name: "The transformation matrix for the default orientation",
			from: Point(0, 0, 0),
			to:   Point(0, 0, -1),
			up:   Vector(0, 1, 0),
			want: M4Identity(),
		},
		{
			name: "A view transformation matrix looking in positive z direction",
			from: Point(0, 0, 0),
			to:   Point(0, 0, 1),
			up:   Vector(0, 1, 0),
			want: Scaling(-1, 1, -1),
		},
		{
			name: "The view transformation moves the world",
			from: Point(0, 0, 8),
			to:   Point(0, 0, 0),
			up:   Vector(0, 1, 0),
			want: Translation(0, 0, -8),
		},
		{
			name: "An arbitrary view transformation",
			from: Point(1, 3, 2),
			to:   Point(4, -2, 8),
			up:   Vector(1, 1, 0),
			want: M4{
				Tuple{-0.50709, 0.50709, 0.67612, -2.36643},
				Tuple{0.76772, 0.60609, 0.12122, -2.82843},
				Tuple{-0.35857, 0.59761, -0.71714, 0.00000},
				Tuple{0.00000, 0.00000, 0.00000, 1.00000},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ViewTransform(tt.from, tt.to, tt.up); !cmp.Equal(got, tt.want) {
				t.Errorf("ViewTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorldT_IsShadowed(t *testing.T) {
	w := DefaultWorld()

	tests := []struct {
		name  string
		point Tuple
		want  bool
	}{
		{
			name:  "There is no shadow when nothing is collinear with point and light",
			point: Point(0, 10, 0),
			want:  false,
		},
		{
			name:  "The shadow when an object is between the point and the light",
			point: Point(10, -10, 10),
			want:  true,
		},
		{
			name:  "There is no shadow when an object is behind the light",
			point: Point(-20, 20, -20),
			want:  false,
		},
		{
			name:  "There is no shadow when an object is behind the point",
			point: Point(-2, 2, -2),
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := w.IsShadowed(tt.point, w.Lights[0]); got != tt.want {
				t.Errorf("WorldT.IsShadowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorldT_ReflectedColor(t *testing.T) {
	w := DefaultWorld()
	r := Ray(Point(0, 0, 0), Vector(0, 0, 1))
	shape := w.Objects[1]
	shape.Material().Ambient = 1
	i := Intersection(1, shape)

	comps := i.PrepareComputations(r)

	if got, want := w.ReflectedColor(comps, maxReflections), Color(0, 0, 0); !got.Equal(want) {
		t.Errorf("w.ReflectedColor = %v, want %v", got, want)
	}
}

func TestWorldT_ReflectedColor_WithReflectiveMaterial(t *testing.T) {
	sq2 := math.Sqrt2 / 2
	w := DefaultWorld()
	shape := Plane()
	shape.Material().Reflective = 0.5
	shape.SetTransform(Translation(0, -1, 0))
	w.Objects = append(w.Objects, shape)
	r := Ray(Point(0, 0, -3), Vector(0, -sq2, sq2))
	i := Intersection(math.Sqrt2, shape)

	comps := i.PrepareComputations(r)
	if got, want := w.ReflectedColor(comps, maxReflections), Color(0.19032, 0.2379, 0.14274); !got.Equal(want) {
		t.Errorf("w.ReflectedColor = %v, want %v", got, want)
	}
}

func TestWorldT_ShadeHit_WithReflectiveMaterial(t *testing.T) {
	sq2 := math.Sqrt2 / 2
	w := DefaultWorld()
	shape := Plane()
	shape.Material().Reflective = 0.5
	shape.SetTransform(Translation(0, -1, 0))
	w.Objects = append(w.Objects, shape)
	r := Ray(Point(0, 0, -3), Vector(0, -sq2, sq2))
	i := Intersection(math.Sqrt2, shape)

	comps := i.PrepareComputations(r)
	if got, want := w.ShadeHit(comps, maxReflections), Color(0.87677, 0.92436, 0.82918); !got.Equal(want) {
		t.Errorf("w.ShadeHit = %v, want %v", got, want)
	}
}

func TestWorldT_ShadeHit_WithMutuallyReflectiveSurfaces(t *testing.T) {
	w := World()
	w.Lights = []*PointLightT{PointLight(Point(0, 0, 0), Color(1, 1, 1))}
	lower := Plane()
	lower.Material().Reflective = 1
	lower.SetTransform(Translation(0, -1, 0))
	upper := Plane()
	upper.Material().Reflective = 1
	upper.SetTransform(Translation(0, 1, 0))
	w.Objects = append(w.Objects, lower, upper)
	r := Ray(Point(0, 0, 0), Vector(0, 1, 0))

	if got, notWant := w.ColorAt(r, maxReflections), Color(0, 0, 0); got.Equal(notWant) {
		t.Errorf("w.ShadeHit = %v, notWant %v", got, notWant)
	}
}

func TestWorldT_ReflectedColor_WithMaximumRecursionDepth(t *testing.T) {
	sq2 := math.Sqrt2 / 2
	w := DefaultWorld()
	shape := Plane()
	shape.Material().Reflective = 0.5
	shape.SetTransform(Translation(0, -1, 0))
	w.Objects = append(w.Objects, shape)
	r := Ray(Point(0, 0, -3), Vector(0, -sq2, sq2))
	i := Intersection(math.Sqrt2, shape)

	comps := i.PrepareComputations(r)
	if got, want := w.ReflectedColor(comps, 0), Color(0, 0, 0); !got.Equal(want) {
		t.Errorf("w.ReflectedColor = %v, want %v", got, want)
	}
}
