package rtc

// Object is an interface that represents an object in the scene.
type Object interface {
	Intersect(ray RayT) []float64
}
