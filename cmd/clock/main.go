// clock renders the ticks on a clock face using transformations.
package main

import (
	"flag"
	"log"
	"math"

	"github.com/gmlewis/rtc/rtc"
)

var (
	size = flag.Int("size", 800, "X and Y size")

	pngFile = flag.String("png", "clock.png", "Output PNG file")
	ppmFile = flag.String("ppm", "clock.ppm", "Output PPM file")
)

func main() {
	flag.Parse()

	canvas := rtc.NewCanvas(*size, *size)

	white := rtc.Color(1, 1, 1)
	r := 0.35 * float64(*size)
	center := 0.5 * float64(*size)

	for i := 0; i < 12; i++ {
		angle := 2.0 * math.Pi * float64(i) / 12
		pt := rtc.M4Identity().Scale(r, 1, r).RotateY(angle).Translate(center, 0, center).MultTuple(rtc.Point(0, 0, 1))
		log.Printf("i=%v, pt=%v", i, pt)
		canvas.WritePixel(int(pt.X()), int(pt.Z()), white)
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
