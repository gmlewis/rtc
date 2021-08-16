package rtc

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTuples(t *testing.T) {
	tests := []struct {
		name       string
		tr         Tuple
		x          float64
		y          float64
		z          float64
		w          float64
		wantPoint  bool
		wantVector bool
	}{
		{
			name:       "nil Tuple",
			wantPoint:  false,
			wantVector: true,
		},
		{
			name:       "book scenario 1 page 27",
			tr:         Tuple{4.3, -4.2, 3.1, 1.0},
			x:          4.3,
			y:          -4.2,
			z:          3.1,
			w:          1.0,
			wantPoint:  true,
			wantVector: false,
		},
		{
			name:       "book scenario 2 page 27",
			tr:         Tuple{4.3, -4.2, 3.1, 0.0},
			x:          4.3,
			y:          -4.2,
			z:          3.1,
			w:          0.0,
			wantPoint:  false,
			wantVector: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.X(); math.Abs(got-tt.x) > epsilon {
				t.Errorf("tp.X() = %v, want %v", got, tt.x)
			}

			if got := tt.tr.Y(); math.Abs(got-tt.y) > epsilon {
				t.Errorf("tp.Y() = %v, want %v", got, tt.y)
			}

			if got := tt.tr.Z(); math.Abs(got-tt.z) > epsilon {
				t.Errorf("tp.Z() = %v, want %v", got, tt.z)
			}

			if got := tt.tr.W(); math.Abs(got-tt.w) > epsilon {
				t.Errorf("tp.W() = %v, want %v", got, tt.w)
			}

			if got := tt.tr.IsPoint(); got != tt.wantPoint {
				t.Errorf("tp.IsPoint() = %v, want %v", got, tt.wantPoint)
			}

			if got := tt.tr.IsVector(); got != tt.wantVector {
				t.Errorf("tp.IsVector() = %v, want %v", got, tt.wantVector)
			}
		})
	}
}

func TestPoint_Vector_Equal(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name       string
		args       args
		wantPoint  Tuple
		wantVector Tuple
	}{
		{
			name:       "book scenarios page 27",
			args:       args{4, -4, 3},
			wantPoint:  Tuple{4, -4, 3, 1},
			wantVector: Tuple{4, -4, 3, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPoint := Point(tt.args.x, tt.args.y, tt.args.z)
			if !cmp.Equal(gotPoint, tt.wantPoint) {
				t.Errorf("Point() = %v, want %v", gotPoint, tt.wantPoint)
			}
			if got := gotPoint.Equal(tt.wantPoint); !got {
				t.Errorf("gotPoint.Equal(wantPoint) = %v, want true", got)
			}
			if got := tt.wantPoint.Equal(gotPoint); !got {
				t.Errorf("wantPoint.Equal(gotPoint) = %v, want true", got)
			}

			gotVector := Vector(tt.args.x, tt.args.y, tt.args.z)
			if !cmp.Equal(gotVector, tt.wantVector) {
				t.Errorf("Vector() = %v, want %v", gotVector, tt.wantVector)
			}
			if got := gotVector.Equal(tt.wantVector); !got {
				t.Errorf("gotVector.Equal(wantVector) = %v, want true", got)
			}
			if got := tt.wantVector.Equal(gotVector); !got {
				t.Errorf("wantVector.Equal(gotVector) = %v, want true", got)
			}
		})
	}
}

func TestTuple_Add(t *testing.T) {
	tests := []struct {
		name  string
		tr    Tuple
		other Tuple
		want  Tuple
	}{
		{
			name:  "Adding two tuples",
			tr:    Tuple{3, -2, 5, 1},
			other: Tuple{-2, 3, 1, 0},
			want:  Tuple{1, 1, 6, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Add(tt.other); !cmp.Equal(got, tt.want) {
				t.Errorf("Tuple.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Sub(t *testing.T) {
	tests := []struct {
		name  string
		tr    Tuple
		other Tuple
		want  Tuple
	}{
		{
			name:  "Subtracting two points",
			tr:    Point(3, 2, 1),
			other: Point(5, 6, 7),
			want:  Vector(-2, -4, -6),
		},
		{
			name:  "Subtracting a vector from a point",
			tr:    Point(3, 2, 1),
			other: Vector(5, 6, 7),
			want:  Point(-2, -4, -6),
		},
		{
			name:  "Subtracting two vectors",
			tr:    Vector(3, 2, 1),
			other: Vector(5, 6, 7),
			want:  Vector(-2, -4, -6),
		},
		{
			name:  "Subtracting a vector from the zero vector",
			tr:    Vector(0, 0, 0),
			other: Vector(1, -2, 3),
			want:  Vector(-1, 2, -3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Sub(tt.other); !cmp.Equal(got, tt.want) {
				t.Errorf("Tuple.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Negate(t *testing.T) {
	tests := []struct {
		name string
		tr   Tuple
		want Tuple
	}{
		{
			name: "Negating a tuple",
			tr:   Tuple{1, -2, 3, -4},
			want: Tuple{-1, 2, -3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Negate(); !cmp.Equal(got, tt.want) {
				t.Errorf("Tuple.Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_MultScalar(t *testing.T) {
	tests := []struct {
		name string
		tr   Tuple
		f    float64
		want Tuple
	}{
		{
			name: "Multiplying a tuple by a scalar",
			tr:   Tuple{1, -2, 3, -4},
			f:    3.5,
			want: Tuple{3.5, -7, 10.5, -14},
		},
		{
			name: "Multiplying a tuple by a fraction",
			tr:   Tuple{1, -2, 3, -4},
			f:    0.5,
			want: Tuple{0.5, -1, 1.5, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.MultScalar(tt.f); !cmp.Equal(got, tt.want) {
				t.Errorf("Tuple.MultScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_DivScalar(t *testing.T) {
	tests := []struct {
		name string
		tr   Tuple
		f    float64
		want Tuple
	}{
		{
			name: "Dividing a tuple by a scalar",
			tr:   Tuple{1, -2, 3, -4},
			f:    2,
			want: Tuple{0.5, -1, 1.5, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.DivScalar(tt.f); !cmp.Equal(got, tt.want) {
				t.Errorf("Tuple.DivScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Magnitude(t *testing.T) {
	tests := []struct {
		name string
		tr   Tuple
		want float64
	}{
		{
			name: "Computing the magnitude of vector(1,0,0)",
			tr:   Vector(1, 0, 0),
			want: 1,
		},
		{
			name: "Computing the magnitude of vector(0,1,0)",
			tr:   Vector(0, 1, 0),
			want: 1,
		},
		{
			name: "Computing the magnitude of vector(0,0,1)",
			tr:   Vector(0, 0, 1),
			want: 1,
		},
		{
			name: "Computing the magnitude of vector(1,2,3)",
			tr:   Vector(1, 2, 3),
			want: math.Sqrt(14),
		},
		{
			name: "Computing the magnitude of vector(-1,-2,-3)",
			tr:   Vector(-1, -2, -3),
			want: math.Sqrt(14),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Magnitude(); math.Abs(got-tt.want) > epsilon {
				t.Errorf("Tuple.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Normalize(t *testing.T) {
	tests := []struct {
		name string
		tr   Tuple
		want Tuple
	}{
		{
			name: "Normalizing vector(4,0,0) gives (1,0,0)",
			tr:   Vector(4, 0, 0),
			want: Vector(1, 0, 0),
		},
		{
			name: "Normalizing vector(1,2,3)",
			tr:   Vector(1, 2, 3),
			want: Vector(0.267261, 0.534522, 0.801783),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Normalize()
			if !got.Equal(tt.want) {
				t.Errorf("Tuple.Normalize() = %v, want %v", got, tt.want)
			}

			if gotMag := got.Magnitude(); math.Abs(gotMag-1) > epsilon {
				t.Errorf("Magnitude of Normalized vector = %v, want 1", gotMag)
			}
		})
	}
}

func TestTuple_Dot(t *testing.T) {
	tests := []struct {
		name  string
		tr    Tuple
		other Tuple
		want  float64
	}{
		{
			name:  "The dot product of two tuples",
			tr:    Vector(1, 2, 3),
			other: Vector(2, 3, 4),
			want:  20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Dot(tt.other); got != tt.want {
				t.Errorf("Tuple.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Cross(t *testing.T) {
	tests := []struct {
		name  string
		tr    Tuple
		other Tuple
		want  Tuple
	}{
		{
			name:  "The cross product of two vectors",
			tr:    Vector(1, 2, 3),
			other: Vector(2, 3, 4),
			want:  Vector(-1, 2, -1),
		},
		{
			name:  "The cross product of two vectors, reversed",
			tr:    Vector(2, 3, 4),
			other: Vector(1, 2, 3),
			want:  Vector(1, -2, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Cross(tt.other); !got.Equal(tt.want) {
				t.Errorf("Tuple.Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Reflect(t *testing.T) {
	sq2 := math.Sqrt(2) / 2

	tests := []struct {
		name   string
		tr     Tuple
		normal Tuple
		want   Tuple
	}{
		{
			name:   "Reflecting a vector approaching at 45°",
			tr:     Vector(1, -1, 0),
			normal: Vector(0, 1, 0),
			want:   Vector(1, 1, 0),
		},
		{
			name:   "Reflecting a vector off a slanted surface",
			tr:     Vector(0, -1, 0),
			normal: Vector(sq2, sq2, 0),
			want:   Vector(1, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Reflect(tt.normal); !cmp.Equal(got, tt.want) {
				t.Errorf("Tuple.Reflect() = %v, want %v", got, tt.want)
			}
		})
	}
}
