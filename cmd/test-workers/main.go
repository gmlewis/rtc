// test-workers benchmarks the use of goroutines to render the
// scene from Chapter 7.
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gmlewis/rtc/rtc"
)

var (
	xsize = flag.Int("xsize", 1280, "X size")
	ysize = flag.Int("ysize", 1024, "Y size")

	pngFile = flag.String("png", "test-workers.png", "Output PNG file")
	ppmFile = flag.String("ppm", "test-workers.ppm", "Output PPM file")
)

func main() {
	flag.Parse()

	world := genWorld()

	camera := rtc.Camera(*xsize, *ysize, math.Pi/3)
	camera.Transform = rtc.ViewTransform(
		rtc.Point(0, 1.5, -5),
		rtc.Point(0, 1, 0),
		rtc.Vector(0, 1, 0))

	workers := []int{
		// 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19, // sweet spot range on my machine.
		// 10, 20, 30, 40, 50, 60, 70, 80, 90,
		// 100, 200, 300, 400, 500, 600, 700, 800, 900,
		// 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000,
		// 10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000,
		// 100000, 200000, 300000, 400000, 500000, 600000, 700000, 800000, 900000,
	}

	fmt.Println("Workers,Seconds")
	for _, num := range workers {
		before := time.Now().Local()
		benchmark(num, camera, world)
		delta := time.Since(before)
		fmt.Printf("%v,%v\n", num, delta.Seconds())
	}
}

func benchmark(numWorkers int, camera *rtc.CameraT, world *rtc.WorldT) {
	camera.NumWorkers = numWorkers
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

	middle := rtc.Sphere()
	middle.SetTransform(rtc.Translation(-0.5, 1, 0.5))
	middle.GetMaterial().Color = rtc.Color(0.1, 1, 0.5)
	middle.GetMaterial().Diffuse = 0.7
	middle.GetMaterial().Specular = 0.3

	right := rtc.Sphere()
	right.SetTransform(rtc.M4Identity().Scale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5))
	right.GetMaterial().Color = rtc.Color(0.5, 1, 0.1)
	right.GetMaterial().Diffuse = 0.7
	right.GetMaterial().Specular = 0.3

	left := rtc.Sphere()
	left.SetTransform(rtc.M4Identity().Scale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75))
	left.GetMaterial().Color = rtc.Color(1, 0.8, 0.1)
	left.GetMaterial().Diffuse = 0.7
	left.GetMaterial().Specular = 0.3

	w.Objects = []rtc.Object{floor, leftWall, rightWall, middle, right, left}
	// w.Lights = []*rtc.PointLightT{rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1))}  // one light
	w.Lights = []*rtc.PointLightT{ // two lights
		rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1)),
		rtc.PointLight(rtc.Point(50, 50, -50), rtc.Color(1, 0.9, 0.8)),
	}
	return w
}
