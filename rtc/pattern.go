package rtc

// Pattern represents a color pattern.
type Pattern interface {
	// PatternAt returns a color at a local point.
	PatternAt(localPoint Tuple) Tuple
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
