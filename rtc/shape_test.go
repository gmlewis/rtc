package rtc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// testShapeT is a wrapper around a shape for testing purposes.
type testShapeT struct {
	shape    *Shape
	savedRay RayT
}

// testShape creates a test shape. It implements the Object interface.
func testShape() *testShapeT {
	return &testShapeT{
		shape: &Shape{transform: M4Identity(), material: Material()},
	}
}

// This is a global test function to save the ray.
var testFunc func(ray RayT)

// LocalIntersect is for testing the testShape only.
func (s *Shape) LocalIntersect(ray RayT) []IntersectionT {
	testFunc(ray)
	return nil
}

func TestShape_NewTestShape(t *testing.T) {
	ts := testShape()
	s := ts.shape
	if !s.Transform().Equal(M4Identity()) {
		t.Errorf("testShape default transform should be 4x4 identity matrix, got %v", s.Transform())
	}
}

func TestShape_SetTransform(t *testing.T) {
	ts := testShape()
	s := ts.shape
	s.SetTransform(Translation(2, 3, 4))
	if got, want := s.Transform(), Translation(2, 3, 4); !got.Equal(want) {
		t.Errorf("testShape setTransform = %v, want %v", s.Transform(), want)
	}
}

func TestShape_Material(t *testing.T) {
	ts := testShape()
	s := ts.shape

	if got, want := s.Material(), Material(); !cmp.Equal(got, &want) {
		t.Errorf("testShape default material = %v, want %v", got, want)
	}

	m := Material()
	m.Ambient = 1
	s.SetMaterial(m)
	if got, want := s.Material(), m; !cmp.Equal(got, &want) {
		t.Errorf("testShape modified material = %v, want %v", got, want)
	}
}

func TestShape_Ray_Transform(t *testing.T) {
	tests := []struct {
		name          string
		ray           RayT
		m             M4
		wantOrigin    Tuple
		wantDirection Tuple
	}{
		{
			name:          "Intersecting a scaled shape with a ray",
			ray:           Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			m:             Scaling(2, 2, 2),
			wantOrigin:    Point(0, 0, -2.5),
			wantDirection: Vector(0, 0, 0.5),
		},
		{
			name:          "Intersecting a translated shape with a ray",
			ray:           Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			m:             Translation(5, 0, 0),
			wantOrigin:    Point(-5, 0, -5),
			wantDirection: Vector(0, 0, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := testShape()
			s := ts.shape
			s.SetTransform(tt.m)
			testFunc = func(ray RayT) { ts.savedRay = ray }
			Intersect(s, tt.ray)

			if got, want := ts.savedRay.Origin, tt.wantOrigin; !cmp.Equal(got, want) {
				t.Errorf("ts.savedRay.Origin = %v, want %v", got, want)
			}

			if got, want := ts.savedRay.Direction, tt.wantDirection; !cmp.Equal(got, want) {
				t.Errorf("ts.savedRay.Direction = %v, want %v", got, want)
			}
		})
	}
}
