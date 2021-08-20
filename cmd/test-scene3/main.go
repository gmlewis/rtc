// test-scene3 implements the test scene in Chapter 14.
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

	w.Objects = []rtc.Object{floor, leftWall, rightWall, hexagon()}
	// w.Lights = []*rtc.PointLightT{rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1))}  // one light
	w.Lights = []*rtc.PointLightT{ // two lights
		rtc.PointLight(rtc.Point(-10, 10, -10), rtc.Color(1, 1, 1)),
		rtc.PointLight(rtc.Point(50, 50, -50), rtc.Color(0.1, 0.09, 0.08)),
	}
	return w
}

func hexagonCorner() rtc.Object {
	return rtc.Sphere().SetTransform(rtc.Translation(0, 0, -1).Mult(rtc.Scaling(0.25, 0.25, 0.25)))
}

func hexagonEdge() rtc.Object {
	edge := rtc.Cylinder()
	edge.Minimum = 0
	edge.Maximum = 1
	edge.SetTransform(rtc.Translation(0, 0, -1).Mult(rtc.RotationY(-math.Pi / 6)).Mult(rtc.RotationZ(-math.Pi / 2)).Mult(rtc.Scaling(0.25, 1, 0.25)))
	return edge
}

func hexagonSide() rtc.Object {
	return rtc.Group(hexagonCorner(), hexagonEdge())
}

func hexagon() rtc.Object {
	hex := rtc.Group()

	for n := 0; n < 6; n++ {
		side := hexagonSide()
		side.SetTransform(rtc.RotationY(float64(n) * math.Pi / 3))
		hex.AddChild(side)
	}

	hex.SetTransform(rtc.Translation(0, 0.5, 0))

	return hex
}
