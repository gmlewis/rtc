package rtc

import (
	"math"
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

// SetTransform sets the object's transform 4x4 matrix.
// Only for testing!
func (s *Shape) SetTransform(m M4) Object {
	s.transform = m
	return s
}

// SetMaterial sets the object's material.
// Only for testing!
func (s *Shape) SetMaterial(material MaterialT) Object {
	s.material = material
	return s
}

// SetParent sets the object's parent group.
// Only for testing!
func (s *Shape) SetParent(parent *GroupT) Object {
	s.parent = parent
	return s
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
// Only for testing!
func (s *Shape) Bounds() *BoundsT {
	return &BoundsT{
		Min: Point(math.Inf(-1), 0, math.Inf(-1)),
		Max: Point(math.Inf(1), 0, math.Inf(1)),
	}
}

// This is a global test function to save the ray.
var testFuncSaveRay func(ray RayT)

// LocalIntersect is for testing the testShape only.
func (s *Shape) LocalIntersect(ray RayT) []IntersectionT {
	testFuncSaveRay(ray)
	return nil
}

// LocalNormalAt is for testing the testShape only.
func (s *Shape) LocalNormalAt(localPoint Tuple, hit *IntersectionT) Tuple {
	return Vector(localPoint.X(), localPoint.Y(), localPoint.Z())
}

func TestShape_NewTestShape(t *testing.T) {
	ts := testShape()
	s := ts.shape
	if !s.Transform().Equal(M4Identity()) {
		t.Errorf("testShape default transform should be 4x4 identity matrix, got %v", s.Transform())
	}

	var want *GroupT
	if got := s.Parent(); got != want {
		t.Errorf("testShape parent = %v, want %v", got, want)
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
			testFuncSaveRay = func(ray RayT) { ts.savedRay = ray }
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

func TestShape_NormalAt_WithTransform(t *testing.T) {
	sq2 := math.Sqrt2 / 2

	tests := []struct {
		name      string
		transform M4
		point     Tuple
		want      Tuple
	}{
		{
			name:      "Computing the normal on a translated shape",
			transform: Translation(0, 1, 0),
			point:     Point(0, 1.70711, -0.70711),
			want:      Vector(0, 0.70711, -0.70711),
		},
		{
			name:      "Computing the normal on a transformed shape",
			transform: Scaling(1, 0.5, 1).Mult(RotationZ(math.Pi / 5)),
			point:     Point(0, sq2, -sq2),
			want:      Vector(0, 0.97014, -0.24254),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := testShape()
			s := ts.shape
			s.SetTransform(tt.transform)
			xs := Intersection(0, s)
			got := xs.NormalAt(tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("NormalAt = %v, want %v", got, tt.want)
			}
		})
	}
}
