package rtc

// MaterialT represents a material.
type MaterialT struct {
	color     Tuple
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

// Material returns a default material.
func Material() *MaterialT {
	return &MaterialT{
		color:     Color(1, 1, 1),
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}
