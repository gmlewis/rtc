package rtc

import (
	"math"
	"sort"
)

// IntersectionT represents an intersection with an object.
type IntersectionT struct {
	T      float64
	Object Object
}

// Intersection returns an IntersectionT.
func Intersection(t float64, object Object) IntersectionT {
	return IntersectionT{
		T:      t,
		Object: object,
	}
}

// Intersections returns a slice of IntersectionT after sorting
// by intersection T values.
func Intersections(args ...IntersectionT) []IntersectionT {
	all := append([]IntersectionT{}, args...)
	sort.Slice(all, func(a, b int) bool {
		return all[a].T < all[b].T
	})
	return all
}

// Hit returns the first non-negative intersection.
// It assumes that the intersections have already been sorted by
// Intersections above.
func Hit(xs []IntersectionT) *IntersectionT {
	for _, x := range xs {
		if x.T > 0 {
			return &x
		}
	}
	return nil
}

// Comps contains precomputed information about an intersection.
type Comps struct {
	T             float64
	Object        Object
	Point         Tuple
	EyeVector     Tuple
	NormalVector  Tuple
	ReflectVector Tuple
	Inside        bool
	OverPoint     Tuple   // For shadow testing - slightly above surface of object.
	UnderPoint    Tuple   // For transparency and index of refraction calculations.
	N1            float64 // Refractive index of material being exited.
	N2            float64 // Refractive index of material being entered.
}

// PrepareComputations returns a new data structure encapsulating information
// about the intersection.
func (i IntersectionT) PrepareComputations(ray RayT, xs []IntersectionT) *Comps {
	point := ray.Position(i.T)
	eyeVector := ray.Direction.Negate()
	normalVector := NormalAt(i.Object, point)
	var inside bool
	if normalVector.Dot(eyeVector) < 0 {
		inside = true
		normalVector = normalVector.Negate()
	}
	reflectVector := ray.Direction.Reflect(normalVector)
	eps := normalVector.MultScalar(epsilon)
	overPoint := point.Add(eps)
	underPoint := point.Sub(eps)

	n1, n2 := 1.0, 1.0
	var containers []Object
	indexOf := func(x Object) int {
		for i, c := range containers {
			if c == x {
				return i
			}
		}
		return -1
	}

	for _, x := range xs {
		if x.T == i.T {
			if len(containers) == 0 {
				n1 = 1.0
			} else {
				n1 = containers[len(containers)-1].Material().RefractiveIndex
			}
		}

		// removeOrAppend(x.Object)
		if index := indexOf(x.Object); index >= 0 {
			containers = append(containers[:index], containers[index+1:]...)
		} else {
			containers = append(containers, x.Object)
		}

		if x.T == i.T {
			if len(containers) == 0 {
				n2 = 1.0
			} else {
				n2 = containers[len(containers)-1].Material().RefractiveIndex
			}
			break
		}
	}

	return &Comps{
		T:             i.T,
		Object:        i.Object,
		Point:         point,
		EyeVector:     eyeVector,
		NormalVector:  normalVector,
		ReflectVector: reflectVector,
		Inside:        inside,
		OverPoint:     overPoint,
		UnderPoint:    underPoint,
		N1:            n1,
		N2:            n2,
	}
}

// Schlick returns the reflectance of the intersection as an approximation
// the Fresnel law, as developed by Christophe Shlick.
func (c *Comps) Schlick() float64 {
	cos := c.EyeVector.Dot(c.NormalVector)

	if c.N1 > c.N2 {
		n := c.N1 / c.N2
		sin2t := n * n * (1 - (cos * cos))
		if sin2t > 1 {
			return 1
		}

		cos = math.Sqrt(1 - sin2t)
	}

	ratio := (c.N1 - c.N2) / (c.N1 + c.N2)
	r0 := ratio * ratio
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
