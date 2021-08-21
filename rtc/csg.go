package rtc

import "log"

// CSGOperation represents a CSG operation.
type CSGOperation int

const (
	CSGUnion CSGOperation = iota
	CSGIntersection
	CSGDifference
)

// CSG represents a constructive solid geometry object.
func CSG(operation CSGOperation, left, right Object) *CSGT {
	c := &CSGT{
		Shape:     Shape{transform: M4Identity(), material: Material()},
		Operation: operation,
		Left:      left,
		Right:     right,
		bounds:    Bounds(),
	}
	left.SetParent(c)
	UpdateTransformedBounds(left, c.bounds)
	right.SetParent(c)
	UpdateTransformedBounds(right, c.bounds)
	return c
}

// CSGT represents a CSG object.
type CSGT struct {
	Shape
	Operation CSGOperation
	Left      Object
	Right     Object

	bounds *BoundsT
}

var _ Object = &CSGT{}

// SetTransform sets the object's transform 4x4 matrix.
func (c *CSGT) SetTransform(m M4) Object {
	c.transform = m
	return c
}

// SetMaterial sets the object's material.
func (c *CSGT) SetMaterial(material MaterialT) Object {
	c.material = material
	return c
}

// SetParent sets the object's parent object.
func (c *CSGT) SetParent(parent Object) Object {
	c.parent = parent
	return c
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (c *CSGT) Bounds() *BoundsT {
	return c.bounds
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (c *CSGT) LocalIntersect(ray RayT) []IntersectionT {
	b := c.Bounds()
	if xs := b.LocalIntersect(ray, c); len(xs) == 0 {
		return nil
	}

	var xs []IntersectionT
	xs = append(xs, Intersect(c.Left, ray)...)
	xs = append(xs, Intersect(c.Right, ray)...)
	return Intersections(xs...) // sort them
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (c *CSGT) LocalNormalAt(objectPoint Tuple, hit *IntersectionT) Tuple {
	log.Fatalf("programming error - groups are abstract and do not have normals")
	return Tuple{}
}

func intersectionAllowed(op CSGOperation, leftHit, inLeft, inRight bool) bool {
	if op == CSGUnion {
		return (leftHit && !inRight) || (!leftHit && !inLeft)
	}

	if op == CSGIntersection {
		return (leftHit && inRight) || (!leftHit && inLeft)
	}

	if op != CSGDifference {
		log.Fatalf("unknown CSG operation: %v", op)
	}

	return (leftHit && !inRight) || (!leftHit && inLeft)
}

// FilterIntersections filters allowed CSG intersections from all
// possible intersections.
func (c *CSGT) FilterIntersections(xs []IntersectionT) []IntersectionT {
	var inLeft, inRight bool
	var result []IntersectionT

	for _, x := range xs {
		leftHit := c.Left.Includes(x.Object)

		if intersectionAllowed(c.Operation, leftHit, inLeft, inRight) {
			result = append(result, x)
		}

		if leftHit {
			inLeft = !inLeft
		} else {
			inRight = !inRight
		}
	}

	return result
}

// Includes returns whether this object includes (or actually is) the
// other object.
func (c *CSGT) Includes(other Object) bool {
	return c.Left.Includes(other) || c.Right.Includes(other)
}
