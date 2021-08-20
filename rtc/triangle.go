package rtc

// Triangle returns a new TriangleT.
func Triangle(p1, p2, p3 Tuple) *TriangleT {
	e1 := p2.Sub(p1)
	e2 := p3.Sub(p1)
	normal := e2.Cross(e1).Normalize()

	return &TriangleT{
		Shape:  Shape{transform: M4Identity(), material: Material()},
		P1:     p1,
		P2:     p2,
		P3:     p3,
		E1:     e1,
		E2:     e2,
		Normal: normal,
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
	return &BoundsT{
		Min: Point(-1, -1, -1),
		Max: Point(1, 1, 1),
	}
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (t *TriangleT) LocalIntersect(ray RayT) []IntersectionT {
	return nil
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (t *TriangleT) LocalNormalAt(objectPoint Tuple) Tuple {
	return t.Normal
}
