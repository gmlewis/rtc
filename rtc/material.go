package rtc

// MaterialT represents a material.
type MaterialT struct {
	Color     Tuple
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

// Material returns a default material.
func Material() MaterialT {
	return MaterialT{
		Color:     Color(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
