// test-irmf reads one or more IRMF objects and renders them.  See irmf.io.
package main

import (
	"flag"
	"log"
	"math"

	"github.com/gmlewis/rtc/irmf"
	"github.com/gmlewis/rtc/rtc"
)

var (
	xsize = flag.Int("xsize", 128, "X size")
	ysize = flag.Int("ysize", 102, "Y size")

	autoFit    = flag.Bool("f", true, "Auto-fit on load")
	scale      = flag.Float64("s", 1, "Scale object factor")
	yTranslate = flag.Float64("ty", 0, "Y translate object")
	xRotate    = flag.Float64("rx", 0, "X rotate object (in degrees)")
	yRotate    = flag.Float64("ry", 180, "Y rotate object (in degrees)")

	pngFile = flag.String("png", "test-irmf.png", "Output PNG file")
	ppmFile = flag.String("ppm", "test-irmf.ppm", "Output PPM file")
)

func main() {
	flag.Parse()

	world := genWorld()

	for _, arg := range flag.Args() {
		obj, err := irmf.ParseFile(arg)
		if err != nil {
			log.Fatal(err)
		}

		toRad := func(deg float64) float64 {
			return deg * math.Pi / 180
		}

		b := obj.Bounds()
		log.Printf("Processed file %q. Bounds: %v", arg, b)

		if *autoFit {
			tx := -0.5 * (b.Min.X() + b.Max.X())
			ty := -0.5 * (b.Min.Y() + b.Max.Y())
			tz := -0.5 * (b.Min.Z() + b.Max.Z())
			maxDim := b.Max.X() - b.Min.X()
			if v := b.Max.Y() - b.Min.Y(); v > maxDim {
				maxDim = v
			}
			if v := b.Max.Z() - b.Min.Z(); v > maxDim {
				maxDim = v
			}
			sf := *scale * 3.0 / maxDim
			log.Printf("auto-fit: center=(%v,%v,%v), scale=%v", tx, ty, tz, sf)
			xfm := rtc.M4Identity().Translate(tx, ty, tz).RotateY(toRad(*yRotate)).RotateX(toRad(*xRotate)).Translate(0, *yTranslate-b.Min.Y(), 0).Scale(sf, sf, sf)
			obj.SetTransform(xfm)
		} else {
			xfm := rtc.M4Identity().RotateY(toRad(*yRotate)).RotateX(toRad(*xRotate)).Translate(0, *yTranslate, 0).Scale(*scale, *scale, *scale)
			obj.SetTransform(xfm)
		}

		world.Objects = append(world.Objects, obj)
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

func genWorld() *rtc.WorldT {
	w := rtc.World()

	floor := rtc.Plane()
	floor.GetMaterial().Color = rtc.Color(1, 0.9, 0.9)
	floor.GetMaterial().Specular = 0

	leftWall := rtc.Plane()
	leftWall.SetTransform(rtc.M4Identity().RotateX(math.Pi/2).RotateY(-math.Pi/4).Translate(0, 0, 5))
	leftWall.SetMaterial(*floor.GetMaterial())

	rightWall := rtc.Plane()
	rightWall.SetTransform(rtc.M4Identity().RotateX(math.Pi/2).RotateY(math.Pi/4).Translate(0, 0, 5))
	rightWall.SetMaterial(*floor.GetMaterial())

	w.Objects = []rtc.Object{floor, leftWall, rightWall}
	w.Lights = []*rtc.PointLightT{ // two lights
		rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1)),
		rtc.PointLight(rtc.Point(50, 50, -50), rtc.Color(0.1, 0.09, 0.08)),
	}
	return w
}
