package rtc

// Translation returns a 4x4 translation matrix.
func Translation(x, y, z float64) M4 {
	return M4{
		Tuple{1, 0, 0, x},
		Tuple{0, 1, 0, y},
		Tuple{0, 0, 1, z},
		Tuple{0, 0, 0, 1},
	}
}

// Translate translates a 4x4 matrix and returns a new one.
func (m M4) Translate(x, y, z float64) M4 {
	t := Translation(x, y, z)
	return t.Mult(m)
}