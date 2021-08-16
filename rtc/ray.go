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

// Position computes a Point a distance t along a given Ray.
func (r RayT) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.MultScalar(t))
}
