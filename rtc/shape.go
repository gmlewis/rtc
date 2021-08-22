package rtc

// Shape represents the common functionality for all shapes.
type Shape struct {
	Transform M4
	Material  MaterialT
	Parent    Object
}

// Transform returns the object's transform 4x4 matrix.
func (s *Shape) GetTransform() M4 {
	return s.Transform
}

// Material returns the object's material.
func (s *Shape) GetMaterial() *MaterialT {
	return &s.Material
}

// Parent returns the object's parent object.
func (s *Shape) GetParent() Object {
	return s.Parent
}
