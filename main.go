package main

import (
	"fmt"
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"github.com/seantur/ray_tracer_challenge/scene"
	"math"
	"time"
)

func saveScene(path string) {
	floorColor := raytracing.RGB{Red: 1, Green: 0.9, Blue: 0.9}

	floor := raytracing.GetPlane()
	mat := floor.GetMaterial()
	mat.RGB = floorColor
	mat.Specular = 0.5
	mat.Reflective = 0.5
	floor.SetMaterial(mat)

	ceiling := raytracing.GetPlane()
	ceiling.SetTransform(datatypes.GetTranslation(0, 20, 0))
	mat.Specular = .5
	ceiling.SetMaterial(mat)

	middle := raytracing.GetSphere()
	middle.SetTransform(datatypes.GetTranslation(0, 1, 0))
	mat = middle.GetMaterial()
	mat.RGB = raytracing.HexColor(raytracing.Red)
	middle.SetMaterial(mat)

	right := raytracing.GetSphere()
	right.SetTransform(datatypes.GetTransform(datatypes.GetScaling(0.5, 0.5, 0.5), datatypes.GetTranslation(1.5, 0.5, -0.5)))
	mat = right.GetMaterial()
	mat.RGB = raytracing.HexColor(raytracing.Red)
	right.SetMaterial(mat)

	left := raytracing.GetSphere()
	left.SetTransform(datatypes.GetTransform(datatypes.GetScaling(0.33, 0.33, 0.33), datatypes.GetTranslation(-1.5, 0.33, -0.75)))
	mat = left.GetMaterial()
	mat.RGB = raytracing.HexColor(raytracing.Teal)
	mat.Diffuse = 0.7
	mat.Specular = 0.3
	left.SetMaterial(mat)

	world := scene.GetWorld()
	world.Shapes = []raytracing.Shape{floor, middle, left, right}

	camera := scene.GetCamera(1000, 1000, math.Pi/3)
	camera.Transform = datatypes.ViewTransform(datatypes.Point(0, 1.5, -5), datatypes.Point(0, 1, 0), datatypes.Vector(0, 1, 0))

	fmt.Println("Rendering...")
	start := time.Now()
	output := camera.RenderConcurrent(world)
	duration := time.Since(start)
	fmt.Printf("done (%v elapsed)\n", duration)

	scene.SavePng(output, path)
}

func main() {
	saveScene("scene.png")
}
