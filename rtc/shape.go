package rtc

// Shape represents the common functionality for all shapes.
type Shape struct {
	transform M4
	material  MaterialT
	parent    *GroupT
}

// Transform returns the object's transform 4x4 matrix.
func (s *Shape) Transform() M4 {
	return s.transform
}

// Material returns the object's material.
func (s *Shape) Material() *MaterialT {
	return &s.material
}

// Parent returns the object's parent group.
func (s *Shape) Parent() *GroupT {
	return s.parent
}
