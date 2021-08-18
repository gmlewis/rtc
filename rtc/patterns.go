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

// RingPatternT is a pattern that draws rings.
// It implements the Pattern interface.
type RingPatternT struct {
	BasePattern
	a Tuple
	b Tuple
}

var _ Pattern = &RingPatternT{}

// RingPattern returns a RingPatternT.
func RingPattern(a, b Tuple) *RingPatternT {
	return &RingPatternT{
		BasePattern: BasePattern{transform: M4Identity()},
		a:           a,
		b:           b,
	}
}

// LocalPatternAt returns a color at a local point.
func (s *RingPatternT) LocalPatternAt(localPoint Tuple) Tuple {
	if int(math.Floor(math.Sqrt(localPoint.X()*localPoint.X()+localPoint.Z()*localPoint.Z())))%2 == 0 {
		return s.a
	}
	return s.b
}

// CheckersPatternT is a pattern that draws checkerss.
// It implements the Pattern interface.
type CheckersPatternT struct {
	BasePattern
	a Tuple
	b Tuple
}

var _ Pattern = &CheckersPatternT{}

// CheckersPattern returns a CheckersPatternT.
func CheckersPattern(a, b Tuple) *CheckersPatternT {
	return &CheckersPatternT{
		BasePattern: BasePattern{transform: M4Identity()},
		a:           a,
		b:           b,
	}
}

// LocalPatternAt returns a color at a local point.
func (s *CheckersPatternT) LocalPatternAt(localPoint Tuple) Tuple {
	t := int(math.Floor(localPoint.X())) + int(math.Floor(localPoint.Y())) + int(math.Floor(localPoint.Z()))
	if t%2 == 0 {
		return s.a
	}
	return s.b
}
