package rtc

// Shape represents the common functionality for all shapes.
// It embeds the Object interface.
type Shape struct {
	transform M4
	material  MaterialT
}

// Transform returns the object's transform 4x4 matrix.
func (s *Shape) Transform() M4 {
	return s.transform
}

// SetTransform sets the object's transform 4x4 matrix.
func (s *Shape) SetTransform(m M4) {
	s.transform = m
}

// Material returns the object's material.
func (s *Shape) Material() *MaterialT {
	return &s.material
}

// SetMaterial sets the object's material.
func (s *Shape) SetMaterial(material MaterialT) {
	s.material = material
}

// NormalAt returns the normal vector at the given point of intersection with the object.
func (s *Shape) NormalAt(worldPoint Tuple) Tuple {
	inv := s.transform.Inverse()
	objectPoint := inv.MultTuple(worldPoint)
	objectNormal := objectPoint.Sub(Point(0, 0, 0))
	worldNormal := inv.Transpose().MultTuple(objectNormal)
	worldNormal[3] = 0 // W
	return worldNormal.Normalize()
}
