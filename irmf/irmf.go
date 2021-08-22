// Package irmf supports loading and rendering IRMF objects.  See irmf.io.
package irmf

import (
	"github.com/gmlewis/rtc/rtc"
)

// IRMFT represents an IRMF object.
// It implements the rtc.Object interface.
type IRMFT struct {
	rtc.Shape
	bounds *rtc.BoundsT
}

var _ rtc.Object = &IRMFT{}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (i *IRMFT) LocalIntersect(ray rtc.RayT) []rtc.IntersectionT {
	return nil
}

// SetParent sets the object's parent object.
func (i *IRMFT) SetParent(parent rtc.Object) rtc.Object {
	return i
}

// SetTransform sets the object's transform 4x4 matrix.
func (i *IRMFT) SetTransform(m rtc.M4) rtc.Object {
	return i
}

// SetMaterial sets the object's material.
func (i *IRMFT) SetMaterial(material rtc.MaterialT) rtc.Object {
	return i
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (i *IRMFT) LocalNormalAt(localPoint rtc.Tuple, hit *rtc.IntersectionT) rtc.Tuple {
	return rtc.Vector(0, 1, 0)
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (i *IRMFT) Bounds() *rtc.BoundsT {
	return i.bounds
}

// Includes returns whether this object includes (or actually is) the
// other object.
func (i *IRMFT) Includes(other rtc.Object) bool {
	return i == other
}
