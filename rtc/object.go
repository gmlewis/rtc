package rtc

// Object is an interface that represents an object in the scene.
type Object interface {
	// LocalIntersect returns a slice of IntersectionT values where the
	// transformed (object space) ray intersects the object.
	LocalIntersect(ray RayT) []IntersectionT

	// Parent returns the object's parent object.
	GetParent() Object
	// SetParent sets the object's parent object.
	SetParent(parent Object) Object

	// Transform returns the object's transform 4x4 matrix.
	GetTransform() M4
	// SetTransform sets the object's transform 4x4 matrix.
	SetTransform(m M4) Object

	// Material returns the object's material.
	GetMaterial() *MaterialT
	// SetMaterial sets the object's material.
	SetMaterial(material MaterialT) Object

	// LocalNormalAt returns the normal vector at the given point of intersection
	// (transformed to object space) with the object.
	LocalNormalAt(localPoint Tuple, hit *IntersectionT) Tuple

	// Bounds returns the minimum bounding box of the object in object
	// (untransformed) space.
	Bounds() *BoundsT

	// Includes returns whether this object includes (or actually is) the
	// other object.
	Includes(other Object) bool
}

// Intersect returns a slice of IntersectionT values where the ray intersects the object.
func Intersect(object Object, ray RayT) []IntersectionT {
	localRay := ray.Transform(object.GetTransform().Inverse())
	return object.LocalIntersect(localRay)
}

// UpdateTransformedBounds returns the updated bounding box of an object, taking
// into account its own transformation.
// If a starting bounding box is supplied, it is updated (expanded), otherwise a new
// one is returned.
func UpdateTransformedBounds(object Object, boundingBox *BoundsT) *BoundsT {
	if boundingBox == nil {
		boundingBox = Bounds()
	}

	bc := object.Bounds()
	boundingBox.UpdateBounds(object.GetTransform().MultTuple(Point(bc.Min.X(), bc.Min.Y(), bc.Min.Z())))
	boundingBox.UpdateBounds(object.GetTransform().MultTuple(Point(bc.Max.X(), bc.Min.Y(), bc.Min.Z())))
	boundingBox.UpdateBounds(object.GetTransform().MultTuple(Point(bc.Max.X(), bc.Max.Y(), bc.Min.Z())))
	boundingBox.UpdateBounds(object.GetTransform().MultTuple(Point(bc.Min.X(), bc.Max.Y(), bc.Min.Z())))
	boundingBox.UpdateBounds(object.GetTransform().MultTuple(Point(bc.Min.X(), bc.Min.Y(), bc.Max.Z())))
	boundingBox.UpdateBounds(object.GetTransform().MultTuple(Point(bc.Max.X(), bc.Min.Y(), bc.Max.Z())))
	boundingBox.UpdateBounds(object.GetTransform().MultTuple(Point(bc.Max.X(), bc.Max.Y(), bc.Max.Z())))
	boundingBox.UpdateBounds(object.GetTransform().MultTuple(Point(bc.Min.X(), bc.Max.Y(), bc.Max.Z())))

	return boundingBox
}
