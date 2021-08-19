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

// SetTransform sets the object's transform 4x4 matrix.
func (p *PlaneT) SetTransform(m M4) Object {
	p.transform = m
	return p
}

// SetMaterial sets the object's material.
func (p *PlaneT) SetMaterial(material MaterialT) Object {
	p.material = material
	return p
}

// SetParent sets the object's parent group.
func (p *PlaneT) SetParent(parent *GroupT) Object {
	p.parent = parent
	return p
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (p *PlaneT) Bounds() *BoundsT {
	return &BoundsT{
		Min: Point(math.Inf(-1), 0, math.Inf(-1)),
		Max: Point(math.Inf(1), 0, math.Inf(1)),
	}
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (p *PlaneT) LocalIntersect(ray RayT) []IntersectionT {
	if math.Abs(ray.Direction.Y()) < epsilon {
		return nil
	}

	t := -ray.Origin.Y() / ray.Direction.Y()
	return []IntersectionT{Intersection(t, p)}
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (p *PlaneT) LocalNormalAt(objectPoint Tuple) Tuple {
	return Vector(0, 1, 0)
}
