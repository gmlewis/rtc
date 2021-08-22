package rtc

import (
	"math"
	"sync"
)

const (
	maxReflections = 4
)

// CameraT represents a camera.
type CameraT struct {
	HSize       int
	VSize       int
	FieldOfView float64
	Transform   M4
	PixelSize   float64
	HalfWidth   float64
	HalfHeight  float64
	NumWorkers  int

	cached       bool
	cachedInv    M4 // Inverse of Transform
	cachedOrigin Tuple
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
		NumWorkers:  18, // sweet spot on my machine
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
	if !c.cached {
		c.cachedInv = c.Transform.Inverse()
		c.cachedOrigin = c.cachedInv.MultTuple(Point(0, 0, 0))
		c.cached = true
	}
	pixel := c.cachedInv.MultTuple(Point(worldX, worldY, -1))
	direction := pixel.Sub(c.cachedOrigin).Normalize()

	return Ray(c.cachedOrigin, direction)
}

// Render renders the world with the camera and returns an image.
func (c *CameraT) Render(world *WorldT) *Canvas {
	canvas := NewCanvas(c.HSize, c.VSize)

	f := func(x, y int) {
		ray := c.RayForPixel(x, y)
		color := world.ColorAt(ray, maxReflections)
		canvas.WritePixel(x, y, color)
	}

	var wg sync.WaitGroup
	if c.NumWorkers > 1 {
		ch := make(chan struct{}, c.NumWorkers)
		origF := f
		f = func(x, y int) {
			wg.Add(1)
			ch <- struct{}{}
			go func(x, y int) {
				origF(x, y)
				wg.Done()
				<-ch
			}(x, y)
		}
	}

	for y := 0; y < c.VSize; y++ {
		for x := 0; x < c.HSize; x++ {
			f(x, y)
		}
	}

	if c.NumWorkers > 1 {
		wg.Wait()
	}

	return canvas
}
