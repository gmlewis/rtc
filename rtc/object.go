package rtc

// Object is an interface that represents an object in the scene.
type Object interface {
	// Intersect returns a slice of IntersectionT values where the ray intersects the object.
	Intersect(ray RayT) []IntersectionT

	// Transform returns the object's transform 4x4 matrix.
	Transform() M4
	// SetTransform sets the object's transform 4x4 matrix.
	SetTransform(m M4)

	// Material returns the object's material.
	Material() *MaterialT
	// SetMaterial sets the object's material.
	SetMaterial(material MaterialT)

	// NormalAt returns the normal vector at the given point of intersection with the object.
	NormalAt(worldPoint Tuple) Tuple
}
