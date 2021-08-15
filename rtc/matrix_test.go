package rtc

import (
	"math"
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
