package rtc

// PointLightT represents a point light.
type PointLightT struct {
	position  Tuple
	intensity Tuple
}

// PointLight returns a point light at the given position (a point Tuple) with
// the provided intensity (a color Tuple).
func PointLight(position Tuple, intensity Tuple) *PointLightT {
	return &PointLightT{position: position, intensity: intensity}
}
