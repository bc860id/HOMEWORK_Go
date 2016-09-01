
// 第3章 練習問題3.4
package main
import (
	"fmt"
	"math"
	"os"
	"net/http"
	"log"
)

const (
	zoom			= 3
	width, height	= 600 * zoom, 320 * zoom
	cells			= 100
	xyrange			= 30.0
	xyscale			= width / 2/ xyrange
	zscale			= height * 0.4
	angle			= math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var z_max, z_min float64 = 0, 0
var colorstep float64

func main() {
	// zの値域を確認して最大値と最小値を記録する。
	getrange()

	http.HandleFunc("/", drawhandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func drawhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: yellow; stroke-width: 0.7' " +
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		// 一応全部書いてもらわないと色が抜けるポリゴンが出るので、
		// 描画は全てに対して行い異常値はエラーを出す様にする。
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner(i + 1, j + 1)
			color := getcolor(i, j)
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

func getrange() {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i) / cells - 0.5)
			y := xyrange * (float64(j) / cells - 0.5)
			z := f(x, y)
			if ( math.IsNaN(z) == true ) { continue }
			if ( z > z_max ) {
				z_max = z
			} else if ( z < z_min ) {
				z_min = z
			}
		}
	}

	colorstep = ((z_max - z_min) / 255)
}

func getcolor(i, j int) uint64 {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
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

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	z := f(x, y)
	if ( math.IsNaN(z) == true ) {
		z = z_max
	}

	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

