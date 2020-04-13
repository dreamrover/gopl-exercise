// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

import (
	"log"
	"net/http"
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
	//!+http
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "image/svg+xml")
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				rgb := color(i, j)
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%06x'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, rgb)
			}
		}
		fmt.Fprintln(w, "</svg>")
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func color(i, j int) int {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	if z > 1.0 {
		z = 1.0
	} else if z < -1.0 {
		z = -1.0
	}
	return 0xff << uint((z+1.0)*8)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	//r := math.Hypot(x, y) // distance from (0,0)
	//return math.Sin(r) / r
	return 0.07*math.Cos(x) + 0.07*math.Sin(y)
	/*return (0.4*math.Sin(2*math.Pi/xyrange*x-0.5*math.Pi) +
	0.4*math.Cos(2*math.Pi/xyrange*y))*/
}

//!-
