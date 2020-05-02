package main

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"github.com/seantur/ray_tracer_challenge/scene"
	"math"
)

func saveScene(path string) {
	floorColor := raytracing.Color{Red: 1, Green: 0.9, Blue: 0.9}
	black := raytracing.Color{Red: 0, Green: 0, Blue: 0}

	floor := raytracing.GetPlane()
	mat := floor.GetMaterial()
	mat.Pattern = raytracing.GetCheckers(floorColor, black)
	mat.Specular = 0
	floor.SetMaterial(mat)

	leftWall := raytracing.GetPlane()
	transform := datatypes.Multiply(datatypes.GetTranslation(0, 0, 5), datatypes.GetRotationY(-math.Pi/4))
	transform = datatypes.Multiply(transform, datatypes.GetRotationX(math.Pi/2))
	leftWall.SetTransform(transform)
	leftWall.SetMaterial(floor.GetMaterial())

	rightWall := raytracing.GetPlane()
	transform = datatypes.Multiply(datatypes.GetTranslation(0, 0, 5), datatypes.GetRotationY(math.Pi/4))
	transform = datatypes.Multiply(transform, datatypes.GetRotationX(math.Pi/2))
	rightWall.SetTransform(transform)
	rightWall.SetMaterial(floor.GetMaterial())

	middle := raytracing.GetSphere()
	middle.SetTransform(datatypes.GetTranslation(-0.5, 1, 0.5))
	mat = middle.GetMaterial()
	mat.Color = raytracing.Color{Red: 1, Green: 0, Blue: 0}
	mat.Diffuse = 0.7
	mat.Specular = 0.3
	mat.Pattern = raytracing.GetGradient(raytracing.Color{Red: 1, Green: 0, Blue: 0}, raytracing.Color{Red: 0, Green: 1, Blue: 0})
	mat.Pattern.SetTransform(datatypes.Multiply(datatypes.GetTranslation(-1, 0, 0), datatypes.GetScaling(2, 1, 1)))
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
	world.Shapes = []raytracing.Shape{floor, middle, right, left}

	camera := scene.GetCamera(500, 500, math.Pi/3)
	camera.Transform = datatypes.ViewTransform(datatypes.Point(0, 1.5, -5), datatypes.Point(0, 1, 0), datatypes.Vector(0, 1, 0))

	canvas := camera.RenderConcurrent(world)

	canvas.SavePPM(path)
}

func main() {
	saveScene("scene.ppm")
}
