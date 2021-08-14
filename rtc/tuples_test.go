package rtc

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTuples(t *testing.T) {
	tests := []struct {
		name       string
		tr         *Tuple
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
			tr:         &Tuple{4.3, -4.2, 3.1, 1.0},
			x:          4.3,
			y:          -4.2,
			z:          3.1,
			w:          1.0,
			wantPoint:  true,
			wantVector: false,
		},
		{
			name:       "book scenario 2 page 27",
			tr:         &Tuple{4.3, -4.2, 3.1, 0.0},
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

			if tt.tr == nil {
				return
			}

			tup := *tt.tr

			if got := tup.X(); math.Abs(got-tt.x) > epsilon {
				t.Errorf("tup.X() = %v, want %v", got, tt.x)
			}

			if got := tup.Y(); math.Abs(got-tt.y) > epsilon {
				t.Errorf("tup.Y() = %v, want %v", got, tt.y)
			}

			if got := tup.Z(); math.Abs(got-tt.z) > epsilon {
				t.Errorf("tup.Z() = %v, want %v", got, tt.z)
			}

			if got := tup.W(); math.Abs(got-tt.w) > epsilon {
				t.Errorf("tup.W() = %v, want %v", got, tt.w)
			}

			if got := tup.IsPoint(); got != tt.wantPoint {
				t.Errorf("tup.IsPoint() = %v, want %v", got, tt.wantPoint)
			}

			if got := tup.IsVector(); got != tt.wantVector {
				t.Errorf("tup.IsVector() = %v, want %v", got, tt.wantVector)
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
			if got := gotPoint.Equal(&tt.wantPoint); !got {
				t.Errorf("gotPoint.Equal(wantPoint) = %v, want true", got)
			}
			if got := tt.wantPoint.Equal(&gotPoint); !got {
				t.Errorf("wantPoint.Equal(gotPoint) = %v, want true", got)
			}

			gotVector := Vector(tt.args.x, tt.args.y, tt.args.z)
			if !cmp.Equal(gotVector, tt.wantVector) {
				t.Errorf("Vector() = %v, want %v", gotVector, tt.wantVector)
			}
			if got := gotVector.Equal(&tt.wantVector); !got {
				t.Errorf("gotVector.Equal(wantVector) = %v, want true", got)
			}
			if got := tt.wantVector.Equal(&gotVector); !got {
				t.Errorf("wantVector.Equal(gotVector) = %v, want true", got)
			}
		})
	}
}

func TestTuple_Add(t *testing.T) {
	tests := []struct {
		name  string
		tr    *Tuple
		other *Tuple
		want  Tuple
	}{
		{
			name:  "book scenario page 29",
			tr:    &Tuple{3, -2, 5, 1},
			other: &Tuple{-2, 3, 1, 0},
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
