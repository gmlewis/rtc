package rtc

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSphere_Intersect(t *testing.T) {
	s := Sphere()

	tests := []struct {
		name string
		ray  RayT
		want []IntersectionT
	}{
		{
			name: "A ray intersects a sphere at two points",
			ray:  Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			want: []IntersectionT{Intersection(4.0, s), Intersection(6.0, s)},
		},
		{
			name: "A ray intersects a sphere at a tangent",
			ray:  Ray(Point(0, 1, -5), Vector(0, 0, 1)),
			want: []IntersectionT{Intersection(5.0, s), Intersection(5.0, s)},
		},
		{
			name: "A ray intersects misses a sphere",
			ray:  Ray(Point(0, 2, -5), Vector(0, 0, 1)),
			want: nil,
		},
		{
			name: "A ray intersects originates inside a sphere",
			ray:  Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			want: []IntersectionT{Intersection(-1.0, s), Intersection(1.0, s)},
		},
		{
			name: "A sphere is behind a ray",
			ray:  Ray(Point(0, 0, 5), Vector(0, 0, 1)),
			want: []IntersectionT{Intersection(-6.0, s), Intersection(-4.0, s)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersect(s, tt.ray)
			if len(got) != len(tt.want) {
				t.Fatalf("Sphere.Intersect() = %v, want %v", got, tt.want)
			}

			for i, w := range tt.want {
				if got[i].T != w.T || got[i].Object != w.Object {
					t.Errorf("Sphere.Intersect[%v] = %v, want %v", i, got[i], w)
				}
			}
		})
	}
}

func TestSphereT_Transform(t *testing.T) {
	s := Sphere()

	if got, want := s.Transform(), M4Identity(); got != want {
		t.Errorf("Sphere default transform = %v, want %v", got, want)
	}

	x := Translation(2, 3, 4)
	s.SetTransform(x)
	if got, want := s.Transform(), x; got != want {
		t.Errorf("Sphere modified transform = %v, want %v", got, want)
	}
}

func TestSphereT_Material(t *testing.T) {
	s := Sphere()

	if got, want := s.Material(), Material(); !cmp.Equal(got, &want) {
		t.Errorf("Sphere default material = %v, want %v", got, want)
	}

	m := Material()
	m.Ambient = 1
	s.SetMaterial(m)
	if got, want := s.Material(), m; !cmp.Equal(got, &want) {
		t.Errorf("Sphere modified material = %v, want %v", got, want)
	}
}

func TestSphere_Ray_Transform(t *testing.T) {
	tests := []struct {
		name string
		ray  RayT
		m    M4
		want []float64
	}{
		{
			name: "Intersecting a scaled sphere with a ray",
			ray:  Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			m:    Scaling(2, 2, 2),
			want: []float64{3, 7},
		},
		{
			name: "Intersecting a translated sphere with a ray",
			ray:  Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			m:    Translation(5, 0, 0),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sphere()
			s.SetTransform(tt.m)
			xs := Intersect(s, tt.ray)

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

func TestSphereT_NormalAt(t *testing.T) {
	s := Sphere()
	sq3 := math.Sqrt(3) / 3

	tests := []struct {
		name  string
		point Tuple
		want  Tuple
	}{
		{
			name:  "The normal on a sphere at a point on the x axis",
			point: Point(1, 0, 0),
			want:  Vector(1, 0, 0),
		},
		{
			name:  "The normal on a sphere at a point on the y axis",
			point: Point(0, 1, 0),
			want:  Vector(0, 1, 0),
		},
		{
			name:  "The normal on a sphere at a point on the z axis",
			point: Point(0, 0, 1),
			want:  Vector(0, 0, 1),
		},
		{
			name:  "The normal on a sphere at a nonaxial point",
			point: Point(sq3, sq3, sq3),
			want:  Vector(sq3, sq3, sq3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalAt(s, tt.point)

			if !cmp.Equal(got, tt.want) {
				t.Errorf("SphereT.NormalAt() = %v, want %v", got, tt.want)
			}

			if want := got.Normalize(); !got.Equal(want) {
				t.Errorf("got.Normalize = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphereT_NormalAt_WithTransform(t *testing.T) {
	sq2 := math.Sqrt2 / 2

	tests := []struct {
		name      string
		transform M4
		point     Tuple
		want      Tuple
	}{
		{
			name:      "Computing the normal on a translated sphere",
			transform: Translation(0, 1, 0),
			point:     Point(0, 1.70711, -0.70711),
			want:      Vector(0, 0.70711, -0.70711),
		},
		{
			name:      "Computing the normal on a transformed sphere",
			transform: Scaling(1, 0.5, 1).Mult(RotationZ(math.Pi / 5)),
			point:     Point(0, sq2, -sq2),
			want:      Vector(0, 0.97014, -0.24254),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sphere()
			s.SetTransform(tt.transform)
			got := NormalAt(s, tt.point)

			if !cmp.Equal(got, tt.want) {
				t.Errorf("SphereT.NormalAt() = %v, want %v", got, tt.want)
			}

			if want := got.Normalize(); !got.Equal(want) {
				t.Errorf("got.Normalize = %v, want %v", got, tt.want)
			}
		})
	}
}
