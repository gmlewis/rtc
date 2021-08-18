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

// IntersectWorld intersects a world with a ray.
func (w *WorldT) IntersectWorld(ray RayT) []IntersectionT {
	var result []IntersectionT
	for _, obj := range w.Objects {
		xs := Intersect(obj, ray)
		result = append(result, xs...)
	}
	return Intersections(result...) // Sorts them.
}

// ShadeHit returns the color (as a Tuple) for the precomputed intersection.
func (w *WorldT) ShadeHit(comps *Comps) Tuple {
	var result Tuple
	for _, light := range w.Lights {
		shadowed := w.IsShadowed(comps.OverPoint, light)
		surface := Lighting(comps.Object.Material(),
			comps.Object,
			light,
			comps.Point,
			comps.EyeVector,
			comps.NormalVector,
			shadowed,
		)
		reflected := w.ReflectedColor(comps)
		result = result.Add(surface).Add(reflected)
	}
	return result
}

// ColorAt returns the color (as a Tuple) when casting the given ray.
func (w *WorldT) ColorAt(ray RayT) Tuple {
	xs := w.IntersectWorld(ray)
	hit := Hit(xs)
	if hit == nil {
		return Color(0, 0, 0)
	}

	comps := hit.PrepareComputations(ray)
	return w.ShadeHit(comps)
}

// ViewTransform creates a camera transformation matrix.
// from and to are Points, and up is a Vector.
func ViewTransform(from, to, up Tuple) M4 {
	forward := to.Sub(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)
	orientation := M4{
		Tuple{left.X(), left.Y(), left.Z(), 0},
		Tuple{trueUp.X(), trueUp.Y(), trueUp.Z(), 0},
		Tuple{-forward.X(), -forward.Y(), -forward.Z(), 0},
		Tuple{0, 0, 0, 1},
	}
	return orientation.Mult(Translation(-from.X(), -from.Y(), -from.Z()))
}

// IsShadowed determines if the provided point is in a shadow for the given light.
func (w *WorldT) IsShadowed(point Tuple, light *PointLightT) bool {
	v := light.position.Sub(point)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := Ray(point, direction)

	intersections := w.IntersectWorld(r)

	h := Hit(intersections)

	return h != nil && h.T < distance
}

// ReflectedColor returns the reflected color for the precomputed intersection.
func (w *WorldT) ReflectedColor(comps *Comps) Tuple {
	if comps.Object.Material().Reflective == 0 {
		return Color(0, 0, 0)
	}

	reflectRay := Ray(comps.OverPoint, comps.ReflectVector)
	color := w.ColorAt(reflectRay)
	return color.MultScalar(comps.Object.Material().Reflective)
}
