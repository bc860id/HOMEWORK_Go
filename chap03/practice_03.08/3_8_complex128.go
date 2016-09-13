
// 第3章 練習問題3.8
package main
import (
	"image"
	"image/color"
	"image/png"
	//"math/cmplx"
	"os"
	"strconv"
	//"fmt"
)

func main() {
	const (
		xmin, ymin, xmax, ymax	= -2, -2, +2, +2
		width_base, height_base	= 1024, 1024
	)
	width := width_base
	height := height_base

	if len(os.Args) >= 2 {
		m, err := strconv.Atoi(os.Args[1])
		if err == nil {
			width *= m
			height *= m
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(((float64(py) / float64(height)) * (ymax - ymin)) + ymin)
		for px := 0; px < width; px++ {
			x := float64(((float64(px) / float64(width)) * (xmax - xmin)) + xmin)
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z_r, z_i float64) color.Color {
	const iterations	= 200
	const contrast		= 15
	var v_r, v_i, v_r_tmp float64

	for n := uint8(0); n < iterations; n++ {
		v_r_tmp = v_r
		v_r = ((v_r * v_r) - (v_i * v_i) + z_r)
		v_i = ((2 * v_r_tmp * v_i) + z_i)

		if ((v_r * v_r) + (v_i * v_i)) > (2 * 2) {
			return color.Gray{255 - (contrast * n)}
		}
	}

	return color.Black
}

