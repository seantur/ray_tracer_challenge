package main

import (
	"fmt"
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"github.com/seantur/ray_tracer_challenge/scene"
	"github.com/seantur/ray_tracer_challenge/shapes"
	"math"
	"time"
)

func saveScene(path string) {
	room := shapes.GetCube()
	mat := room.GetMaterial()
	mat.Pattern = raytracing.GetCheckers(raytracing.HexColor(raytracing.White), raytracing.HexColor(raytracing.Black))
	mat.Pattern.SetTransform(datatypes.GetScaling(0.1, 0.1, 0.1))
	room.SetMaterial(mat)
	room.SetTransform(datatypes.GetScaling(100, 100, 100))

	obj := shapes.GetSphere()
	//obj.Min = 0
	//obj.Max = 1
	//obj.Closed = true
	obj.SetTransform(datatypes.GetTransform(datatypes.GetRotationX(math.Pi/4), datatypes.GetTranslation(0, 0, -5)))

	mat = obj.GetMaterial()
	mat.RGB = raytracing.HexColor(raytracing.Red)
	//mat.Transparency = 1
	//mat.RefractiveIndex = 1.52
	mat.Specular = 0.5
	//mat.Reflective = 0.5
	obj.SetMaterial(mat)

	world := scene.GetWorld()
	world.Light.Position = datatypes.Point(10, 10, 10)
	world.Shapes = []shapes.Shape{room, obj}

	camera := scene.GetCamera(500, 500, math.Pi/3)
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
