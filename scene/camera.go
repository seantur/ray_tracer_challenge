package scene

import (
	"fmt"
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"image"
	"math"
	"runtime"
	"sync"
)

type camera struct {
	Hsize, Vsize                          int
	Fov, PixelSize, HalfWidth, HalfHeight float64
	Transform                             datatypes.Matrix
}

func GetCamera(hsize, vsize int, fov float64) camera {
	c := camera{Hsize: hsize, Vsize: vsize, Fov: fov, Transform: datatypes.GetIdentity()}

	half_view := math.Tan(fov / 2)
	aspect_ratio := float64(hsize) / float64(vsize)

	if aspect_ratio >= 1 {
		c.HalfWidth = half_view
		c.HalfHeight = half_view / aspect_ratio
	} else {
		c.HalfWidth = half_view * aspect_ratio
		c.HalfHeight = half_view
	}

	c.PixelSize = (c.HalfWidth * 2) / float64(c.Hsize)

	return c
}

func (c *camera) RayForPixel(px, py int) datatypes.Ray {
	xoffset := (float64(px) + 0.5) * c.PixelSize
	yoffset := (float64(py) + 0.5) * c.PixelSize

	world_x := c.HalfWidth - xoffset
	world_y := c.HalfHeight - yoffset

	transform_inv, _ := c.Transform.Inverse()

	pixel := datatypes.TupleMultiply(transform_inv, datatypes.Point(world_x, world_y, -1))
	origin := datatypes.TupleMultiply(transform_inv, datatypes.Point(0, 0, 0))

	direction := datatypes.Subtract(pixel, origin)
	direction = direction.Normalize()

	return datatypes.Ray{Origin: origin, Direction: direction}
}

func (c *camera) Render(w World) image.Image {
	im := InitCanvas(c.Hsize, c.Vsize)

	for y := 0; y < c.Vsize; y++ {
		for x := 0; x < c.Hsize; x++ {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r, 5)
			im.Set(x, y, color)
		}
	}

	return im
}

type pnt struct {
	x, y int
}

func worker(channel chan pnt, w World, c *camera, im *image.RGBA64, wg *sync.WaitGroup) {
	defer wg.Done()

	for pnt := range channel {
		r := c.RayForPixel(pnt.x, pnt.y)
		color := w.ColorAt(r, 5)
		im.Set(pnt.x, pnt.y, color)
	}
}

func (c *camera) RenderConcurrent(w World) image.Image {

	numWorkers := runtime.NumCPU() * 4

	im := InitCanvas(c.Hsize, c.Vsize)

	channel := make(chan pnt)
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	fmt.Printf("Starting %d goroutines\n", numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(channel, w, c, im, &wg)
	}

	for y := 0; y < c.Vsize; y++ {
		for x := 0; x < c.Hsize; x++ {
			channel <- pnt{x, y}
		}
	}

	close(channel)
	wg.Wait()

	return im
}
