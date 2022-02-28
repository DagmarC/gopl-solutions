// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

const (
	mandelbrotImg = iota
	acosImg
	sqrtImg
	newtonImg
	typeSize
)

func main() {
	// Query example: 
	// http://localhost:8000/?type=0&x=5000&y=400&scale=12
	// http://localhost:8000/?type=3&x=200&y=300&scale=15

	handler := func(w http.ResponseWriter, r *http.Request) {
		f, x, y, s := parseQuery(r.URL)
		fmt.Println(f, x, y, s)
		drawFractal(w, int(f)%typeSize, x, y, s)
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseQuery(u *url.URL) (f int64, x, y, s float64) {

	f, err := strconv.ParseInt(u.Query().Get("type"), 10, 32)
	if err != nil {
		f = 0
	}

	x, err = strconv.ParseFloat(u.Query().Get("x"), 64)
	if err != nil {
		x = 2
	}

	y, err = strconv.ParseFloat(u.Query().Get("y"), 64)
	if err != nil {
		y = 2
	}

	s, err = strconv.ParseFloat(u.Query().Get("scale"), 64)
	if err != nil {
		s = 1
	}

	return
}

func drawFractal(w http.ResponseWriter, f int, ix, iy, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	s := math.Abs(zoom)

	v := float64(2) / s // Determine boundaries of image.
	if math.IsNaN(v) {
		v = 2
	}

	dx, dy := ix/(s*width), iy/(s*height)
	if math.IsNaN(dx) {
		dx = 0
	}
	if math.IsNaN(dy) {
		dy = 0
	}

	var xmin, ymin, xmax, ymax = -v, -v, +v, +v

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			z := complex(x+dx, y+dy)
			// Image point (px, py) represents complex value z.

			switch f {
			case mandelbrotImg:
				img.Set(px, py, mandelbrot(z))
			case acosImg:
				img.Set(px, py, acos(z))
			case sqrtImg:
				img.Set(px, py, sqrt(z))
			case newtonImg:
				img.Set(px, py, newton(z))

			}
		}
	}

	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colorize(v, z, n)
		}
	}
	return color.Black
}

// colorize obtained from ray-g/gopl/ch03/ex3.09/mandelbrot.go
func colorize(v, z complex128, n uint8) color.Color {
	const contrast = 15
	blue := 255 - contrast*n
	red := 255 - blue
	green := lerp(red, blue, n%1)

	return color.RGBA{red, green, blue, 255}
}

// lerp: linear interpolation
func lerp(v0, v1, t uint8) uint8 {
	return v0 + t*(v1-v0)
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
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
