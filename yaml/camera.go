package yaml

import (
	"math"

	"github.com/gmlewis/rtc/rtc"
)

// Camera returns the YAML camera definition.
// xsize, ysize, and field-of-view are all optional and may be nil.
func (y *YAMLFile) Camera(xsize, ysize *int, fov *float64) *rtc.CameraT {
	for _, item := range y.Items {
		if item.Add == nil {
			continue
		}
		if *item.Add == "camera" {
			width, height := 1280, 1024
			if item.Width != nil {
				width = *item.Width
			}
			if xsize != nil {
				width = *xsize
			}
			if item.Height != nil {
				height = *item.Height
			}
			if xsize != nil {
				height = *ysize
			}
			finalFOV := math.Pi / 3
			if item.FOV != nil {
				finalFOV = *item.FOV
			}
			if fov != nil {
				finalFOV = *fov
			}
			camera := rtc.Camera(width, height, finalFOV)
			from := rtc.Vector(item.From[0], item.From[1], item.From[2])
			to := rtc.Vector(item.To[0], item.To[1], item.To[2])
			up := rtc.Vector(item.Up[0], item.Up[1], item.Up[2])
			camera.Transform = rtc.ViewTransform(from, to, up)
			return camera
		}
	}
	return nil
}
