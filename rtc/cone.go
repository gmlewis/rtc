package rtc

import (
	"math"
)

// Cone creates a cone at the origin with its axis on the Y axis.
// It implements the Object interface.
func Cone() *ConeT {
	return &ConeT{
		Shape:   Shape{Transform: M4Identity(), Material: GetMaterial()},
		Minimum: math.Inf(-1),
		Maximum: math.Inf(1),
		Closed:  false,
	}
}

// ConeT represents a Cone.
type ConeT struct {
	Shape
	Minimum float64
	Maximum float64
	Closed  bool
}

var _ Object = &ConeT{}

// SetTransform sets the object's transform 4x4 matrix.
func (c *ConeT) SetTransform(m M4) Object {
	c.Transform = m
	return c
}

// SetMaterial sets the object's material.
func (c *ConeT) SetMaterial(material MaterialT) Object {
	c.Material = material
	return c
}

// SetParent sets the object's parent object.
func (c *ConeT) SetParent(parent Object) Object {
	c.Parent = parent
	return c
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (c *ConeT) Bounds() *BoundsT {
	return &BoundsT{
		Min: Point(c.Minimum, c.Minimum, c.Minimum),
		Max: Point(c.Maximum, c.Maximum, c.Maximum),
	}
}

func (c *ConeT) intersectCaps(ray RayT, xs []IntersectionT) []IntersectionT {
	if !c.Closed || math.Abs(ray.Direction.Y()) < epsilon {
		return xs
	}

	t := (c.Minimum - ray.Origin.Y()) / ray.Direction.Y()
	if checkCap(ray, t, c.Minimum) { // Abs not necessary for radius here.
		xs = append(xs, Intersection(t, c))
	}

	t = (c.Maximum - ray.Origin.Y()) / ray.Direction.Y()
	if checkCap(ray, t, c.Maximum) { // Abs not necessary for radius here.
		xs = append(xs, Intersection(t, c))
	}

	return xs
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (c *ConeT) LocalIntersect(ray RayT) []IntersectionT {
	a := ray.Direction.X()*ray.Direction.X() - ray.Direction.Y()*ray.Direction.Y() + ray.Direction.Z()*ray.Direction.Z()
	b := 2*ray.Origin.X()*ray.Direction.X() - 2*ray.Origin.Y()*ray.Direction.Y() + 2*ray.Origin.Z()*ray.Direction.Z()
	if math.Abs(a) < epsilon && math.Abs(b) < epsilon {
		return c.intersectCaps(ray, nil)
	}

	c2 := ray.Origin.X()*ray.Origin.X() - ray.Origin.Y()*ray.Origin.Y() + ray.Origin.Z()*ray.Origin.Z()
	if math.Abs(a) < epsilon {
		t := -c2 / (2 * b)
		return c.intersectCaps(ray, []IntersectionT{Intersection(t, c)})
	}

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
func (c *ConeT) LocalNormalAt(objectPoint Tuple, hit *IntersectionT) Tuple {
	dist := objectPoint.X()*objectPoint.X() + objectPoint.Z()*objectPoint.Z()
	if dist < 1 && objectPoint.Y() >= c.Maximum-epsilon {
		return Vector(0, 1, 0)
	}
	if dist < 1 && objectPoint.Y() <= c.Minimum+epsilon {
		return Vector(0, -1, 0)
	}

	y := math.Sqrt(objectPoint.X()*objectPoint.X() + objectPoint.Z()*objectPoint.Z())
	if objectPoint.Y() > 0 {
		y = -y
	}

	return Vector(objectPoint.X(), y, objectPoint.Z())
}

// Includes returns whether this object includes (or actually is) the
// other object.
func (c *ConeT) Includes(other Object) bool {
	return c == other
}
