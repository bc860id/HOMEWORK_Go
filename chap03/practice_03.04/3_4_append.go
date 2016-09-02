
// 第3章 練習問題3.4 後半
package main
import (
	"fmt"
	"math"
	"os"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	def_width, def_height	= 600, 320
	def_cells				= 100
	def_xyrange				= 30.0
	def_line				= "yellow"
	angle					= math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", drawhandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func drawhandler(w http.ResponseWriter, r *http.Request) {
	// パラメータを決定する。
	width, height, cells, xyrange, xyscale, zscale, line := determparam(r)

	// zの値域を確認して最大値と最小値を記録する。
	z_max, z_min, colorstep := getrange(cells, xyrange)

	w.Header().Set("Content-type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: %s; stroke-width: 0.7' " +
		"width='%d' height='%d'>", line, width, height)
	for i := 0; i < cells; i++ {
		// 一応全部書いてもらわないと色が抜けるポリゴンが出るので、
		// 描画は全てに対して行い異常値はエラーを出す様にする。
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j,
				width, height, cells, xyrange, xyscale, zscale, z_max)
			bx, by := corner(i, j,
				width, height, cells, xyrange, xyscale, zscale, z_max)
			cx, cy := corner(i, j + 1,
				width, height, cells, xyrange, xyscale, zscale, z_max)
			dx, dy := corner(i + 1, j + 1,
				width, height, cells, xyrange, xyscale, zscale, z_max)
			color := getcolor(i, j, cells, xyrange, colorstep, z_min)
			fmt.Fprintf(
				w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%06X'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)

			if ( !((0 <= by) && (dy <= height)) ) {
				fmt.Fprintf(os.Stderr,
					"skip! out of canvas:(%d, %d) " +
					"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					i, j, ax, ay, bx, by, cx ,cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(w, "</svg>\n")
}

func determparam(r *http.Request) (/* width		*/float64,
								   /* height	*/float64,
								   /* cells		*/int,
								   /* xyrange	*/float64,
								   /* xyscale	*/float64,
								   /* zscale	*/float64,
								   /* line		*/string) {
	var width	float64	= def_width
	var height	float64	= def_height
	var cells	int		= def_cells
	var xyrange	float64	= def_xyrange
	var xyscale	float64
	var zscale	float64
	var line	string	= def_line

	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "URL parse error:%v\n", err)
		xyscale	= (width / 2 / xyrange)
		zscale = (height * 0.4)
		return width, height, cells, xyrange, xyscale, zscale, line
	}

	if val, exist := m["width"]; exist == true {
		width_bup := width
		//width, err = strconv.Atoui(val[0])
		width, err = strconv.ParseFloat(val[0], 64)
		if err != nil {
			width = width_bup
			fmt.Fprintf(os.Stderr, "width error:%v\n", err)
		}
	}

	if val, exist := m["height"]; exist == true {
		height_bup := width
		//height, err = strconv.Atoui(val[0])
		height, err = strconv.ParseFloat(val[0], 64)
		if err != nil {
			height = height_bup
			fmt.Fprintf(os.Stderr, "height error:%v\n", err)
		}
	}

	if val, exist := m["cells"]; exist == true {
		cells_bup := cells
		cells, err = strconv.Atoi(val[0])
		if err != nil {
			cells = cells_bup
			fmt.Fprintf(os.Stderr, "cells error:%v\n", err)
		}
	}

	if val, exist := m["xyrange"]; exist == true {
		xyrange_bup := xyrange
		xyrange, err = strconv.ParseFloat(val[0], 64)
		if err != nil {
			xyrange = xyrange_bup
			fmt.Fprintf(os.Stderr, "xyrange error:%v\n", err)
		}
	}

	if val, exist := m["line"]; exist == true {
		line = val[0]
	}

	xyscale	= (width / 2 / xyrange)
	zscale = (height * 0.4)
	return width, height, cells, xyrange, xyscale, zscale, line
}

func getrange(cells int, xyrange float64) (/* z_max		*/float64,
										   /* z_min		*/float64,
										   /* colorstep	*/float64) {
	var z_min, z_max float64 = 0, 0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * ((float64(i) / float64(cells)) - 0.5)
			y := xyrange * ((float64(j) / float64(cells)) - 0.5)
			z := f(x, y)
			if ( math.IsNaN(z) == true ) { continue }
			if ( z > z_max ) {
				z_max = z
			} else if ( z < z_min ) {
				z_min = z
			}
		}
	}

	colorstep := ((z_max - z_min) / 255)
	return z_max, z_min, colorstep
}

func getcolor(i, j int, cells int, xyrange, colorstep, z_min float64) uint64 {
	x := xyrange * ((float64(i) / float64(cells)) - 0.5)
	y := xyrange * ((float64(j) / float64(cells)) - 0.5)
	z := f(x, y)

	// 極はゼロ割でぶっ飛んでるが、
	// sin(r)が2乗ノルムなので正値から0に近づくと考えて、
	// 符号は正で無限大としておく。
	// (x == 0) && (z == 0)で見てもいい。
	if ( math.IsNaN(z) == true ) {
		return 0x00FF0000
	}

	colornotch := uint64((z - z_min) / colorstep)
	return((0x000000FF - colornotch) + (colornotch << 16))
}

func corner(i, j int,
	width, height float64,
	cells int, xyrange, xyscale, zscale, z_max float64) (/* sx	*/float64,
														 /* sy	*/float64) {
	x := xyrange * ((float64(i) / float64(cells)) - 0.5)
	y := xyrange * ((float64(j) / float64(cells)) - 0.5)
	z := f(x, y)
	if ( math.IsNaN(z) == true ) {
		z = z_max
	}

	sx := (width / 2) + ((x - y) * cos30 * xyscale)
	sy := (height / 2) + ((x + y) * sin30 * xyscale) - (z * zscale)
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return (math.Sin(r) / r)
}

