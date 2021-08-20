// cast-test is the "Putting it together" section of Chapter 5.
// It renders the shadow of a sphere on a wall.
package main

import (
	"flag"
	"log"

	"github.com/gmlewis/rtc/rtc"
)

var (
	size     = flag.Int("size", 100, "Canvas X and Y size")
	wallSize = flag.Float64("wallSize", 7.0, "Wall size")
	wallZ    = flag.Float64("wallZ", 10.0, "Wall Z")

	pngFile = flag.String("png", "cast-test.png", "Output PNG file")
	ppmFile = flag.String("ppm", "cast-test.ppm", "Output PPM file")
)

func main() {
	flag.Parse()

	rayOrigin := rtc.Point(0, 0, -5)
	pixelSize := *wallSize / float64(*size)
	half := *wallSize / 2
	canvas := rtc.NewCanvas(*size, *size)
	shape := rtc.Sphere()
	shape.Material().Color = rtc.Color(1, 0.2, 1)

	lightPosition := rtc.Point(-10, 10, -10)
	lightColor := rtc.Color(1, 1, 1)
	light := rtc.PointLight(lightPosition, lightColor)

	// shrink it along the y axis​
	// shape.SetTransform(rtc.Scaling(1, 0.5, 1))

	// shrink it along the x axis​
	// shape.SetTransform(rtc.Scaling(0.5, 1, 1))

	// shrink it, and rotate it!​
	// shape.SetTransform(rtc.RotationZ(math.Pi / 4).Mult(rtc.Scaling(0.5, 1, 1)))

	// shrink it, and skew it!​
	// shape.SetTransform(rtc.Shearing(1, 0, 0, 0, 0, 0).Mult(rtc.Scaling(0.5, 1, 1)))

	for y := 0; y < *size; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < *size; x++ {
			worldX := -half + pixelSize*float64(x)
			position := rtc.Point(worldX, worldY, *wallZ)
			r := rtc.Ray(rayOrigin, position.Sub(rayOrigin).Normalize())
			xs := rtc.Intersect(shape, r)
			if hit := rtc.Hit(xs); hit != nil {
				point := r.Position(hit.T)
				normal := hit.NormalAt(point)
				eye := r.Direction.Negate()
				color := rtc.Lighting(hit.Object.Material(), hit.Object, light, point, eye, normal, false)
				canvas.WritePixel(x, y, color)
			}
		}
	}

	if *pngFile != "" {
		if err := canvas.WritePNGFile(*pngFile); err != nil {
			log.Fatal(err)
		}
	}

	if *ppmFile != "" {
		if err := canvas.WritePPMFile(*ppmFile); err != nil {
			log.Fatal(err)
		}
	}
}
