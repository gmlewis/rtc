package rtc

// Pattern represents a color pattern.
type Pattern interface {
	// LocalPatternAt returns a color at a local point.
	LocalPatternAt(localPoint Tuple) Tuple

	// Transform returns the object's transform 4x4 matrix.
	Transform() M4
	// SetTransform sets the object's transform 4x4 matrix.
	SetTransform(m M4)
}

// BasePattern represents the common functionality for all patterns.
type BasePattern struct {
	transform M4
}

// Transform returns the object's transform 4x4 matrix.
func (s *BasePattern) Transform() M4 {
	return s.transform
}

// SetTransform sets the object's transform 4x4 matrix.
func (s *BasePattern) SetTransform(m M4) {
	s.transform = m
}

// PatternAt returns the pattern at the given point of intersection with the object.
func PatternAt(pattern Pattern, object Object, worldPoint Tuple) Tuple {
	localPoint := WorldToObject(object, worldPoint)
	patternPoint := pattern.Transform().Inverse().MultTuple(localPoint)
	return pattern.LocalPatternAt(patternPoint)
}
