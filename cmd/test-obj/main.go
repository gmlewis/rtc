// test-obj renders one or more Wavefront OBJ files.
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

	autoFit    = flag.Bool("f", true, "Auto-fit on load")
	scale      = flag.Float64("s", 1, "Scale object factor")
	yTranslate = flag.Float64("ty", 0, "Y translate object")
	xRotate    = flag.Float64("rx", 0, "X rotate object (in degrees)")
	yRotate    = flag.Float64("ry", 180, "Y rotate object (in degrees)")

	pngFile = flag.String("png", "test-obj.png", "Output PNG file")
	ppmFile = flag.String("ppm", "test-obj.ppm", "Output PPM file")
)

func main() {
	flag.Parse()

	world := genWorld()

	for _, arg := range flag.Args() {
		obj, err := rtc.ParseObjFile(arg)
		if err != nil {
			log.Fatal(err)
		}

		toRad := func(deg float64) float64 {
			return deg * math.Pi / 180
		}

		g := obj.ToGroup()
		b := g.Bounds()
		log.Printf("Processed file %q, %v lines ignored. Bounds: %v", arg, obj.IgnoredLines, b)

		if *autoFit {
			tx := -0.5 * (b.Min.X() + b.Max.X())
			ty := -b.Min.Y()
			tz := -0.5 * (b.Min.Z() + b.Max.Z())
			maxDim := b.Max.X() - b.Min.X()
			if v := b.Max.Y() - b.Min.Y(); v > maxDim {
				maxDim = v
			}
			if v := b.Max.Z() - b.Min.Z(); v > maxDim {
				maxDim = v
			}
			sf := 3.0 / maxDim
			log.Printf("auto-fit: translate=(%v,%v,%v), scale=%v", tx, ty, tz, sf)
			g.SetTransform(rtc.M4Identity().Translate(tx, ty, tz).Scale(sf, sf, sf))
		}

		xfm := g.Transform().Translate(0, *yTranslate, 0).Scale(*scale, *scale, *scale).RotateY(toRad(*yRotate)).RotateX(toRad(*xRotate))
		g.SetTransform(xfm)

		world.Objects = append(world.Objects, g)
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
	floor.Material().Color = rtc.Color(1, 0.9, 0.9)
	floor.Material().Specular = 0

	leftWall := rtc.Plane()
	leftWall.SetTransform(rtc.M4Identity().RotateX(math.Pi/2).RotateY(-math.Pi/4).Translate(0, 0, 5))
	leftWall.SetMaterial(*floor.Material())

	rightWall := rtc.Plane()
	rightWall.SetTransform(rtc.M4Identity().RotateX(math.Pi/2).RotateY(math.Pi/4).Translate(0, 0, 5))
	rightWall.SetMaterial(*floor.Material())

	w.Objects = []rtc.Object{floor, leftWall, rightWall}
	// w.Lights = []*rtc.PointLightT{rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1))}  // one light
	w.Lights = []*rtc.PointLightT{ // two lights
		rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1)),
		rtc.PointLight(rtc.Point(50, 50, -50), rtc.Color(0.1, 0.09, 0.08)),
	}
	return w
}
