// cannon is a simple test game recommended on page 38 of the book.
package main

import (
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/gmlewis/rtc/rtc"
)

var (
	posx  = flag.Float64("px", 0, "Projectile position x")
	posy  = flag.Float64("py", 1, "Projectile position y")
	posz  = flag.Float64("pz", 0, "Projectile position z")
	vels  = flag.Float64("vs", 11.25, "Projectile velocity vector scale factor")
	velx  = flag.Float64("vx", 1, "Projectile velocity vector x")
	vely  = flag.Float64("vy", 1.8, "Projectile velocity vector y")
	velz  = flag.Float64("vz", 0, "Projectile velocity vector z")
	gravx = flag.Float64("gx", 0, "Environment gravity vector x")
	gravy = flag.Float64("gy", -0.1, "Environment gravity vector y")
	gravz = flag.Float64("gz", 0, "Environment gravity vector z")
	windx = flag.Float64("wx", -0.01, "Environment wind vector x")
	windy = flag.Float64("wy", 0, "Environment wind vector y")
	windz = flag.Float64("wz", 0, "Environment wind vector z")

	pngFile = flag.String("png", "cannon.png", "Output PNG file")
	ppmFile = flag.String("ppm", "cannon.ppm", "Output PPM file")
)

type Projectile struct {
	Position rtc.Tuple // Point
	Velocity rtc.Tuple // Vector
}

type Environment struct {
	Gravity rtc.Tuple // Vector
	Wind    rtc.Tuple // Vector
}

func main() {
	p := &Projectile{
		Position: rtc.Point(*posx, *posy, *posz),
		Velocity: rtc.Vector(*velx, *vely, *velz).Normalize().MultScalar(*vels),
	}
	e := &Environment{
		Gravity: rtc.Vector(*gravx, *gravy, *gravz),
		Wind:    rtc.Vector(*windx, *windy, *windz),
	}

	var maxx, maxy float64
	var xvals []int
	var yvals []int

	var ticks int
	for {
		if p.Position.X() > maxx {
			maxx = p.Position.X()
		}
		if p.Position.Y() > maxy {
			maxy = p.Position.Y()
		}
		xvals = append(xvals, int(math.Floor(0.5+p.Position.X())))
		yvals = append(yvals, int(math.Floor(0.5+p.Position.Y())))

		fmt.Printf("After tick #%v, position: (%0.2f,%0.2f)\n", ticks, p.Position.X(), p.Position.Y())
		if p.Position.Y() <= 0 {
			break
		}
		p = tick(e, p)
		ticks++
	}
	fmt.Printf("Projectile hit the ground after %v ticks.\n", ticks)

	var canvas *rtc.Canvas
	if *ppmFile != "" || *pngFile != "" {
		width := int(math.Floor(1.5 + maxx))
		height := int(math.Floor(1.5 + maxy))
		canvas = rtc.NewCanvas(width, height)

		red := rtc.Color(1, 0, 0)
		for i, xv := range xvals {
			canvas.WritePixel(xv, height-yvals[i]-1, red)
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

func tick(env *Environment, proj *Projectile) *Projectile {
	pos := proj.Position.Add(proj.Velocity)
	vel := proj.Velocity.Add(env.Gravity).Add(env.Wind)
	return &Projectile{Position: pos, Velocity: vel}
}
