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
func (c *CameraT) RayForPixel(x, y int) RayT {
	return RayT{}
}
