package rtc

import (
	"testing"
)

func TestPatternAt(t *testing.T) {
	black := Color(0, 0, 0)
	white := Color(1, 1, 1)

	tests := []struct {
		name             string
		objectTransform  M4
		patternTransform M4
		point            Tuple
		want             Tuple
	}{
		{
			name:             "Stripes with an object transformation",
			objectTransform:  Scaling(2, 2, 2),
			patternTransform: M4Identity(),
			point:            Point(1.5, 0, 0),
			want:             white,
		},
		{
			name:             "Stripes with an pattern transformation",
			objectTransform:  M4Identity(),
			patternTransform: Scaling(2, 2, 2),
			point:            Point(1.5, 0, 0),
			want:             white,
		},
		{
			name:             "Stripes with both an object and a pattern transformation",
			objectTransform:  Scaling(2, 2, 2),
			patternTransform: Translation(0.5, 0, 0),
			point:            Point(2.5, 0, 0),
			want:             white,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			object := Sphere()
			object.SetTransform(tt.objectTransform)
			pattern := StripePattern(white, black)
			pattern.SetTransform(tt.patternTransform)

			if got := PatternAt(pattern, object, tt.point); !got.Equal(tt.want) {
				t.Errorf("PatternAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
