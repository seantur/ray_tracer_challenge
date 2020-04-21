package main

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"github.com/seantur/ray_tracer_challenge/scene"
	"math"
	"os"
	"runtime/pprof"
)

func saveScene(path string) {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	floor := raytracing.GetSphere()
	floor.SetTransform(datatypes.GetScaling(10, 0.01, 10))
	mat := floor.GetMaterial()
	mat.Color = raytracing.Color{Red: 1, Green: 0.9, Blue: 0.9}
	mat.Specular = 0
	floor.SetMaterial(mat)

	leftWall := raytracing.GetSphere()
	transform := datatypes.Multiply(datatypes.GetTranslation(0, 0, 5), datatypes.GetRotationY(-math.Pi/4))
	transform = datatypes.Multiply(transform, datatypes.GetRotationX(math.Pi/2))
	transform = datatypes.Multiply(transform, datatypes.GetScaling(10, 0.01, 10))
	leftWall.SetTransform(transform)
	leftWall.SetMaterial(floor.GetMaterial())

	rightWall := raytracing.GetSphere()
	transform = datatypes.Multiply(datatypes.GetTranslation(0, 0, 5), datatypes.GetRotationY(math.Pi/4))
	transform = datatypes.Multiply(transform, datatypes.GetRotationX(math.Pi/2))
	transform = datatypes.Multiply(transform, datatypes.GetScaling(10, 0.01, 10))
	rightWall.SetTransform(transform)
	rightWall.SetMaterial(floor.GetMaterial())

	middle := raytracing.GetSphere()
	middle.SetTransform(datatypes.GetTranslation(-0.5, 1, 0.5))
	mat = middle.GetMaterial()
	mat.Color = raytracing.Color{Red: 1, Green: 0, Blue: 0}
	mat.Diffuse = 0.7
	mat.Specular = 0.3
	middle.SetMaterial(mat)

	right := raytracing.GetSphere()
	right.SetTransform(datatypes.Multiply(datatypes.GetTranslation(1.5, 0.5, -0.5), datatypes.GetScaling(0.5, 0.5, 0.5)))
	mat = right.GetMaterial()
	mat.Color = raytracing.Color{Red: 0.5, Green: 1, Blue: 0.1}
	mat.Diffuse = 0.7
	mat.Specular = 0.3
	right.SetMaterial(mat)

	left := raytracing.GetSphere()
	left.SetTransform(datatypes.Multiply(datatypes.GetTranslation(-1.5, 0.33, -0.75), datatypes.GetScaling(0.33, 0.33, 0.33)))
	mat = left.GetMaterial()
	mat.Color = raytracing.Color{Red: 1, Green: 0.8, Blue: 0.1}
	mat.Diffuse = 0.7
	mat.Specular = 0.3
	left.SetMaterial(mat)

	world := scene.GetWorld()
	world.Shapes = []raytracing.Shape{floor, leftWall, rightWall, middle, right, left}

	camera := scene.GetCamera(1000, 500, math.Pi/2)
	camera.Transform = datatypes.ViewTransform(datatypes.Point(0, 3.5, -5), datatypes.Point(0, 1, 0), datatypes.Vector(0, 1, 0))

	canvas := camera.Render(world)
	canvas.SavePPM(path)
}

func main() {
	saveScene("1st_scene.ppm")
}
