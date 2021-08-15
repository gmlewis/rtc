package rtc

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"strings"
)

const (
	maxPPMLineLen = 70
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

func clamp8(v float64) uint8 {
	clamped := math.Max(math.Min(v, 1), 0)
	return uint8(math.Floor(0.5 + 255*clamped))
}

func clamp16(v float64) uint16 {
	clamped := math.Max(math.Min(v, 1), 0)
	return uint16(math.Floor(0.5 + 65535*clamped))
}

// At returns the color at the provides location in the canvas.
func (c *Canvas) At(x, y int) color.Color {
	if x < 0 || y < 0 || x >= c.width || y >= c.height {
		return color.Black
	}
	idx := y*c.width + x
	pixel := &c.pixels[idx]
	r := clamp16(pixel.Red())
	g := clamp16(pixel.Green())
	b := clamp16(pixel.Blue())
	a := clamp16(pixel.Alpha())
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
	var lines []string
	for y := 0; y < c.height; y++ {
		var line string
		for x := 0; x < c.width; x++ {
			pixel := c.PixelAt(x, y)
			p := fmt.Sprintf("%v %v %v", clamp8(pixel.Red()), clamp8(pixel.Green()), clamp8(pixel.Blue()))
			if len(line)+1+len(p) > maxPPMLineLen {
				lines = append(lines, line)
				line = p
			} else if len(line) == 0 {
				line = p
			} else {
				line += " " + p
			}
		}
		lines = append(lines, line)
	}
	return fmt.Sprintf("%v\n%v\n", header, strings.Join(lines, "\n"))
}
