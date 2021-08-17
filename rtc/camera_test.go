package rtc

import (
	"math"
	"testing"
)

func TestCamera(t *testing.T) {
	tests := []struct {
		name          string
		hsize         int
		vsize         int
		fov           float64
		wantPixelSize float64
	}{
		{
			name:          "Constructing a camera",
			hsize:         160,
			vsize:         120,
			fov:           math.Pi / 2,
			wantPixelSize: 0.0125,
		},
		{
			name:          "The pixel size for a horizontal canvas",
			hsize:         200,
			vsize:         125,
			fov:           math.Pi / 2,
			wantPixelSize: 0.01,
		},
		{
			name:          "The pixel size for a vertical canvas",
			hsize:         125,
			vsize:         200,
			fov:           math.Pi / 2,
			wantPixelSize: 0.01,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Camera(tt.hsize, tt.vsize, tt.fov)

			if got.HSize != tt.hsize {
				t.Errorf("HSize = %v, want %v", got.HSize, tt.hsize)
			}

			if got.VSize != tt.vsize {
				t.Errorf("VSize = %v, want %v", got.VSize, tt.vsize)
			}

			if got.FieldOfView != tt.fov {
				t.Errorf("FieldOfView = %v, want %v", got.FieldOfView, tt.fov)
			}

			if got.PixelSize != tt.wantPixelSize {
				t.Errorf("Camera.PixelSize = %v, want %v", got.PixelSize, tt.wantPixelSize)
			}
		})
	}
}

func TestCameraT_RayForPixel(t *testing.T) {
	sq2 := math.Sqrt(2) / 2

	tests := []struct {
		name          string
		hsize         int
		vsize         int
		fov           float64
		transform     M4
		x             int
		y             int
		wantOrigin    Tuple
		wantDirection Tuple
	}{
		{
			name:          "Constructing a ray through the center of the canvas",
			hsize:         201,
			vsize:         101,
			fov:           math.Pi / 2,
			x:             100,
			y:             50,
			wantOrigin:    Point(0, 0, 0),
			wantDirection: Vector(0, 0, -1),
		},
		{
			name:          "Constructing a ray through a corner of the canvas",
			hsize:         201,
			vsize:         101,
			fov:           math.Pi / 2,
			x:             0,
			y:             0,
			wantOrigin:    Point(0, 0, 0),
			wantDirection: Vector(0.66519, 0.33259, -0.66851),
		},
		{
			name:          "Constructing a ray when the camera is transformed",
			hsize:         201,
			vsize:         101,
			fov:           math.Pi / 2,
			transform:     RotationY(math.Pi / 4).Mult(Translation(0, -2, 5)),
			x:             100,
			y:             50,
			wantOrigin:    Point(0, 2, -5),
			wantDirection: Vector(sq2, 0, -sq2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Camera(tt.hsize, tt.vsize, tt.fov)
			r := c.RayForPixel(tt.x, tt.y)

			if r.Origin != tt.wantOrigin {
				t.Errorf("Origin = %v, want %v", r.Origin, tt.wantOrigin)
			}

			if r.Direction != tt.wantDirection {
				t.Errorf("Direction = %v, want %v", r.Direction, tt.wantDirection)
			}
		})
	}
}
