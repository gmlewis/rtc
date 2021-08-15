package rtc

// M4 is a 4x4 matrix.
type M4 [4]Tuple

// Get returns a value within the matrix.
func (m M4) Get(row, col int) float64 {
	return m[row][col]
}

// Equal tests if two matrices are equal.
func (m M4) Equal(other M4) bool {
	return m[0].Equal(other[0]) &&
		m[1].Equal(other[1]) &&
		m[2].Equal(other[2]) &&
		m[3].Equal(other[3])
}

// M3 is a 3x3 matrix.
type M3 [3]Tuple

// Get returns a value within the matrix.
func (m M3) Get(row, col int) float64 {
	return m[row][col]
}

// Equal tests if two matrices are equal.
func (m M3) Equal(other M3) bool {
	return m[0].Equal(other[0]) &&
		m[1].Equal(other[1]) &&
		m[2].Equal(other[2])
}

// M2 is a 2x2 matrix.
type M2 [2]Tuple

// Get returns a value within the matrix.
func (m M2) Get(row, col int) float64 {
	return m[row][col]
}

// Equal tests if two matrices are equal.
func (m M2) Equal(other M2) bool {
	return m[0].Equal(other[0]) &&
		m[1].Equal(other[1])
}
