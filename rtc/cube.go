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

func checkAxis(origin, direction float64) (tmin float64, tmax float64) {
	const infinity = 1e100
	tminNumerator := -1 - origin
	tmaxNumerator := 1 - origin

	if math.Abs(direction) >= epsilon {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
	} else {
		tmin = tminNumerator * infinity
		tmax = tmaxNumerator * infinity
	}

	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}
	return tmin, tmax
}

// LocalIntersect returns a slice of IntersectionT values where the
// transformed (object space) ray intersects the object.
func (c *CubeT) LocalIntersect(ray RayT) []IntersectionT {
	xtmin, xtmax := checkAxis(ray.Origin.X(), ray.Direction.X())
	ytmin, ytmax := checkAxis(ray.Origin.Y(), ray.Direction.Y())
	ztmin, ztmax := checkAxis(ray.Origin.Z(), ray.Direction.Z())

	tmin := math.Max(xtmin, math.Max(ytmin, ztmin))
	tmax := math.Min(xtmax, math.Min(ytmax, ztmax))

	if tmin > tmax {
		return nil
	}

	return []IntersectionT{Intersection(tmin, c), Intersection(tmax, c)}
}

// LocalNormalAt returns the normal vector at the given point of intersection
// (transformed to object space) with the object.
func (c *CubeT) LocalNormalAt(objectPoint Tuple) Tuple {
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
