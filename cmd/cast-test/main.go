// cast-test is the "Putting it together" section of Chapter 5.
// It renders the shadow of a sphere on a wall.
package main

import (
	"flag"
	"image/png"
	"io/ioutil"
	"log"
	"os"

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
	rayOrigin := rtc.Point(0, 0, -5)
	pixelSize := *wallSize / float64(*size)
	half := *wallSize / 2
	canvas := rtc.NewCanvas(*size, *size)
	color := rtc.Color(1, 0, 0)
	shape := rtc.Sphere()

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
			xs := shape.Intersect(r)
			if hit := rtc.Hit(xs); hit != nil {
				canvas.WritePixel(x, y, color)
			}
		}
	}

	if *pngFile != "" {
		f, err := os.Create(*pngFile)
		if err != nil {
			log.Fatal(err)
		}

		if err := png.Encode(f, canvas); err != nil {
			f.Close()
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	if *ppmFile != "" {
		ppm := canvas.ToPPM()
		if err := ioutil.WriteFile(*ppmFile, []byte(ppm), 0644); err != nil {
			log.Fatal(err)
		}
	}
}
