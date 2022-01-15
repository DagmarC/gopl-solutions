// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"os"

	svg "github.com/ajstarks/svgo"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

type Extrema int

const (
	middle Extrema = iota
	valley
	peak
)

func FindExtrema(x float64) (e Extrema) {
	// To find Extremes -> use function derivates
	// f(x) = sin(x)/x, f'(x) = (x*cos(x)-sin(x))/x^2
	// f'(x) = 0 ==> x = tan(x), peak or vally
	// if f''(x) > 0, vally
	// if f''(x) < 0, peak
	// f''(x) = {2(sin(x)-x*cos(x)) - x*x*sin(x)}/x*x*x

	// f`(x) == 0
	//firstDer := math.Abs(x*math.Cos(x) - math.Sin(x))

	e = middle
	if int((2*(math.Sin(x)-x*math.Cos(x))-x*x*math.Sin(x))/x*x*x) == 0 {
		e = peak
		secDer := 2*(math.Sin(x)-x*math.Cos(x)) - x*x*math.Sin(x)
		if secDer > 0 {
			e = valley
		}
	}
	return e
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	out, err := os.OpenFile("tmp.svg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	canvas := svg.New(out)
	canvas.Start(width, height, "style='stroke: white; fill: black; stroke-width: 0.7'")

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, e1, ok1 := corner(i+1, j)
			bx, by, e2, ok2 := corner(i, j)
			cx, cy, e3, ok3 := corner(i, j+1)
			dx, dy, e4, ok4 := corner(i+1, j+1)

			// Ex 3.3
			c := "grey"
			if e1 == valley || e2 == valley || e3 == valley || e4 == valley {
				c = "blue"
			} else if e1 == peak || e2 == peak || e3 == peak || e4 == peak {
				c = "red"
			}

			if ok1 && ok2 && ok3 && ok4 {
				canvas.Writer.Write([]byte(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s; fill: none; stroke-width: 0.7' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, c)))
			}
		}
	}
	canvas.End()
}

func corner(i, j int) (float64, float64, Extrema, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, 0, false
	}
	extrema := FindExtrema(z)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, extrema, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	res := math.Sin(r) / r

	if math.IsInf(res, 0) || math.IsNaN(res) {
		return 0, false
	}
	return res, true
}

//!-
