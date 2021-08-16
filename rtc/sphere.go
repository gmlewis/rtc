package rtc

import "math"

// Sphere creates a unit sphere at the origin. It implements the Object interface.
type Sphere struct {
}

var _ Object = Sphere{}

// Intersect returns the collection of t values where the ray intersects the object.
func (s Sphere) Intersect(ray RayT) []IntersectionT {
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
