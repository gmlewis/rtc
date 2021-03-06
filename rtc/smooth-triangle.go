package rtc

import "math"

// SmoothTriangle returns a new smooth SmoothTriangleT.
func SmoothTriangle(p1, p2, p3, n1, n2, n3 Tuple) *SmoothTriangleT {
	t := Triangle(p1, p2, p3)
	return &SmoothTriangleT{
		TriangleT: *t,
		N1:        n1,
		N2:        n2,
		N3:        n3,
	}
}

// SmoothTriangleT represents a smooth triangle object.
type SmoothTriangleT struct {
	TriangleT

	N1 Tuple
	N2 Tuple
	N3 Tuple
}

var _ Object = &SmoothTriangleT{}

// SetTransform sets the object's transform 4x4 matrix.
func (s *SmoothTriangleT) SetTransform(m M4) Object {
	s.Transform = m
	return s
}

// SetMaterial sets the object's material.
func (s *SmoothTriangleT) SetMaterial(material MaterialT) Object {
	s.Material = material
	return s
}

// SetParent sets the object's parent object.
func (s *SmoothTriangleT) SetParent(parent Object) Object {
	s.Parent = parent
	return s
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (s *SmoothTriangleT) Bounds() *BoundsT {
	return s.bounds
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (s *SmoothTriangleT) LocalIntersect(ray RayT) []IntersectionT {
	dirCrossE2 := ray.Direction.Cross(s.E2)
	det := s.E1.Dot(dirCrossE2)
	if math.Abs(det) < epsilon {
		return nil
	}

	f := 1 / det
	p1ToOrigin := ray.Origin.Sub(s.P1)
	u := f * p1ToOrigin.Dot(dirCrossE2)
	if u < 0 || u > 1 {
		return nil
	}

	originCrossE1 := p1ToOrigin.Cross(s.E1)
	v := f * ray.Direction.Dot(originCrossE1)
	if v < 0 || u+v > 1 {
		return nil
	}

	tv := f * s.E2.Dot(originCrossE1)
	return Intersections(IntersectionWithUV(tv, s, u, v))
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (s *SmoothTriangleT) LocalNormalAt(objectPoint Tuple, hit *IntersectionT) Tuple {
	return s.N2.MultScalar(hit.U).Add(s.N3.MultScalar(hit.V).Add(s.N1.MultScalar(1 - hit.U - hit.V)))
}

// Includes returns whether this object includes (or actually is) the
// other object.
func (s *SmoothTriangleT) Includes(other Object) bool {
	return s == other
}
