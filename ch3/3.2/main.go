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
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			// Ex 3.1
			if ok1 && ok2 && ok3 && ok4 {
				canvas.Writer.Write([]byte(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: black; fill: none; stroke-width: 0.7' />\n",
					ax, ay, bx, by, cx, cy, dx, dy)))
			}
		}
	}
	canvas.End()
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (float64, bool) {
	res := math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12

	// Ex 2.1 Check for non-finite float64 number
	if math.IsInf(res, 0) || math.IsNaN(res) {
		return 0, false
	}
	return res, true
}

//!-
