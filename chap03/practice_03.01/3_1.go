
// 第3章 練習問題3.1
package main
import (
	"fmt"
	"math"
	"os"
)

const (
	//width, height	= 600, 320
	width, height	= 600, 100
	cells			= 100
	xyrange			= 30.0
	xyscale			= width / 2/ xyrange
	zscale			= height * 0.4
	angle			= math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		// キャンバスのY軸方向の領域(height)に収まらなければ排除.
		// 4点確認するのは大変なので、最小側はby、最大側はdyを検査.
		// corner()はX軸方向についてはwidthを考慮した値を戻している.
		// 領域外を強制的に発生させるためheightを小さく定義.
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner(i + 1, j + 1)
			if ( (0 <= by) && (dy <= height) ) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Fprintf(os.Stderr,
					"skip! out of canvas:(%d, %d) " +
					"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					i, j, ax, ay, bx, by, cx ,cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	z := f(x, y)

	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

