package yaml

import "github.com/gmlewis/rtc/rtc"

// Camera returns the YAML camera definition.
func (y *YAMLFile) Camera() *rtc.CameraT {
	for _, item := range y.Items {
		if item.Add == nil || item.Width == nil || item.Height == nil || item.FOV == nil {
			continue
		}
		if *item.Add == "camera" {
			camera := rtc.Camera(*item.Width, *item.Height, *item.FOV)
			from := rtc.Vector(item.From[0], item.From[1], item.From[2])
			to := rtc.Vector(item.To[0], item.To[1], item.To[2])
			up := rtc.Vector(item.Up[0], item.Up[1], item.Up[2]).Normalize()
			forward := from.Sub(to).Normalize()
			right := up.Cross(forward)
			camera.Transform = rtc.M4{
				rtc.Tuple{right.X(), right.Y(), right.Z(), 0},
				rtc.Tuple{up.X(), up.Y(), up.Z(), 0},
				rtc.Tuple{forward.X(), forward.Y(), forward.Z(), 0},
				rtc.Tuple{from.X(), from.Y(), from.Z(), 0},
			}
			return camera
		}
	}
	return nil
}
