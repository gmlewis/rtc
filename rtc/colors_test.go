package rtc

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			a:    0.0,
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

func TestColor_Add(t *testing.T) {
	tests := []struct {
		name  string
		tr    *Tuple
		other *Tuple
		want  *Tuple
	}{
		{
			name:  "Adding colors",
			tr:    Color(0.9, 0.6, 0.75),
			other: Color(0.7, 0.1, 0.25),
			want:  Color(1.6, 0.7, 1.0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Add(tt.other); !cmp.Equal(got, tt.want) {
				t.Errorf("Color.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Sub(t *testing.T) {
	tests := []struct {
		name  string
		tr    *Tuple
		other *Tuple
		want  *Tuple
	}{
		{
			name:  "Subtracting colors",
			tr:    Color(0.9, 0.6, 0.75),
			other: Color(0.7, 0.1, 0.25),
			want:  Color(0.2, 0.5, 0.5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Sub(tt.other); !cmp.Equal(got, tt.want) {
				t.Errorf("Color.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_MulScalar(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tuple
		f    float64
		want *Tuple
	}{
		{
			name: "Multiplying a color by a scalar",
			tr:   Color(0.2, 0.3, 0.4),
			f:    2,
			want: Color(0.4, 0.6, 0.8),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.MulScalar(tt.f); !cmp.Equal(got, tt.want) {
				t.Errorf("Color.MulScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_HadamardProduct(t *testing.T) {
	tests := []struct {
		name  string
		tr    *Tuple
		other *Tuple
		want  *Tuple
	}{
		{
			name:  "Multiplying colors",
			tr:    Color(1, 0.2, 0.4),
			other: Color(0.9, 1, 0.1),
			want:  Color(0.9, 0.2, 0.04),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.HadamardProduct(tt.other); !got.Equal(tt.want) {
				t.Errorf("Color.HadamardProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
