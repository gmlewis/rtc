package rtc

import "math"

// BoundsT represents a minimum bounding box of an object.
type BoundsT struct {
	Min Tuple
	Max Tuple
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (b *BoundsT) LocalIntersect(ray RayT, group *GroupT) []IntersectionT {
	xtmin, xtmax := checkAxis(ray.Origin.X(), ray.Direction.X(), b.Min.X(), b.Max.X())
	ytmin, ytmax := checkAxis(ray.Origin.Y(), ray.Direction.Y(), b.Min.Y(), b.Max.Y())
	ztmin, ztmax := checkAxis(ray.Origin.Z(), ray.Direction.Z(), b.Min.Z(), b.Max.Z())

	tmin := math.Max(xtmin, math.Max(ytmin, ztmin))
	tmax := math.Min(xtmax, math.Min(ytmax, ztmax))

	if tmin > tmax {
		return nil
	}

	return []IntersectionT{Intersection(tmin, group), Intersection(tmax, group)}
}
