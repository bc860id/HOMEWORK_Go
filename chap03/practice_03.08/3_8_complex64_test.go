
// 第3章 練習問題3.8
package main
import (
	"image"
	"image/color"
	"image/png"
	//"math/cmplx"
	"os"
	"strconv"
	"testing"
	//"fmt"
)

const (
	xmin, ymin, xmax, ymax	= -2, -2, +2, +2
	width_base, height_base	= 1024, 1024
)

func main() {
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
		y := float32(((float32(py) / float32(height)) * (ymax - ymin)) + ymin)
		for px := 0; px < width; px++ {
			x := float32(((float32(px) / float32(width)) * (xmax - xmin)) + xmin)
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z_r, z_i float32) color.Color {
	const iterations	= 200
	const contrast		= 15
	var v_r, v_i, v_r_tmp float32

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

func BenchmarkForMeasure(b *testing.B) {
	/*
	for i := 0; i < b.N; i++ {
		ForMeasure()
	}
	*/
	ForMeasure()
}

func ForMeasure() {
	width := width_base
	height := height_base

	/*
	f, err := os.Create("measure_complex64.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "file create error:%v\n", err)
		return
	}
	*/

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(((float32(py) / float32(height)) * (ymax - ymin)) + ymin)
		for px := 0; px < width; px++ {
			x := float32(((float32(px) / float32(width)) * (xmax - xmin)) + xmin)
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	//png.Encode(os.Stdout, img)
	/*
	png.Encode(f, img)
	f.Close()
	*/
	png.Encode(os.Stderr, img)
}

