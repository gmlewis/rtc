// test-yaml reads a YAML scene description as demonstrated in Appendix 1.
package main

import (
	"flag"
	"log"
	"math"

	"github.com/gmlewis/rtc/rtc"
)

var (
	xsize = flag.Int("xsize", 1280, "X size")
	ysize = flag.Int("ysize", 1024, "Y size")

	pngFile = flag.String("png", "test-yaml.png", "Output PNG file")
	ppmFile = flag.String("ppm", "test-yaml.ppm", "Output PPM file")
)

func main() {
	flag.Parse()

	world := rtc.World()

	for _, arg := range flag.Args() {
		yaml, err := rtc.ParseYAMLFile(arg)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("yaml=%#v", yaml)
	}

	camera := rtc.Camera(*xsize, *ysize, math.Pi/3)
	camera.Transform = rtc.ViewTransform(
		rtc.Point(0, 1.5, -5),
		rtc.Point(0, 1, 0),
		rtc.Vector(0, 1, 0))
	canvas := camera.Render(world)

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
