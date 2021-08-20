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

	scale      = flag.Float64("s", 1, "Scale object factor")
	yTranslate = flag.Float64("y", 0, "Y translate object")
	yRotate    = flag.Float64("r", 0, "Y rotate object (in degrees)")

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

		g := obj.ToGroup()
		g.SetTransform(rtc.RotationY(*yRotate * math.Pi / 180).Mult(rtc.Scaling(*scale, *scale, *scale).Mult(rtc.Translation(0, *yTranslate, 0))))
		log.Printf("%q bounds: %v", arg, g.Bounds())
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
