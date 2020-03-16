// This package generates an image of the mandelbrot set
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Parameters struct {
	width  int
	height int

	minx float64
	maxx float64
	miny float64
	maxy float64

	stepx float64
	stepy float64

	max_iterations int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func get_color(px int, py int, params Parameters) uint8 {
	x0 := params.minx + (float64(px) * params.stepx)
	y0 := params.miny + (float64(py) * params.stepy)

	x := 0.0
	y := 0.0

	iterations := 0
	for iterations <= params.max_iterations && x*x+y*y < 2*2 {
		x2 := float64(x*x-y*y) + x0
		y = float64(2*x*y) + y0
		x = x2
		iterations++
	}
	return uint8(min(iterations*4, 155))
}

func main() {
	width := 1400
	height := 800

	minx := -2.5
	maxx := 1.0
	miny := -1.0
	maxy := 1.0

	stepx := float64(math.Abs(maxx-minx)) / float64(width)
	stepy := float64(math.Abs(maxy-miny)) / float64(height)

	max_iterations := 1000

	params := Parameters{width, height, minx, maxx, miny, maxy, stepx,
		stepy, max_iterations}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pxl_color := color.RGBA{0, 0,
				get_color(x, y, params) + uint8(100),
				0xff}
			img.Set(x, y, pxl_color)
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
