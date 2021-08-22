package rtc

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	sq2 := math.Sqrt2 / 2

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
			transform:     M4Identity(),
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
			transform:     M4Identity(),
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
			c.Transform = tt.transform
			r := c.RayForPixel(tt.x, tt.y)

			if !r.Origin.Equal(tt.wantOrigin) {
				t.Errorf("Origin = %v, want %v", r.Origin, tt.wantOrigin)
			}

			if !r.Direction.Equal(tt.wantDirection) {
				t.Errorf("Direction = %v, want %v", r.Direction, tt.wantDirection)
			}
		})
	}
}

func TestViewTransform(t *testing.T) {
	tests := []struct {
		name string
		from Tuple
		to   Tuple
		up   Tuple
		want M4
	}{
		{
			name: "The transformation matrix for the default orientation",
			from: Point(0, 0, 0),
			to:   Point(0, 0, -1),
			up:   Vector(0, 1, 0),
			want: M4Identity(),
		},
		{
			name: "A view transformation matrix looking in positive z direction",
			from: Point(0, 0, 0),
			to:   Point(0, 0, 1),
			up:   Vector(0, 1, 0),
			want: Scaling(-1, 1, -1),
		},
		{
			name: "The view transformation moves the world",
			from: Point(0, 0, 8),
			to:   Point(0, 0, 0),
			up:   Vector(0, 1, 0),
			want: Translation(0, 0, -8),
		},
		{
			name: "An arbitrary view transformation",
			from: Point(1, 3, 2),
			to:   Point(4, -2, 8),
			up:   Vector(1, 1, 0),
			want: M4{
				Tuple{-0.50709, 0.50709, 0.67612, -2.36643},
				Tuple{0.76772, 0.60609, 0.12122, -2.82843},
				Tuple{-0.35857, 0.59761, -0.71714, 0.00000},
				Tuple{0.00000, 0.00000, 0.00000, 1.00000},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ViewTransform(tt.from, tt.to, tt.up); !cmp.Equal(got, tt.want) {
				t.Errorf("ViewGetTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCameraT_Render(t *testing.T) {
	w := DefaultWorld()
	c := Camera(11, 11, math.Pi/2)
	from := Point(0, 0, -5)
	to := Point(0, 0, 0)
	up := Point(0, 1, 0)
	c.Transform = ViewTransform(from, to, up)
	canvas := c.Render(w)

	if got, want := canvas.PixelAt(5, 5), Color(0.38066, 0.47583, 0.2855); !got.Equal(want) {
		t.Errorf("canvas.PixelAt(5,5) = %v, want %v", got, want)
	}
}
