package rtc

import "math"

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
	s1.GetMaterial().Color = Color(0.8, 1.0, 0.6)
	s1.GetMaterial().Diffuse = 0.7
	s1.GetMaterial().Specular = 0.2

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
func (w *WorldT) ShadeHit(comps *Comps, remaining int) Tuple {
	var result Tuple
	for _, light := range w.Lights {
		shadowed := w.IsShadowed(comps.OverPoint, light)
		surface := Lighting(comps.Object.GetMaterial(),
			comps.Object,
			light,
			comps.Point,
			comps.EyeVector,
			comps.NormalVector,
			shadowed,
		)

		reflected := w.ReflectedColor(comps, remaining)
		refracted := w.RefractedColor(comps, remaining)

		material := comps.Object.GetMaterial()
		if material.Reflective > 0 && material.Transparency > 0 {
			reflectance := comps.Schlick()
			result = result.Add(surface).Add(reflected.MultScalar(reflectance)).Add(refracted.MultScalar(1 - reflectance))
		} else {
			result = result.Add(surface).Add(reflected).Add(refracted)
		}
	}
	return result
}

// ColorAt returns the color (as a Tuple) when casting the given ray.
func (w *WorldT) ColorAt(ray RayT, remaining int) Tuple {
	xs := w.IntersectWorld(ray)
	hit := Hit(xs)
	if hit == nil {
		return Color(0, 0, 0)
	}

	comps := hit.PrepareComputations(ray, xs)
	return w.ShadeHit(comps, remaining)
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
func (w *WorldT) ReflectedColor(comps *Comps, remaining int) Tuple {
	if remaining < 1 || comps.Object.GetMaterial().Reflective == 0 {
		return Color(0, 0, 0)
	}

	reflectRay := Ray(comps.OverPoint, comps.ReflectVector)
	color := w.ColorAt(reflectRay, remaining-1)
	return color.MultScalar(comps.Object.GetMaterial().Reflective)
}

// RefractedColor returns the refracted color for the precomputed intersection.
func (w *WorldT) RefractedColor(comps *Comps, remaining int) Tuple {
	if remaining < 1 || comps.Object.GetMaterial().Transparency == 0 {
		return Color(0, 0, 0)
	}

	nRatio := comps.N1 / comps.N2                   // precompute?
	cosI := comps.EyeVector.Dot(comps.NormalVector) // precompute?
	sin2t := nRatio * nRatio * (1 - (cosI * cosI))

	if sin2t > 1 {
		return Color(0, 0, 0)
	}

	cosT := math.Sqrt(1 - sin2t)
	direction := comps.NormalVector.MultScalar(nRatio*cosI - cosT).Sub(comps.EyeVector.MultScalar(nRatio))

	refractedRay := Ray(comps.UnderPoint, direction)
	color := w.ColorAt(refractedRay, remaining-1)
	return color.MultScalar(comps.Object.GetMaterial().Transparency)
}

// WorldToObject converts a world-space point to object space, taking into
// account all the parents of the object.
func WorldToObject(object Object, point Tuple) Tuple {
	if p := object.GetParent(); p != nil {
		point = WorldToObject(p, point)
	}
	return object.GetTransform().Inverse().MultTuple(point)
}

// NormalToWorld converts an object-space normal to world space, taking into
// account all the parents of the object.
func NormalToWorld(object Object, normal Tuple) Tuple {
	inv := object.GetTransform().Inverse()
	worldNormal := inv.Transpose().MultTuple(normal)
	worldNormal[3] = 0 // W
	normal = worldNormal.Normalize()

	if p := object.GetParent(); p != nil {
		normal = NormalToWorld(p, normal)
	}
	return normal
}
