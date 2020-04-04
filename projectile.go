package main

import (
	"fmt"
	"github.com/seantur/ray_tracer_challenge/canvas"
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
	start := tuples.Point(0, 1, 0)
	velocity := tuples.Vector(1, 1.8, 0)
	velocity = velocity.Normalize()
	velocity = velocity.Multiply(11.25)

	p := projectile{start, velocity}
	gravity := tuples.Vector(0, -0.1, 0)
	wind := tuples.Vector(-0.01, 0, 0)

	color := canvas.Color{Red: 1, Green: 0, Blue: 0}

	e := environment{gravity, wind}

	c := canvas.Canvas{Height: 550, Width: 900}
	c.Init()

	for i := 0; i < 200; i++ {
		p = tick(e, p)
		fmt.Println(int(p.position.X), int(550-p.position.Y))
		c.WritePixel(int(p.position.X), int(550-p.position.Y), color)
	}

	c.SavePPM("tmp.ppm")

}
