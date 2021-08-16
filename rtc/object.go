package rtc

import "sort"

// Object is an interface that represents an object in the scene.
type Object interface {
	Intersect(ray RayT) []IntersectionT
	Transform() M4
	SetTransform(m M4)
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
