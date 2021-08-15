// cannon is a simple test game recommended on page 38 of the book.
package main

import (
	"github.com/gmlewis/rtc/rtc"
)

type Projectile struct {
	Position *rtc.Tuple
	Velocity *rtc.Tuple
}

type Environment struct {
	Gravity *rtc.Tuple
	Wind    *rtc.Tuple
}

func main() {

}

func tick(env *Environment, proj *Projectile) *Projectile {
	pos := proj.Position.Add(proj.Velocity)
	vel := proj.Velocity.Add(env.Gravity).Add(env.Wind)
	return &Projectile{Position: pos, Velocity: vel}
}
