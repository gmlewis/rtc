package rtc

import "math"

// Cube creates a cube at the origin ranging from -1 to 1 on each axis.
// It implements the Object interface.
func Cube() *CubeT {
	return &CubeT{Shape{transform: M4Identity(), material: Material()}}
}

// CubeT represents a Cube.
type CubeT struct {
	Shape
}

var _ Object = &CubeT{}

// SetTransform sets the object's transform 4x4 matrix.
func (c *CubeT) SetTransform(m M4) Object {
	c.transform = m
	return c
}

// SetMaterial sets the object's material.
func (c *CubeT) SetMaterial(material MaterialT) Object {
	c.material = material
	return c
}

// SetParent sets the object's parent object.
func (c *CubeT) SetParent(parent Object) Object {
	c.parent = parent
	return c
}

// Bounds returns the minimum bounding box of the object in object
// (untransformed) space.
func (c *CubeT) Bounds() *BoundsT {
	return &BoundsT{
		Min: Point(-1, -1, -1),
		Max: Point(1, 1, 1),
	}
}

func checkAxis(origin, direction, min, max float64) (tmin float64, tmax float64) {
	tminNumerator := min - origin
	tmaxNumerator := max - origin

	if math.Abs(direction) >= epsilon {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
	} else {
		tmin = tminNumerator * math.Inf(1)
		tmax = tmaxNumerator * math.Inf(1)
	}

	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}
	return tmin, tmax
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (c *CubeT) LocalIntersect(ray RayT) []IntersectionT {
	xtmin, xtmax := checkAxis(ray.Origin.X(), ray.Direction.X(), -1, 1)
	ytmin, ytmax := checkAxis(ray.Origin.Y(), ray.Direction.Y(), -1, 1)
	ztmin, ztmax := checkAxis(ray.Origin.Z(), ray.Direction.Z(), -1, 1)

	tmin := math.Max(xtmin, math.Max(ytmin, ztmin))
	tmax := math.Min(xtmax, math.Min(ytmax, ztmax))

	if tmin > tmax {
		return nil
	}

	return []IntersectionT{Intersection(tmin, c), Intersection(tmax, c)}
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (c *CubeT) LocalNormalAt(objectPoint Tuple, hit *IntersectionT) Tuple {
	absX := math.Abs(objectPoint.X())
	absY := math.Abs(objectPoint.Y())
	maxc := math.Max(absX, math.Max(absY, math.Abs(objectPoint.Z())))

	if maxc == absX {
		return Vector(objectPoint.X(), 0, 0)
	}
	if maxc == absY {
		return Vector(0, objectPoint.Y(), 0)
	}
	return Vector(0, 0, objectPoint.Z())
}
