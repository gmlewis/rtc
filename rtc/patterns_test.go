package rtc

import (
	"testing"
)

func TestStripePattern(t *testing.T) {
	black := Color(0, 0, 0)
	white := Color(1, 1, 1)

	pattern := StripePattern(white, black)

	if got, want := pattern.a, white; !got.Equal(want) {
		t.Errorf("pattern.a = %v, want %v", got, want)
	}

	if got, want := pattern.b, black; !got.Equal(want) {
		t.Errorf("pattern.b = %v, want %v", got, want)
	}
}

func TestStripePatternT_LocalPatternAt(t *testing.T) {
	black := Color(0, 0, 0)
	white := Color(1, 1, 1)

	pattern := StripePattern(white, black)

	tests := []struct {
		name string
		p    Tuple
		want Tuple
	}{
		{
			name: "A stripe pattern is constant in y",
			p:    Point(0, 0, 0),
			want: white,
		},
		{
			name: "A stripe pattern is constant in y",
			p:    Point(0, 1, 0),
			want: white,
		},
		{
			name: "A stripe pattern is constant in y",
			p:    Point(0, 2, 0),
			want: white,
		},
		{
			name: "A stripe pattern is constant in z",
			p:    Point(0, 0, 1),
			want: white,
		},
		{
			name: "A stripe pattern is constant in z",
			p:    Point(0, 0, 2),
			want: white,
		},
		{
			name: "A stripe pattern alternates in x",
			p:    Point(0.9, 0, 0),
			want: white,
		},
		{
			name: "A stripe pattern alternates in x",
			p:    Point(1, 0, 0),
			want: black,
		},
		{
			name: "A stripe pattern alternates in x",
			p:    Point(1.9, 0, 0),
			want: black,
		},
		{
			name: "A stripe pattern alternates in x",
			p:    Point(-1, 0, 0),
			want: black,
		},
		{
			name: "A stripe pattern alternates in x",
			p:    Point(-1.1, 0, 0),
			want: white,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pattern.LocalPatternAt(tt.p); !got.Equal(tt.want) {
				t.Errorf("StripePatternT.PatternAt(%v) = %v, want %v", tt.p, got, tt.want)
			}
		})
	}
}

func TestStripePatternAt(t *testing.T) {
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
