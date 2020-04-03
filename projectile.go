package main

import (
	"fmt"
	"github.com/seantur/ray_tracer_challenge/tuples"
)

type projectile struct {
	position tuples.Tuple
	velocity tuples.Tuple
}

type environment struct {
	gravity tuples.Tuple
	wind    tuples.Tuple
}

func tick(env environment, proj projectile) projectile {
	position := tuples.Add(proj.position, proj.velocity)
	velocity := tuples.Add(proj.velocity, env.gravity)
	velocity = tuples.Add(velocity, env.wind)

	return projectile{position, velocity}
}

func main() {
	v := tuples.Vector(2, 2, 0)
	p := projectile{tuples.Point(0, 1, 0), v.Normalize()}
	e := environment{tuples.Vector(0, -0.1, 0), tuples.Vector(-0.01, 0, 0)}

	for i := 1; i <= 10; i++ {
		p = tick(e, p)

		fmt.Printf("%d: position: <%f %f %f> velocity: <%f %f %f %f>\n",
			i, p.position.X, p.position.Y, p.position.Z, p.velocity.X, p.velocity.Y, p.velocity.Z, p.velocity.W)
	}

}
