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

// Bounds returns an empty bounding box.
func Bounds() *BoundsT {
	return &BoundsT{
		Min: Point(math.Inf(1), math.Inf(1), math.Inf(1)),
		Max: Point(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	}
}

// UpdateBounds updates the bounding box with the provided point.
func (b *BoundsT) UpdateBounds(p Tuple) {
	if p.X() < b.Min.X() {
		b.Min[0] = p.X()
	}
	if p.Y() < b.Min.Y() {
		b.Min[1] = p.Y()
	}
	if p.Z() < b.Min.Z() {
		b.Min[2] = p.Z()
	}
	if p.X() > b.Max.X() {
		b.Max[0] = p.X()
	}
	if p.Y() > b.Max.Y() {
		b.Max[1] = p.Y()
	}
	if p.Z() > b.Max.Z() {
		b.Max[2] = p.Z()
	}
}
