package rtc

// Object is an interface that represents an object in the scene.
type Object interface {
	// LocalIntersect returns a slice of IntersectionT values where the
	// transformed (object space) ray intersects the object.
	LocalIntersect(ray RayT) []IntersectionT

	// Transform returns the object's transform 4x4 matrix.
	Transform() M4
	// SetTransform sets the object's transform 4x4 matrix.
	SetTransform(m M4)

	// Material returns the object's material.
	Material() *MaterialT
	// SetMaterial sets the object's material.
	SetMaterial(material MaterialT)

	// LocalNormalAt returns the normal vector at the given point of intersection
	// (transformed to object space) with the object.
	LocalNormalAt(localPoint Tuple) Tuple
}

// Intersect returns a slice of IntersectionT values where the ray intersects the object.
func Intersect(object Object, ray RayT) []IntersectionT {
	localRay := ray.Transform(object.Transform().Inverse())
	return object.LocalIntersect(localRay)
}

// NormalAt returns the normal vector at the given point of intersection with the object.
func NormalAt(object Object, worldPoint Tuple) Tuple {
	inv := object.Transform().Inverse()
	objectPoint := inv.MultTuple(worldPoint)
	objectNormal := object.LocalNormalAt(objectPoint)
	worldNormal := inv.Transpose().MultTuple(objectNormal)
	worldNormal[3] = 0 // W
	return worldNormal.Normalize()
}
