package rtc

// Color returns a new Tuple as a color.
func Color(red, green, blue float64) Tuple {
	return Tuple{red, green, blue, 0}
}

// Red returns the red component of a color (Tuple).
func (t Tuple) Red() float64 {
	if len(t) == 0 {
		return 0
	}
	return t[0]
}

// Green returns the green component of a color (Tuple).
func (t Tuple) Green() float64 {
	if len(t) == 0 {
		return 0
	}
	return t[1]
}

// Blue returns the blue component of a color (Tuple).
func (t Tuple) Blue() float64 {
	if len(t) == 0 {
		return 0
	}
	return t[2]
}

// Alpha returns the alpha component of a color (Tuple).
func (t Tuple) Alpha() float64 {
	if len(t) == 0 {
		return 0
	}
	return t[3]
}

// HadamardProduct computes the product of two colors.
func (t Tuple) HadamardProduct(other Tuple) Tuple {
	return Color(
		t.X()*other.X(),
		t.Y()*other.Y(),
		t.Z()*other.Z(),
	)
}
