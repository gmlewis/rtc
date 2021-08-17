package rtc

import "math"

// CameraT represents a camera.
type CameraT struct {
	HSize       int
	VSize       int
	FieldOfView float64
	Transform   M4
	PixelSize   float64
	HalfWidth   float64
	HalfHeight  float64
}

// Camera creates a new CameraT with the provided canvas size and
// field of view (in radians).
func Camera(hsize, vsize int, fov float64) *CameraT {
	halfView := math.Tan(fov / 2)
	aspect := float64(hsize) / float64(vsize)

	c := &CameraT{
		HSize:       hsize,
		VSize:       vsize,
		FieldOfView: fov,
		Transform:   M4Identity(),
	}

	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}

	c.PixelSize = c.HalfWidth * 2 / float64(c.HSize)

	return c
}

// RayForPixel returns a ray for the camera at the given pixel.
func (c *CameraT) RayForPixel(px, py int) RayT {
	// The offset from the edge of the canvas to the pixel's center
	xoffset := (float64(px) + 0.5) * c.PixelSize
	yoffset := (float64(py) + 0.5) * c.PixelSize

	// The untransformed coordinates of the pixel in world space.
	// The camera looks toward -Z, so +X is to the *left*.
	worldX := c.HalfWidth - xoffset
	worldY := c.HalfHeight - yoffset

	// Using the camera matrix, transform the canvas point and the origin
	// and then compute the ray's direction vector.
	// The canvas is at Z = -1.
	inv := c.Transform.Inverse()
	pixel := inv.MultTuple(Point(worldX, worldY, -1))
	origin := inv.MultTuple(Point(0, 0, 0))
	direction := pixel.Sub(origin).Normalize()

	return Ray(origin, direction)
}
