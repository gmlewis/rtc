package rtc

import (
	"math"
	"reflect"
	"testing"
)

func TestMatrix_Construction4x4(t *testing.T) {
	m := M4{
		Tuple{1, 2, 3, 4},
		Tuple{5.5, 6.5, 7.5, 8.5},
		Tuple{9, 10, 11, 12},
		Tuple{13.5, 14.5, 15.5, 16.5},
	}

	assertValue := func(row, col int, want float64) {
		got := m.Get(row, col)
		if math.Abs(got-want) > epsilon {
			t.Errorf("m[%v,%v] = %v, want %v", row, col, got, want)
		}
	}

	assertValue(0, 0, 1)
	assertValue(0, 3, 4)
	assertValue(1, 0, 5.5)
	assertValue(1, 2, 7.5)
	assertValue(2, 2, 11)
	assertValue(3, 0, 13.5)
	assertValue(3, 2, 15.5)
}

func TestMatrix_Construction2x2(t *testing.T) {
	m := M2{
		Tuple{-3, 5},
		Tuple{1, -2},
	}

	assertValue := func(row, col int, want float64) {
		got := m.Get(row, col)
		if math.Abs(got-want) > epsilon {
			t.Errorf("m[%v,%v] = %v, want %v", row, col, got, want)
		}
	}

	assertValue(0, 0, -3)
	assertValue(0, 1, 5)
	assertValue(1, 0, 1)
	assertValue(1, 1, -2)
}

func TestMatrix_Construction3x3(t *testing.T) {
	m := M3{
		Tuple{-3, 5, 0},
		Tuple{1, -2, -7},
		Tuple{0, 1, 1},
	}

	assertValue := func(row, col int, want float64) {
		got := m.Get(row, col)
		if math.Abs(got-want) > epsilon {
			t.Errorf("m[%v,%v] = %v, want %v", row, col, got, want)
		}
	}

	assertValue(0, 0, -3)
	assertValue(0, 1, 5)
	assertValue(1, 0, 1)
	assertValue(1, 1, -2)
	assertValue(2, 2, 1)
}

func TestMatrix_Equal4x4(t *testing.T) {
	m1 := M4{
		Tuple{1, 2, 3, 4},
		Tuple{5, 6, 7, 8},
		Tuple{9, 8, 7, 6},
		Tuple{5, 4, 3, 2},
	}
	m2 := M4{
		Tuple{1, 2, 3, 4},
		Tuple{5, 6, 7, 8},
		Tuple{9, 8, 7, 6},
		Tuple{5, 4, 3, 2},
	}

	if !m1.Equal(m2) {
		t.Errorf("4x4 m1.Equal(m2) = %v, want true", false)
	}
	if !m2.Equal(m1) {
		t.Errorf("4x4 m2.Equal(m1) = %v, want true", false)
	}

	m3 := M4{
		Tuple{2, 3, 4, 5},
		Tuple{6, 7, 8, 9},
		Tuple{8, 7, 6, 5},
		Tuple{4, 3, 2, 1},
	}

	if m1.Equal(m3) {
		t.Errorf("4x4 m1.Equal(m3) = %v, want false", true)
	}
	if m3.Equal(m1) {
		t.Errorf("4x4 m3.Equal(m1) = %v, want false", true)
	}
}

func TestMatrix_Mult4x4(t *testing.T) {
	m1 := M4{
		Tuple{1, 2, 3, 4},
		Tuple{5, 6, 7, 8},
		Tuple{9, 8, 7, 6},
		Tuple{5, 4, 3, 2},
	}
	m2 := M4{
		Tuple{-2, 1, 2, 3},
		Tuple{3, 2, 1, -1},
		Tuple{4, 3, 6, 5},
		Tuple{1, 2, 7, 8},
	}

	want := M4{
		Tuple{20, 22, 50, 48},
		Tuple{44, 54, 114, 108},
		Tuple{40, 58, 110, 102},
		Tuple{16, 26, 46, 42},
	}
	if got := m1.Mult(m2); !got.Equal(want) {
		t.Errorf("4x4 m1.Mult(m2) = %v, want %v", got, want)
	}

	want = M4{
		Tuple{36, 30, 24, 18},
		Tuple{17, 22, 27, 32},
		Tuple{98, 94, 90, 86},
		Tuple{114, 102, 90, 78},
	}
	if got := m2.Mult(m1); !got.Equal(want) {
		t.Errorf("4x4 m2.Mult(m1) = %v, want %v", got, want)
	}
}

func TestMatrix_MultTuple(t *testing.T) {
	a := M4{
		Tuple{1, 2, 3, 4},
		Tuple{2, 4, 4, 2},
		Tuple{8, 6, 4, 1},
		Tuple{0, 0, 0, 1},
	}
	b := Tuple{1, 2, 3, 1}

	want := Tuple{18, 24, 33, 1}
	if got := a.MultTuple(b); !got.Equal(want) {
		t.Errorf("a.MultTuple(b) = %v, want %v", got, want)
	}
}

func TestMatrix_Mult_Identity(t *testing.T) {
	m1 := M4{
		Tuple{1, 2, 3, 4},
		Tuple{5, 6, 7, 8},
		Tuple{9, 8, 7, 6},
		Tuple{5, 4, 3, 2},
	}
	m2 := M4Identity()

	if got := m1.Mult(m2); !got.Equal(m1) {
		t.Errorf("4x4 m1.Mult(m2) = %v, want %v", got, m1)
	}

	t1 := Tuple{1, 2, 3, 4}
	if got := m2.MultTuple(t1); !got.Equal(t1) {
		t.Errorf("m2.Mult(t1) = %v, want %v", got, t1)
	}
}

func TestMatrix_Transpose(t *testing.T) {
	m1 := M4{
		Tuple{0, 9, 3, 0},
		Tuple{9, 8, 0, 8},
		Tuple{1, 8, 5, 3},
		Tuple{0, 0, 5, 8},
	}
	got := m1.Transpose()

	want := M4{
		Tuple{0, 9, 1, 0},
		Tuple{9, 8, 8, 0},
		Tuple{3, 0, 5, 5},
		Tuple{0, 8, 3, 8},
	}
	if !got.Equal(want) {
		t.Errorf("4x4 m1.Transpose = %v, want %v", got, want)
	}

	ident := M4Identity()
	if got := ident.Transpose(); !got.Equal(ident) {
		t.Errorf("4x4 ident.Transpose = %v, want %v", got, want)
	}
}

func TestM2_Determinant(t *testing.T) {
	tests := []struct {
		name string
		m    M2
		want float64
	}{
		{
			name: "Calculating the determinant of a 2x2 matrix",
			m:    M2{Tuple{1, 5}, Tuple{-3, 2}},
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Determinant(); got != tt.want {
				t.Errorf("M2.Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestM3_Submatrix(t *testing.T) {
	tests := []struct {
		name string
		m    M3
		row  int
		col  int
		want M2
	}{
		{
			name: "A submatrix of a 3x3 matrix is a 2x2 matrix",
			m:    M3{Tuple{1, 5, 0}, Tuple{-3, 2, 7}, Tuple{0, 6, -3}},
			row:  0,
			col:  2,
			want: M2{Tuple{-3, 2}, Tuple{0, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Submatrix(tt.row, tt.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("M3.Submatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestM4_Submatrix(t *testing.T) {
	tests := []struct {
		name string
		m    M4
		row  int
		col  int
		want M3
	}{
		{
			name: "A submatrix of a 4x4 matrix is a 3x3 matrix",
			m: M4{Tuple{-6, 1, 1, 6},
				Tuple{-8, 5, 8, 6},
				Tuple{-1, 0, 8, 2},
				Tuple{-7, 1, -1, 1}},
			row:  2,
			col:  1,
			want: M3{Tuple{-6, 1, 6}, Tuple{-8, 8, 6}, Tuple{-7, -1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Submatrix(tt.row, tt.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("M4.Submatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestM3_Minor(t *testing.T) {
	tests := []struct {
		name string
		m    M3
		row  int
		col  int
		want float64
	}{
		{
			name: "Calculating a minor of a 3x3 matrix",
			m: M3{Tuple{3, 5, 0},
				Tuple{2, -1, -7},
				Tuple{6, -1, 5}},
			row:  1,
			col:  0,
			want: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.m.Submatrix(tt.row, tt.col)
			if got := b.Determinant(); got != tt.want {
				t.Errorf("b.Determinant() = %v, want %v", got, tt.want)
			}

			if got := tt.m.Minor(tt.row, tt.col); got != tt.want {
				t.Errorf("M3.Minor() = %v, want %v", got, tt.want)
			}
		})
	}
}
