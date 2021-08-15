package rtc

import (
	"math"
	"testing"
)

func TestColors(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tuple
		r    float64
		g    float64
		b    float64
		a    float64
	}{
		{
			name: "nil Color",
		},
		{
			name: "Colors are (red, green, blue) tuples",
			tr:   Color(-0.5, 0.4, 1.7),
			r:    -0.5,
			g:    0.4,
			b:    1.7,
			a:    1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Red(); math.Abs(got-tt.r) > epsilon {
				t.Errorf("tp.Red() = %v, want %v", got, tt.r)
			}

			if got := tt.tr.Green(); math.Abs(got-tt.g) > epsilon {
				t.Errorf("tp.Green() = %v, want %v", got, tt.g)
			}

			if got := tt.tr.Blue(); math.Abs(got-tt.b) > epsilon {
				t.Errorf("tp.Blue() = %v, want %v", got, tt.b)
			}

			if got := tt.tr.Alpha(); math.Abs(got-tt.a) > epsilon {
				t.Errorf("tp.Alpha() = %v, want %v", got, tt.a)
			}

			if tt.tr == nil {
				return
			}

			tup := *tt.tr

			if got := tup.Red(); math.Abs(got-tt.r) > epsilon {
				t.Errorf("tup.Red() = %v, want %v", got, tt.r)
			}

			if got := tup.Green(); math.Abs(got-tt.g) > epsilon {
				t.Errorf("tup.Green() = %v, want %v", got, tt.g)
			}

			if got := tup.Blue(); math.Abs(got-tt.b) > epsilon {
				t.Errorf("tup.Blue() = %v, want %v", got, tt.b)
			}

			if got := tup.Alpha(); math.Abs(got-tt.a) > epsilon {
				t.Errorf("tup.Alpha() = %v, want %v", got, tt.a)
			}
		})
	}
}
