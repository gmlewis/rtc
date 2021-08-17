package rtc

import "math"

// Plane creates a plane at the origin on the X-Z axes and +Y is up.
// It implements the Object interface.
func Plane() *PlaneT {
	return &PlaneT{Shape{transform: M4Identity(), material: Material()}}
}

// PlaneT represents a Plane.
type PlaneT struct {
	Shape
}

var _ Object = &PlaneT{}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (s *PlaneT) LocalIntersect(ray RayT) []IntersectionT {
	if math.Abs(ray.Direction.Y()) < epsilon {
		return nil
	}

	t := -ray.Origin.Y() / ray.Direction.Y()
	return []IntersectionT{Intersection(t, s)}
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (s *PlaneT) LocalNormalAt(objectPoint Tuple) Tuple {
	return Vector(0, 1, 0)
}
