package rtc

// WorldT represents the world to be rendered.
type WorldT struct {
	Objects []Object
	Lights  []*PointLightT // TODO: Replace with light interfaces.
}

// World creates an empty world.
func World() *WorldT {
	return &WorldT{}
}

// DefaultWorld returns a default test world.
func DefaultWorld() *WorldT {
	s1 := Sphere()
	s1.Material().Color = Color(0.8, 1.0, 0.6)
	s1.Material().Diffuse = 0.7
	s1.Material().Specular = 0.2

	s2 := Sphere()
	s2.SetTransform(Scaling(0.5, 0.5, 0.5))

	return &WorldT{
		Objects: []Object{s1, s2},
		Lights:  []*PointLightT{PointLight(Point(-10, 10, -10), Color(1, 1, 1))},
	}
}
