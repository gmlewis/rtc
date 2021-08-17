// test-scene implements the test scene in Chapter 7.
package main

import (
	"flag"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/gmlewis/rtc/rtc"
)

var (
	xsize = flag.Int("xsize", 100, "X size")
	ysize = flag.Int("ysize", 50, "Y size")

	pngFile = flag.String("png", "test-scene.png", "Output PNG file")
	ppmFile = flag.String("ppm", "test-scene.ppm", "Output PPM file")
)

func main() {
	flag.Parse()

	world := genWorld()

	camera := rtc.Camera(*xsize, *ysize, math.Pi/3)
	camera.Transform = rtc.ViewTransform(
		rtc.Point(0, 1.5, -5),
		rtc.Point(0, 1, 0),
		rtc.Vector(0, 1, 0))
	canvas := camera.Render(world)

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

	middle := rtc.Sphere()
	middle.SetTransform(rtc.Translation(-0.5, 1, 0.5))
	middle.Material().Color = rtc.Color(0.1, 1, 0.5)
	middle.Material().Diffuse = 0.7
	middle.Material().Specular = 0.3

	right := rtc.Sphere()
	right.SetTransform(rtc.M4Identity().Scale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5))
	right.Material().Color = rtc.Color(0.5, 1, 0.1)
	right.Material().Diffuse = 0.7
	right.Material().Specular = 0.3

	left := rtc.Sphere()
	left.SetTransform(rtc.M4Identity().Scale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75))
	left.Material().Color = rtc.Color(1, 0.8, 0.1)
	left.Material().Diffuse = 0.7
	left.Material().Specular = 0.3

	w.Objects = []rtc.Object{floor, leftWall, rightWall, middle, right, left}
	// w.Lights = []*rtc.PointLightT{rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1))}  // one light
	w.Lights = []*rtc.PointLightT{ // two lights
		rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1)),
		rtc.PointLight(rtc.Point(50, 50, -50), rtc.Color(1, 0.9, 0.8)),
	}
	return w
}
