package rtc

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

// Canvas represents an image canvas and implements the image.Image interface.
type Canvas struct {
	width  int
	height int
	pixels []Tuple
}

var _ image.Image = &Canvas{}

// NewCanvas returns a new canvas with the given dimensions.
func NewCanvas(width, height int) *Canvas {
	return &Canvas{
		width:  width,
		height: height,
		pixels: make([]Tuple, width*height),
	}
}

func clamp(v float64) uint16 {
	clamped := math.Max(math.Min(v, 1), 0)
	return uint16(math.Floor(65535 * clamped))
}

// At returns the color at the provides location in the canvas.
func (c *Canvas) At(x, y int) color.Color {
	if x < 0 || y < 0 || x >= c.width || y >= c.height {
		return color.Black
	}
	idx := y*c.width + x
	pixel := &c.pixels[idx]
	r := clamp(pixel.Red())
	g := clamp(pixel.Green())
	b := clamp(pixel.Blue())
	a := clamp(pixel.Alpha())
	return color.NRGBA64{R: r, G: g, B: b, A: a}
}

// Bounds returns the bounding box of the canvas.
func (c *Canvas) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.width, c.height)
}

// ColorModel returns the Image's color model.
func (c *Canvas) ColorModel() color.Model {
	return color.NRGBA64Model
}

// WritePixel writes a pixel to the canvas.
func (c *Canvas) WritePixel(x, y int, color *Tuple) {
	if x < 0 || y < 0 || x >= c.width || y >= c.height {
		return
	}
	idx := y*c.width + x
	c.pixels[idx] = *color
}

// PixelAt returns the color (Tuple) at the given pixel.
func (c *Canvas) PixelAt(x, y int) *Tuple {
	if x < 0 || y < 0 || x >= c.width || y >= c.height {
		return &Tuple{}
	}
	idx := y*c.width + x
	return &c.pixels[idx]
}

// ToPPM returns a string PPM representation of the canvas.
func (c *Canvas) ToPPM() string {
	header := fmt.Sprintf("P3\n%v %v\n255\n", c.width, c.height)
	return header
}
