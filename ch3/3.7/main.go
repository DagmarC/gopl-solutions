// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"

	"github.com/DagmarC/gopl-solutions/utils"
)

// 16 colors, for same escape time
var palettes = []color.RGBA{
	{66, 30, 15, 255},
	{25, 7, 26, 255},
	{9, 1, 47, 255},
	{4, 4, 73, 255},
	{0, 7, 100, 255},
	{12, 44, 138, 255},
	{24, 82, 177, 255},
	{57, 125, 209, 255},
	{134, 181, 229, 255},
	{211, 236, 248, 255},
	{241, 233, 191, 255},
	{248, 201, 95, 255},
	{255, 170, 0, 255},
	{204, 128, 0, 255},
	{153, 87, 0, 255},
	{106, 52, 3, 255},
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	out := utils.GetFile("mandelbrot.png")
	defer out.Close()

	// Enlarge the image resolution
	imgMax := image.NewRGBA(image.Rect(0, 0, width*2, height*2))
	for py := 0; py < height*2; py++ {
		y := float64(py)/(height*2)*(ymax-ymin) + ymin

		for px := 0; px < width*2; px++ {
			x := float64(px)/(width*2)*(xmax-xmin) + xmin

			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			imgMax.Set(px, py, newton(z))
		}
	}

	// Supersampling. Obtain original image size.
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height*2; py += 2 {
		for px := 0; px < width*2; px += 2 {

			resColor := superSampling(imgMax, px, py)
			img.SetRGBA(px/2, py/2, resColor)
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func superSampling(maxImg *image.RGBA, x, y int) color.RGBA {
	pixelColors := []color.RGBA{
		maxImg.RGBAAt(x, y),
		maxImg.RGBAAt(x+1, y),
		maxImg.RGBAAt(x, y+1),
		maxImg.RGBAAt(x+1, y+1),
	}

	var r, g, b, a uint64

	for i := 0; i < 4; i++ {
		r += uint64(pixelColors[i].R)
		g += uint64(pixelColors[i].G)
		b += uint64(pixelColors[i].B)
		a += uint64(pixelColors[i].A)
	}
	return color.RGBA{R: uint8(r / 4), G: uint8(g / 4), B: uint8(b / 4), A: uint8(a / 4)}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palettes[n%16]
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return palettes[i%16]
		}
	}
	return color.Black
}
