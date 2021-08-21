package rtc

import "math"

// Sphere creates a unit sphere at the origin.
// It implements the Object interface.
func Sphere() *SphereT {
	return &SphereT{Shape{transform: M4Identity(), material: Material()}}
}

// GlassSphere creates a unit glass sphere at the origin.
// It implements the Object interface.
func GlassSphere() *SphereT {
	m := Material()
	m.Transparency = 1
	m.RefractiveIndex = 1.5
	return &SphereT{Shape{transform: M4Identity(), material: m}}
}

// SphereT represents a sphere.
type SphereT struct {
	Shape
}

var _ Object = &SphereT{}

// SetTransform sets the object's transform 4x4 matrix.
func (s *SphereT) SetTransform(m M4) Object {
	s.transform = m
	return s
}

// SetMaterial sets the object's material.
func (s *SphereT) SetMaterial(material MaterialT) Object {
	s.material = material
	return s
}

// SetParent sets the object's parent object.
func (s *SphereT) SetParent(parent Object) Object {
	s.parent = parent
	return s
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (s *SphereT) Bounds() *BoundsT {
	return &BoundsT{
		Min: Point(-1, -1, -1),
		Max: Point(1, 1, 1),
	}
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (s *SphereT) LocalIntersect(ray RayT) []IntersectionT {
	sphereToRay := ray.Origin.Sub(Point(0, 0, 0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	sr := math.Sqrt(discriminant)
	t1 := (-b - sr) / (2 * a)
	t2 := (-b + sr) / (2 * a)
	return []IntersectionT{Intersection(t1, s), Intersection(t2, s)}
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (s *SphereT) LocalNormalAt(objectPoint Tuple, hit *IntersectionT) Tuple {
	return objectPoint.Sub(Point(0, 0, 0))
}

// Includes returns whether this object includes (or actually is) the
// other object.
func (s *SphereT) Includes(other Object) bool {
	return s == other
}
