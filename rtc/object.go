package rtc

// Object is an interface that represents an object in the scene.
type Object interface {
	Intersect(ray RayT) []float64
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

// Intersections returns a slice of IntersectionT.
func Intersections(args ...IntersectionT) []IntersectionT {
	return append([]IntersectionT{}, args...)
}
