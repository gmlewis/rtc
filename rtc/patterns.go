package rtc

import "math"

// StripePatternT is a pattern that draws stripes.
// It implements the Pattern interface.
type StripePatternT struct {
	BasePattern
	a Tuple
	b Tuple
}

// StripePattern returns a StripePatternT.
func StripePattern(a, b Tuple) *StripePatternT {
	return &StripePatternT{
		BasePattern: BasePattern{transform: M4Identity()},
		a:           a,
		b:           b,
	}
}

var _ Pattern = &StripePatternT{}

// LocalPatternAt returns a color at a local point.
func (s *StripePatternT) LocalPatternAt(localPoint Tuple) Tuple {
	if int(math.Floor(localPoint.X()))%2 == 0 {
		return s.a
	}
	return s.b
}
