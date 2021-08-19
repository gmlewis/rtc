package rtc

import (
	"math"
)

// Cylinder creates a cylinder at the origin with its axis on the Y axis.
// It implements the Object interface.
func Cylinder() *CylinderT {
	return &CylinderT{
		Shape:   Shape{transform: M4Identity(), material: Material()},
		Minimum: math.Inf(-1),
		Maximum: math.Inf(1),
	}
}

// CylinderT represents a Cylinder.
type CylinderT struct {
	Shape
	Minimum float64
	Maximum float64
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

	if t1 > t2 {
		t1, t2 = t2, t1
	}
	y1 := ray.Origin.Y() + t1*ray.Direction.Y()
	y2 := ray.Origin.Y() + t2*ray.Direction.Y()

	var xs []IntersectionT
	if c.Minimum < y1 && y1 < c.Maximum {
		xs = append(xs, Intersection(t1, c))
	}
	if c.Minimum < y2 && y2 < c.Maximum {
		xs = append(xs, Intersection(t2, c))
	}

	return xs
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (c *CylinderT) LocalNormalAt(objectPoint Tuple) Tuple {
	return Vector(objectPoint.X(), 0, objectPoint.Z())
}
