package rtc

import "log"

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
		xs := obj.Intersect(ray)
		result = append(result, xs...)
	}
	return Intersections(result...) // Sorts them.
}

// ShadeHit returns the color (as a Tuple) for the precomputed intersection.
func (w *WorldT) ShadeHit(comps *Comps) Tuple {
	var result Tuple
	for _, light := range w.Lights {
		color := Lighting(comps.Object.Material(),
			light,
			comps.Point,
			comps.EyeVector,
			comps.NormalVector,
		)
		result = result.Add(color)
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
	log.Printf("from=%v, to=%v, up=%v", from, to, up)
	forward := to.Sub(from).Normalize()
	log.Printf("forward=%v", forward)
	upn := up.Normalize()
	log.Printf("upn=%v", upn)
	left := forward.Cross(upn)
	log.Printf("left=%v", left)
	trueUp := left.Cross(forward)
	log.Printf("trueUp=%v", trueUp)
	orientation := M4{
		Tuple{left.X(), left.Y(), left.Z(), 0},
		Tuple{trueUp.X(), trueUp.Y(), trueUp.Z(), 0},
		Tuple{-forward.X(), -forward.Y(), -forward.Z(), 0},
		Tuple{0, 0, 0, 1},
	}
	log.Printf("orientation=%v", orientation)
	return orientation.Mult(Translation(-from.X(), -from.Y(), -from.Z()))
}
