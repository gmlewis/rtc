package rtc

import "math"

// Triangle returns a new TriangleT.
func Triangle(p1, p2, p3 Tuple) *TriangleT {
	e1 := p2.Sub(p1)
	e2 := p3.Sub(p1)
	normal := e2.Cross(e1).Normalize()

	bounds := Bounds()
	bounds.UpdateBounds(p1)
	bounds.UpdateBounds(p2)
	bounds.UpdateBounds(p3)

	return &TriangleT{
		Shape:  Shape{transform: M4Identity(), material: Material()},
		P1:     p1,
		P2:     p2,
		P3:     p3,
		E1:     e1,
		E2:     e2,
		Normal: normal,
		bounds: bounds,
	}
}

// TriangleT represents a triangle object.
type TriangleT struct {
	Shape
	P1 Tuple
	P2 Tuple
	P3 Tuple

	E1     Tuple
	E2     Tuple
	Normal Tuple

	bounds *BoundsT
}

var _ Object = &TriangleT{}

// SetTransform sets the object's transform 4x4 matrix.
func (t *TriangleT) SetTransform(m M4) Object {
	t.transform = m
	return t
}

// SetMaterial sets the object's material.
func (t *TriangleT) SetMaterial(material MaterialT) Object {
	t.material = material
	return t
}

// SetParent sets the object's parent group.
func (t *TriangleT) SetParent(parent *GroupT) Object {
	t.parent = parent
	return t
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (t *TriangleT) Bounds() *BoundsT {
	return t.bounds
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (t *TriangleT) LocalIntersect(ray RayT) []IntersectionT {
	dirCrossE2 := ray.Direction.Cross(t.E2)
	det := t.E1.Dot(dirCrossE2)
	if math.Abs(det) < epsilon {
		return nil
	}

	f := 1 / det
	p1ToOrigin := ray.Origin.Sub(t.P1)
	u := f * p1ToOrigin.Dot(dirCrossE2)
	if u < 0 || u > 1 {
		return nil
	}

	originCrossE1 := p1ToOrigin.Cross(t.E1)
	v := f * ray.Direction.Dot(originCrossE1)
	if v < 0 || u+v > 1 {
		return nil
	}

	tv := f * t.E2.Dot(originCrossE1)
	return Intersections(Intersection(tv, t))
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (t *TriangleT) LocalNormalAt(objectPoint Tuple) Tuple {
	return t.Normal
}
