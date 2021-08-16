package rtc

import "math"

// Sphere creates a unit sphere at the origin. It implements the Object interface.
func Sphere() *SphereT {
	return &SphereT{transform: M4Identity()}
}

// SphereT represents a sphere.
type SphereT struct {
	transform M4
}

var _ Object = &SphereT{}

// Intersect returns a slice of IntersectionT values where the ray intersects the object.
func (s *SphereT) Intersect(ray RayT) []IntersectionT {
	r := ray.Transform(s.Transform().Inverse())

	sphereToRay := r.Origin.Sub(Point(0, 0, 0))

	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	sr := math.Sqrt(discriminant)
	t1 := (-b - sr) / (2 * a)
	t2 := (-b + sr) / (2 * a)
	return []IntersectionT{Intersection(t1, s), Intersection(t2, s)}
}

// Transform returns the object's transform 4x4 matrix.
func (s *SphereT) Transform() M4 {
	return s.transform
}

// SetTransform sets the object's transform 4x4 matrix.
func (s *SphereT) SetTransform(m M4) {
	s.transform = m
}

// NormalAt returns the normal vector at the given point of intersection with the object.
func (s *SphereT) NormalAt(worldPoint Tuple) Tuple {
	inv := s.transform.Inverse()
	objectPoint := inv.MultTuple(worldPoint)
	objectNormal := objectPoint.Sub(Point(0, 0, 0))
	worldNormal := inv.Transpose().MultTuple(objectNormal)
	worldNormal[3] = 0 // W
	return worldNormal.Normalize()
}
