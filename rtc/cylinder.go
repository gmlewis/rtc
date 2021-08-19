package rtc

import (
	"math"
)

// Cylinder creates a cylinder at the origin ranging from -1 to 1 on each axis.
// It implements the Object interface.
func Cylinder() *CylinderT {
	return &CylinderT{Shape{transform: M4Identity(), material: Material()}}
}

// CylinderT represents a Cylinder.
type CylinderT struct {
	Shape
}

var _ Object = &CylinderT{}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (c *CylinderT) LocalIntersect(ray RayT) []IntersectionT {
	a := ray.Direction.X()*ray.Direction.X() + ray.Direction.Z()*ray.Direction.Z()
	if math.Abs(a) < epsilon {
		return nil
	}

	b := 2*ray.Origin.X()*ray.Direction.X() + 2*ray.Origin.Z()*ray.Direction.Z()
	c2 := ray.Origin.X()*ray.Origin.X() + ray.Origin.Z()*ray.Origin.Z() - 1
	discriminant := b*b - 4*a*c2

	if discriminant < 0 {
		return nil
	}

	sr := math.Sqrt(discriminant)
	t1 := (-b - sr) / (2 * a)
	t2 := (-b + sr) / (2 * a)
	return []IntersectionT{Intersection(t1, c), Intersection(t2, c)}
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (c *CylinderT) LocalNormalAt(objectPoint Tuple) Tuple {
	return Vector(objectPoint.X(), 0, objectPoint.Z())
}
