package rtc

import "sort"

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

// Comps contains precomputed information about an intersection.
type Comps struct {
	T            float64
	Object       Object
	Point        Tuple
	EyeVector    Tuple
	NormalVector Tuple
	Inside       bool
}

// PrepareComputations returns a new data structure encapsulating information
// about the intersection.
func (i IntersectionT) PrepareComputations(ray RayT) *Comps {
	point := ray.Position(i.T)
	eyeVector := ray.Direction.Negate()
	normalVector := i.Object.NormalAt(point)
	var inside bool
	if normalVector.Dot(eyeVector) < 0 {
		inside = true
		normalVector = normalVector.Negate()
	}

	return &Comps{
		T:            i.T,
		Object:       i.Object,
		Point:        point,
		EyeVector:    eyeVector,
		NormalVector: normalVector,
		Inside:       inside,
	}
}
