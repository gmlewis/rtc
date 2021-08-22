package rtc

import (
	"log"
)

// Group creates a group of objects at the origin.
// It implements the Object interface.
func Group(shapes ...Object) *GroupT {
	g := &GroupT{Shape: Shape{Transform: M4Identity(), Material: GetMaterial()}, bounds: Bounds()}
	g.AddChild(shapes...)
	return g
}

// AddChild adds shape(s) to a group.
func (g *GroupT) AddChild(shapes ...Object) {
	g.Children = append(g.Children, shapes...)
	for _, child := range shapes {
		child.SetParent(g)
		UpdateTransformedBounds(child, g.bounds)
	}
}

// GroupT represents a group of objects with its own transformation matrix.
// It implements the Object interface.
type GroupT struct {
	Shape
	Children []Object

	bounds *BoundsT
}

var _ Object = &GroupT{}

// SetTransform sets the object's transform 4x4 matrix.
func (g *GroupT) SetTransform(m M4) Object {
	g.Transform = m
	return g
}

// SetMaterial sets the object's material.
func (g *GroupT) SetMaterial(material MaterialT) Object {
	g.Material = material
	return g
}

// SetParent sets the object's parent object.
func (g *GroupT) SetParent(parent Object) Object {
	g.Parent = parent
	return g
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (g *GroupT) Bounds() *BoundsT {
	return g.bounds
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (g *GroupT) LocalIntersect(ray RayT) []IntersectionT {
	b := g.Bounds()
	if xs := b.LocalIntersect(ray, g); len(xs) == 0 {
		return nil
	}

	var xs []IntersectionT
	for _, child := range g.Children {
		xs = append(xs, Intersect(child, ray)...)
	}
	return Intersections(xs...) // sort them
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (g *GroupT) LocalNormalAt(objectPoint Tuple, hit *IntersectionT) Tuple {
	log.Fatalf("programming error - groups are abstract and do not have normals")
	return Tuple{}
}

// Includes returns whether this object includes (or actually is) the
// other object.
func (g *GroupT) Includes(other Object) bool {
	for _, child := range g.Children {
		if child.Includes(other) {
			return true
		}
	}

	return false
}
