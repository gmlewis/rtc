package rtc

// Object is an interface that represents an object in the scene.
type Object interface {
	// LocalIntersect returns a slice of IntersectionT values where the
	// transformed (object space) ray intersects the object.
	LocalIntersect(ray RayT) []IntersectionT

	// Parent returns the object's parent group.
	Parent() *GroupT
	// SetParent sets the object's parent group.
	SetParent(parent *GroupT) Object

	// Transform returns the object's transform 4x4 matrix.
	Transform() M4
	// SetTransform sets the object's transform 4x4 matrix.
	SetTransform(m M4) Object

	// Material returns the object's material.
	Material() *MaterialT
	// SetMaterial sets the object's material.
	SetMaterial(material MaterialT) Object

	// LocalNormalAt returns the normal vector at the given point of intersection
	// (transformed to object space) with the object.
	LocalNormalAt(localPoint Tuple, xs *IntersectionT) Tuple

	// Bounds returns the minimum bounding box of the object in object
	// (untransformed) space.
	Bounds() *BoundsT
}

// Intersect returns a slice of IntersectionT values where the ray intersects the object.
func Intersect(object Object, ray RayT) []IntersectionT {
	localRay := ray.Transform(object.Transform().Inverse())
	return object.LocalIntersect(localRay)
}
