package rtc

// RayT represents an origin and a direction in 3D space.
type RayT struct {
	Origin    Tuple
	Direction Tuple
}

// Ray returns a new RayT.
func Ray(origin, direction Tuple) RayT {
	return RayT{Origin: origin, Direction: direction}
}
