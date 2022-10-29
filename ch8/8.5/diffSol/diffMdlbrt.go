package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	maxGors                = 100
)

func main() {
	var start time.Time

	for n := 1; n < maxGors; n++ {
		start = time.Now()
		buildImg(n) // img result omitted
		fmt.Printf("For %d gors: time elapsed %fms\n", n, time.Since(start).Seconds()*1000)
		// png.Encode(os.Stdout, img) // NOTE: ignoring errors
	}

}

func buildImg(gors int) *image.RGBA {

	var wg sync.WaitGroup

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	yAxis := make(chan int, height)

	for i := 0; i < gors; i++ {
		wg.Add(1)
		go execImage(yAxis, &wg, img)
	}

	for y := 0; y < height; y++ {
		yAxis <- y // sending goroutine
	}
	close(yAxis)
	wg.Wait()

	return img
}

func execImage(yAxis <-chan int, wg *sync.WaitGroup, img *image.RGBA) {
	defer wg.Done()

	for py := range yAxis { // receiving goroutine
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
