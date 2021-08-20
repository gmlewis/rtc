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
		Closed:  false,
	}
}

// CylinderT represents a Cylinder.
type CylinderT struct {
	Shape
	Minimum float64
	Maximum float64
	Closed  bool
}

var _ Object = &CylinderT{}

// SetTransform sets the object's transform 4x4 matrix.
func (c *CylinderT) SetTransform(m M4) Object {
	c.transform = m
	return c
}

// SetMaterial sets the object's material.
func (c *CylinderT) SetMaterial(material MaterialT) Object {
	c.material = material
	return c
}

// SetParent sets the object's parent group.
func (c *CylinderT) SetParent(parent *GroupT) Object {
	c.parent = parent
	return c
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (c *CylinderT) Bounds() *BoundsT {
	return &BoundsT{
		Min: Point(-1, c.Minimum, -1),
		Max: Point(1, c.Maximum, 1),
	}
}

func checkCap(ray RayT, t, radius float64) bool {
	x := ray.Origin.X() + t*ray.Direction.X()
	z := ray.Origin.Z() + t*ray.Direction.Z()
	return x*x+z*z <= radius*radius
}

func (c *CylinderT) intersectCaps(ray RayT, xs []IntersectionT) []IntersectionT {
	if !c.Closed || math.Abs(ray.Direction.Y()) < epsilon {
		return xs
	}

	t := (c.Minimum - ray.Origin.Y()) / ray.Direction.Y()
	if checkCap(ray, t, 1) {
		xs = append(xs, Intersection(t, c))
	}

	t = (c.Maximum - ray.Origin.Y()) / ray.Direction.Y()
	if checkCap(ray, t, 1) {
		xs = append(xs, Intersection(t, c))
	}

	return xs
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (c *CylinderT) LocalIntersect(ray RayT) []IntersectionT {
	a := ray.Direction.X()*ray.Direction.X() + ray.Direction.Z()*ray.Direction.Z()
	if math.Abs(a) < epsilon {
		return c.intersectCaps(ray, nil)
	}

	b := 2*ray.Origin.X()*ray.Direction.X() + 2*ray.Origin.Z()*ray.Direction.Z()
	c2 := ray.Origin.X()*ray.Origin.X() + ray.Origin.Z()*ray.Origin.Z() - 1
	discriminant := b*b - 4*a*c2

	if discriminant < 0 {
		return c.intersectCaps(ray, nil)
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
	xs = c.intersectCaps(ray, xs)

	return xs
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (c *CylinderT) LocalNormalAt(objectPoint Tuple, hit *IntersectionT) Tuple {
	dist := objectPoint.X()*objectPoint.X() + objectPoint.Z()*objectPoint.Z()
	if dist < 1 && objectPoint.Y() >= c.Maximum-epsilon {
		return Vector(0, 1, 0)
	}
	if dist < 1 && objectPoint.Y() <= c.Minimum+epsilon {
		return Vector(0, -1, 0)
	}

	return Vector(objectPoint.X(), 0, objectPoint.Z())
}
