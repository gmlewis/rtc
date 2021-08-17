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

func TestStripePatternT_PatternAt(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pattern.PatternAt(tt.p); !got.Equal(tt.want) {
				t.Errorf("StripePatternT.PatternAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
