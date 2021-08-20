package rtc

import (
	"log"
)

// Group creates a group of objects at the origin.
// It implements the Object interface.
func Group(shapes ...Object) *GroupT {
	g := &GroupT{Shape: Shape{transform: M4Identity(), material: Material()}, bounds: Bounds()}
	g.AddChild(shapes...)
	return g
}

// AddChild adds shape(s) to a group.
func (g *GroupT) AddChild(shapes ...Object) {
	g.Children = append(g.Children, shapes...)
	for _, child := range shapes {
		child.SetParent(g)

		bc := child.Bounds()
		g.bounds.UpdateBounds(child.Transform().MultTuple(Point(bc.Min.X(), bc.Min.Y(), bc.Min.Z())))
		g.bounds.UpdateBounds(child.Transform().MultTuple(Point(bc.Max.X(), bc.Min.Y(), bc.Min.Z())))
		g.bounds.UpdateBounds(child.Transform().MultTuple(Point(bc.Max.X(), bc.Max.Y(), bc.Min.Z())))
		g.bounds.UpdateBounds(child.Transform().MultTuple(Point(bc.Min.X(), bc.Max.Y(), bc.Min.Z())))
		g.bounds.UpdateBounds(child.Transform().MultTuple(Point(bc.Min.X(), bc.Min.Y(), bc.Max.Z())))
		g.bounds.UpdateBounds(child.Transform().MultTuple(Point(bc.Max.X(), bc.Min.Y(), bc.Max.Z())))
		g.bounds.UpdateBounds(child.Transform().MultTuple(Point(bc.Max.X(), bc.Max.Y(), bc.Max.Z())))
		g.bounds.UpdateBounds(child.Transform().MultTuple(Point(bc.Min.X(), bc.Max.Y(), bc.Max.Z())))
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
	g.transform = m
	return g
}

// SetMaterial sets the object's material.
func (g *GroupT) SetMaterial(material MaterialT) Object {
	g.material = material
	return g
}

// SetParent sets the object's parent group.
func (g *GroupT) SetParent(parent *GroupT) Object {
	g.parent = parent
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
	b := g.Bounds() // consider cacheing bounds!
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
func (g *GroupT) LocalNormalAt(objectPoint Tuple, xs *IntersectionT) Tuple {
	log.Fatalf("programming error - groups are abstract and do not have normals")
	return Tuple{}
}
