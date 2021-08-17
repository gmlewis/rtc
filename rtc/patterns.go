package rtc

import "math"

// StripePatternT is a pattern that draws stripes.
// It implements the Pattern interface.
type StripePatternT struct {
	BasePattern
	a Tuple
	b Tuple
}

var _ Pattern = &StripePatternT{}

// StripePattern returns a StripePatternT.
func StripePattern(a, b Tuple) *StripePatternT {
	return &StripePatternT{
		BasePattern: BasePattern{transform: M4Identity()},
		a:           a,
		b:           b,
	}
}

// LocalPatternAt returns a color at a local point.
func (s *StripePatternT) LocalPatternAt(localPoint Tuple) Tuple {
	if int(math.Floor(localPoint.X()))%2 == 0 {
		return s.a
	}
	return s.b
}

// GradientPatternT is a pattern that draws gradients.
// It implements the Pattern interface.
type GradientPatternT struct {
	BasePattern
	a        Tuple
	b        Tuple
	distance Tuple
}

var _ Pattern = &GradientPatternT{}

// GradientPattern returns a GradientPatternT.
func GradientPattern(a, b Tuple) *GradientPatternT {
	distance := b.Sub(a)
	return &GradientPatternT{
		BasePattern: BasePattern{transform: M4Identity()},
		a:           a,
		b:           b,
		distance:    distance,
	}
}

// LocalPatternAt returns a color at a local point.
func (s *GradientPatternT) LocalPatternAt(localPoint Tuple) Tuple {
	t := localPoint.X() - math.Floor(localPoint.X())
	return s.a.Add(s.distance.MultScalar(t))
}
