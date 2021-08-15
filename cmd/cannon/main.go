// cannon is a simple test game recommended on page 38 of the book.
package main

import (
	"flag"
	"fmt"

	"github.com/gmlewis/rtc/rtc"
)

var (
	posx  = flag.Float64("px", 0, "Projectile position x")
	posy  = flag.Float64("py", 1, "Projectile position y")
	posz  = flag.Float64("pz", 0, "Projectile position z")
	velx  = flag.Float64("vx", 1, "Projectile velocity vector x")
	vely  = flag.Float64("vy", 1, "Projectile velocity vector y")
	velz  = flag.Float64("vz", 0, "Projectile velocity vector z")
	gravx = flag.Float64("gx", 0, "Environment gravity vector x")
	gravy = flag.Float64("gy", -0.1, "Environment gravity vector y")
	gravz = flag.Float64("gz", 0, "Environment gravity vector z")
	windx = flag.Float64("wx", -0.01, "Environment wind vector x")
	windy = flag.Float64("wy", 0, "Environment wind vector y")
	windz = flag.Float64("wz", 0, "Environment wind vector z")
)

type Projectile struct {
	Position *rtc.Tuple // Point
	Velocity *rtc.Tuple // Vector
}

type Environment struct {
	Gravity *rtc.Tuple // Vector
	Wind    *rtc.Tuple // Vector
}

func main() {
	p := &Projectile{
		Position: rtc.Point(*posx, *posy, *posz),
		Velocity: rtc.Vector(*velx, *vely, *velz),
	}
	e := &Environment{
		Gravity: rtc.Vector(*gravx, *gravy, *gravz),
		Wind:    rtc.Vector(*windx, *windy, *windz),
	}

	var ticks int
	for {
		fmt.Printf("After tick #%v, position: (%0.2f,%0.2f,%0.2f)\n", ticks, p.Position.X(), p.Position.Y(), p.Position.Z())
		if p.Position.Y() <= 0 {
			break
		}
		p = tick(e, p)
		ticks++
	}
	fmt.Printf("Projectile hit the ground after %v ticks.\n", ticks)
}

func tick(env *Environment, proj *Projectile) *Projectile {
	pos := proj.Position.Add(proj.Velocity)
	vel := proj.Velocity.Add(env.Gravity).Add(env.Wind)
	return &Projectile{Position: pos, Velocity: vel}
}
