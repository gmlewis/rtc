package rtc

import (
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
