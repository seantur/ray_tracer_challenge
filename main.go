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
	floor := raytracing.GetPlane()
	mat := floor.GetMaterial()
	mat.Pattern = raytracing.GetCheckers(raytracing.HexColor(raytracing.White), raytracing.HexColor(raytracing.Black))
	mat.Specular = 0.5
	mat.Reflective = 0.5
	floor.SetMaterial(mat)
	floor.SetTransform(datatypes.GetTransform(
		datatypes.GetRotationX(math.Pi/2),
		datatypes.GetTranslation(0, 0, -20)))

	sphere := raytracing.GetSphere()
	sphere.SetTransform(datatypes.GetTransform(
		datatypes.GetScaling(0.5, 0.5, 0.5),
		datatypes.GetTranslation(0, 0, -1)))
	mat = sphere.GetMaterial()
	mat.Diffuse = 0
	mat.Transparency = 1
	mat.RefractiveIndex = 1.52
	mat.Ambient = 0
	mat.Shininess = 300
	mat.Specular = 1
	mat.Reflective = 1
	sphere.SetMaterial(mat)

	bubble := raytracing.GetSphere()
	bubble.SetTransform(datatypes.GetTransform(
		datatypes.GetScaling(0.25, 0.25, 0.25),
		datatypes.GetTranslation(0, 0, -1)))
	mat = bubble.GetMaterial()
	mat.RefractiveIndex = 1.00029
	mat.Transparency = 1
	mat.Shininess = 300
	mat.Specular = 1
	mat.Reflective = 1
	bubble.SetMaterial(mat)

	world := scene.GetWorld()
	world.Light.Position = datatypes.Point(10, 10, 10)
	world.Shapes = []raytracing.Shape{floor, sphere} //, bubble}

	camera := scene.GetCamera(1000, 1000, math.Pi/3)
	camera.Transform = datatypes.ViewTransform(datatypes.Point(0, 0, 0), datatypes.Point(0, 0, -1), datatypes.Vector(0, 1, 0))

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
