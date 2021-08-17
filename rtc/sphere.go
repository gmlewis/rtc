package rtc

import "math"

// Sphere creates a unit sphere at the origin.
// It implements the Object interface.
func Sphere() *SphereT {
	return &SphereT{Shape{transform: M4Identity(), material: Material()}}
}

// SphereT represents a sphere.
type SphereT struct {
	Shape
}

var _ Object = &SphereT{}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (s *SphereT) LocalIntersect(ray RayT) []IntersectionT {
	sphereToRay := ray.Origin.Sub(Point(0, 0, 0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
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

// NormalAt returns the normal vector at the given point of intersection with the object.
func (s *SphereT) NormalAt(worldPoint Tuple) Tuple {
	inv := s.transform.Inverse()
	objectPoint := inv.MultTuple(worldPoint)
	objectNormal := objectPoint.Sub(Point(0, 0, 0))
	worldNormal := inv.Transpose().MultTuple(objectNormal)
	worldNormal[3] = 0 // W
	return worldNormal.Normalize()
}
