package main

import (
	"github.com/seantur/ray_tracer_challenge/canvas"
	"github.com/seantur/ray_tracer_challenge/matrices"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"github.com/seantur/ray_tracer_challenge/tuples"
	"math"
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

func saveProjectile(path string) {
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
		c.WritePixel(int(p.position.X), int(550-p.position.Y), color)
	}

	c.SavePPM(path)
}

func saveClock(path string, size int) {

	scale := 3 / 8. * float64(size)

	c := canvas.Canvas{Height: size, Width: size}
	c.Init()

	rot := matrices.GetRotationY(math.Pi * 2 / 12.)

	white := canvas.Color{Red: 255, Green: 255, Blue: 255}

	offset := int(size / 2)
	p := tuples.Point(0, 0, 1)

	for i := 0; i < 12; i++ {
		p = matrices.TupleMultiply(rot, p)
		c.WritePixel(int(p.X*scale)+offset, int(p.Z*scale)+offset, white)
	}

	c.SavePPM(path)
}

func saveShadow(path string) {
	canvas_pixels := 500
	wall_z := 10.0
	wall_size := 7.0
	pixel_size := wall_size / float64(canvas_pixels)
	half := wall_size / 2.0

	c := canvas.Canvas{Height: canvas_pixels, Width: canvas_pixels}
	c.Init()

	red := canvas.Color{Red: 1, Green: 0, Blue: 0}
	shape := raytracing.Sphere{}
	shape.Init()

	ray_origin := tuples.Point(0, 0, -5)

	for y := 0; y < canvas_pixels; y++ {
		world_y := half - pixel_size*float64(y)
		for x := 0; x < canvas_pixels; x++ {
			world_x := half - pixel_size*float64(x)
			position := tuples.Point(world_x, world_y, wall_z)

			pos := tuples.Subtract(position, ray_origin)
			r := raytracing.Ray{Origin: ray_origin, Direction: pos.Normalize()}
			xs := shape.Intersect(r)

			if len(xs) > 0 {
				_, err := raytracing.Hit(xs)
				if err == nil {
					c.WritePixel(x, y, red)
				}
			}
		}
	}

	c.SavePPM(path)

}

func main() {
	saveShadow("shadow.ppm")
}
