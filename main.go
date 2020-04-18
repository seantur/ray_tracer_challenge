package main

import (
	"fmt"
	"github.com/seantur/ray_tracer_challenge/canvas"
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
	"time"
)

type projectile struct {
	position datatypes.Tuple
	velocity datatypes.Tuple
}

type environment struct {
	gravity datatypes.Tuple
	wind    datatypes.Tuple
}

func tick(env environment, proj projectile) projectile {
	position := datatypes.Add(proj.position, proj.velocity)
	velocity := datatypes.Add(proj.velocity, env.gravity)
	velocity = datatypes.Add(velocity, env.wind)

	return projectile{position, velocity}
}

func saveProjectile(path string) {
	start := datatypes.Point(0, 1, 0)
	velocity := datatypes.Vector(1, 1.8, 0)
	velocity = velocity.Normalize()
	velocity = velocity.Multiply(11.25)

	p := projectile{start, velocity}
	gravity := datatypes.Vector(0, -0.1, 0)
	wind := datatypes.Vector(-0.01, 0, 0)

	color := raytracing.Color{Red: 1, Green: 0, Blue: 0}

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

	rot := datatypes.GetRotationY(math.Pi * 2 / 12.)

	white := raytracing.Color{Red: 255, Green: 255, Blue: 255}

	offset := int(size / 2)
	p := datatypes.Point(0, 0, 1)

	for i := 0; i < 12; i++ {
		p = datatypes.TupleMultiply(rot, p)
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

	red := raytracing.Color{Red: 1, Green: 0, Blue: 0}
	shape := raytracing.GetSphere()

	ray_origin := datatypes.Point(0, 0, -5)

	for y := 0; y < canvas_pixels; y++ {
		world_y := half - pixel_size*float64(y)
		for x := 0; x < canvas_pixels; x++ {
			world_x := -half + pixel_size*float64(x)
			position := datatypes.Point(world_x, world_y, wall_z)

			pos := datatypes.Subtract(position, ray_origin)
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

func save3DSphere(path string) {

	start := time.Now()

	shape := raytracing.GetSphere()
	shape.Material.Color = raytracing.Color{Red: 1, Green: 0, Blue: 0}
	shape.SetTransform(datatypes.GetScaling(0.5, 0.5, 0.5))

	light := raytracing.PointLight{Position: datatypes.Point(-10, 10, -10), Intensity: raytracing.Color{Red: 1, Green: 1, Blue: 1}}

	canvas_pixels := 1000
	wall_z := 10.0
	wall_size := 7.0
	pixel_size := wall_size / float64(canvas_pixels)
	half := wall_size / 2.0

	c := canvas.Canvas{Height: canvas_pixels, Width: canvas_pixels}
	c.Init()

	ray_origin := datatypes.Point(0, 0, -5)

	for y := 0; y < canvas_pixels; y++ {
		world_y := half - pixel_size*float64(y)
		for x := 0; x < canvas_pixels; x++ {
			world_x := -half + pixel_size*float64(x)
			position := datatypes.Point(world_x, world_y, wall_z)

			pos := datatypes.Subtract(position, ray_origin)
			r := raytracing.Ray{Origin: ray_origin, Direction: pos.Normalize()}
			xs := shape.Intersect(r)

			if len(xs) > 0 {
				hit, err := raytracing.Hit(xs)
				if err == nil {
					point := r.Position(hit.T)
					normal := hit.Object.GetNormal(point)
					eye := r.Direction.Negate()

					color := raytracing.Lighting(hit.Object.Material, light, point, eye, normal)
					c.WritePixel(x, y, color)
				}

			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Rendering took %s\n", elapsed.Round(time.Millisecond))
	fmt.Println("Saving...")

	start = time.Now()
	c.SavePPM(path)
	elapsed = time.Since(start)
	fmt.Printf("Saving took %s\n", elapsed.Round(time.Millisecond))

}

func main() {
	save3DSphere("3Dsphere.ppm")
}
