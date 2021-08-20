package rtc

import (
	"log"
	"math"
)

// Group creates a group of objects at the origin.
// It implements the Object interface.
func Group(shapes ...Object) *GroupT {
	g := &GroupT{Shape: Shape{transform: M4Identity(), material: Material()}}
	g.AddChild(shapes...)
	return g
}

// AddChild adds shape(s) to a group.
func (g *GroupT) AddChild(shapes ...Object) {
	g.Children = append(g.Children, shapes...)
	for _, shape := range shapes {
		shape.SetParent(g)
	}
}

// GroupT represents a group of objects with its own transformation matrix.
// It implements the Object interface.
type GroupT struct {
	Shape
	Children []Object
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
	b := &BoundsT{
		Min: Point(math.Inf(1), math.Inf(1), math.Inf(1)),
		Max: Point(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	}

	f := func(p Tuple) {
		if p.X() < b.Min.X() {
			b.Min[0] = p.X()
		}
		if p.Y() < b.Min.Y() {
			b.Min[1] = p.Y()
		}
		if p.Z() < b.Min.Z() {
			b.Min[2] = p.Z()
		}
		if p.X() > b.Max.X() {
			b.Max[0] = p.X()
		}
		if p.Y() > b.Max.Y() {
			b.Max[1] = p.Y()
		}
		if p.Z() > b.Max.Z() {
			b.Max[2] = p.Z()
		}
	}

	for _, child := range g.Children {
		bc := child.Bounds()
		f(child.Transform().MultTuple(Point(bc.Min.X(), bc.Min.Y(), bc.Min.Z())))
		f(child.Transform().MultTuple(Point(bc.Max.X(), bc.Min.Y(), bc.Min.Z())))
		f(child.Transform().MultTuple(Point(bc.Max.X(), bc.Max.Y(), bc.Min.Z())))
		f(child.Transform().MultTuple(Point(bc.Min.X(), bc.Max.Y(), bc.Min.Z())))
		f(child.Transform().MultTuple(Point(bc.Min.X(), bc.Min.Y(), bc.Max.Z())))
		f(child.Transform().MultTuple(Point(bc.Max.X(), bc.Min.Y(), bc.Max.Z())))
		f(child.Transform().MultTuple(Point(bc.Max.X(), bc.Max.Y(), bc.Max.Z())))
		f(child.Transform().MultTuple(Point(bc.Min.X(), bc.Max.Y(), bc.Max.Z())))
	}

	return b
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
func (g *GroupT) LocalNormalAt(objectPoint Tuple) Tuple {
	log.Fatalf("programming error - groups are abstract and do not have normals")
	return Tuple{}
}
