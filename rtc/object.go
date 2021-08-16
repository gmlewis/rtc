package rtc

import "sort"

// Object is an interface that represents an object in the scene.
type Object interface {
	// Intersect returns a slice of IntersectionT values where the ray intersects the object.
	Intersect(ray RayT) []IntersectionT

	// Transform returns the object's transform 4x4 matrix.
	Transform() M4
	// SetTransform sets the object's transform 4x4 matrix.
	SetTransform(m M4)

	// NormalAt returns the normal vector at the given point of intersection with the object.
	NormalAt(point Tuple) Tuple
}

// IntersectionT represents an intersection with an object.
type IntersectionT struct {
	T      float64
	Object Object
}

// Intersection returns an IntersectionT.
func Intersection(t float64, object Object) IntersectionT {
	return IntersectionT{
		T:      t,
		Object: object,
	}
}

// Intersections returns a slice of IntersectionT after sorting
// by intersection T values.
func Intersections(args ...IntersectionT) []IntersectionT {
	all := append([]IntersectionT{}, args...)
	sort.Slice(all, func(a, b int) bool {
		return all[a].T < all[b].T
	})
	return all
}

// Hit returns the first non-negative intersection.
// It assumes that the intersections have already been sorted by
// Intersections above.
func Hit(xs []IntersectionT) *IntersectionT {
	for _, x := range xs {
		if x.T > 0 {
			return &x
		}
	}
	return nil
}
