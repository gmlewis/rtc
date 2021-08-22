package rtc

// MaterialT represents a material.
type MaterialT struct {
	Color           Tuple
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Reflective      float64
	Transparency    float64
	RefractiveIndex float64
	Pattern         Pattern
}

// Material returns a default material.
func GetMaterial() MaterialT {
	return MaterialT{
		Color:           Color(1, 1, 1),
		Ambient:         0.1,
		Diffuse:         0.9,
		Specular:        0.9,
		Shininess:       200.0,
		Reflective:      0,
		Transparency:    0,
		RefractiveIndex: 1,
	}
}
