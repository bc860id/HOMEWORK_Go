
// 第3章 練習問題3.5
package main
import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax	= -2, -2, +2, +2
		width, height			= 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := ((float64(py) / height) * (ymax - ymin)) + ymin
		for px := 0; px < width; px++ {
			x := ((float64(px) / width) * (xmax - xmin)) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.RGBA {
	const iterations	= 200
	const contrast		= 15
	var v complex128
	//var c color.RGBA
	//c := color.RGBA{ R:0xFF, G:0xFF, B:0xFF, A:0xFF }
	c := color.RGBA{ R: 0xFF, G:0xFF, A:0xFF }

	for n := uint8(0); n < iterations; n++ {
		v = (v * v) + z
		if cmplx.Abs(v) > 2 {
			c.R = uint8(((math.Abs(real(v)) / 2) * 256))
			c.B = uint8(((math.Abs(imag(v)) / 2) * 256))
			c.G = 0
			//c.A = (contrast * n)
			c.A -= (contrast * n)
			return c
		}
	}
	return c
}

