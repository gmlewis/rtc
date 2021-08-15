package rtc

import (
	"strings"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	got := c.Bounds()

	if got.Min.X != 0 {
		t.Errorf("NewCanvas Min.X = %v, want 0", got.Min.X)
	}
	if got.Min.Y != 0 {
		t.Errorf("NewCanvas Min.Y = %v, want 0", got.Min.Y)
	}
	if want := 10; got.Max.X != want {
		t.Errorf("NewCanvas Max.X = %v, want %v", got.Max.X, want)
	}
	if want := 20; got.Max.Y != want {
		t.Errorf("NewCanvas Max.Y = %v, want %v", got.Max.Y, want)
	}

	black := Color(0, 0, 0)
	for y := got.Min.Y; y < got.Max.Y; y++ {
		for x := got.Min.X; x < got.Max.X; x++ {
			got := c.PixelAt(x, y)
			if !got.Equal(black) {
				t.Errorf("pixel (%v,%v) = %v, want %v", x, y, got, black)
			}
		}
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	c := NewCanvas(10, 20)
	red := Color(1, 0, 0)
	c.WritePixel(2, 3, red)
	got := c.PixelAt(2, 3)
	if !got.Equal(red) {
		t.Errorf("PixelAt(2,3) = %v, want %v", got, red)
	}
}

func TestCanvas_ToPPM_Header(t *testing.T) {
	c := NewCanvas(5, 3)

	ppm := c.ToPPM()

	lines := strings.Split(ppm, "\n")
	want := `P3
5 3
255`
	got := strings.Join(lines[0:3], "\n")
	if got != want {
		t.Errorf("ToPPM header =\n%v\nwant:\n%v", got, want)
	}
}

func TestCanvas_ToPPM_Pixel_Data(t *testing.T) {
	c := NewCanvas(5, 3)
	c1 := Color(1.5, 0, 0)
	c2 := Color(0, 0.5, 0)
	c3 := Color(-0.5, 0, 1)

	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	ppm := c.ToPPM()

	lines := strings.Split(ppm, "\n")
	want := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255`
	got := strings.Join(lines[4:7], "\n")
	if got != want {
		t.Errorf("ToPPM pixel data =\n%v\nwant:\n%v", got, want)
	}
}
