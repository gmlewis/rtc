package rtc

import (
	"testing"
)

// testPatternT is a wrapper around a BasePattern for testing purposes.
// It implements the Pattern interface.
type testPatternT struct {
	BasePattern
}

var _ Pattern = &testPatternT{}

func (t *testPatternT) LocalPatternAt(localPoint Tuple) Tuple {
	return Color(localPoint.X(), localPoint.Y(), localPoint.Z())
}

// testPattern creates a test BasePattern. It implements the Object interface.
func testPattern() *testPatternT {
	return &testPatternT{
		BasePattern{transform: M4Identity()},
	}
}

func TestPattern_NewTestPattern(t *testing.T) {
	ts := testPattern()
	s := ts.BasePattern
	if !s.Transform().Equal(M4Identity()) {
		t.Errorf("testPattern default transform should be 4x4 identity matrix, got %v", s.Transform())
	}
}

func TestPattern_SetTransform(t *testing.T) {
	ts := testPattern()
	s := ts.BasePattern
	s.SetTransform(Translation(1, 2, 3))
	if got, want := s.Transform(), Translation(1, 2, 3); !got.Equal(want) {
		t.Errorf("testPattern setTransform = %v, want %v", s.Transform(), want)
	}
}

func TestPatternAt(t *testing.T) {
	tests := []struct {
		name             string
		objectTransform  M4
		patternTransform M4
		point            Tuple
		want             Tuple
	}{
		{
			name:             "A pattern with an object transformation",
			objectTransform:  Scaling(2, 2, 2),
			patternTransform: M4Identity(),
			point:            Point(2, 3, 4),
			want:             Color(1, 1.5, 2),
		},
		{
			name:             "A pattern with an pattern transformation",
			objectTransform:  M4Identity(),
			patternTransform: Scaling(2, 2, 2),
			point:            Point(2, 3, 4),
			want:             Color(1, 1.5, 2),
		},
		{
			name:             "A pattern with both an object and a pattern transformation",
			objectTransform:  Scaling(2, 2, 2),
			patternTransform: Translation(0.5, 1, 1.5),
			point:            Point(2.5, 3, 3.5),
			want:             Color(0.75, 0.5, 0.25),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			object := Sphere()
			object.SetTransform(tt.objectTransform)
			pattern := testPattern()
			pattern.SetTransform(tt.patternTransform)

			if got := PatternAt(pattern, object, tt.point); !got.Equal(tt.want) {
				t.Errorf("PatternAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
