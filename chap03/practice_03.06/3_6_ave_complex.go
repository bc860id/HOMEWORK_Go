
// 第3章 練習問題3.6 複素数値で平均化
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
			z := smooth(x, y, ((xmax - xmin) / width), ((ymax - ymin) / height))
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func smooth(x, y float64, pitch_x, pitch_y float64) complex128 {
	var sx	[2]float64
	var sy	[2]float64
	var sum complex128

	sx[0] = (x + (1 * (pitch_x / 4)))
	sx[1] = (x + (3 * (pitch_x / 4)))
	sy[0] = (y + (1 * (pitch_y / 4)))
	sy[1] = (y + (3 * (pitch_y / 4)))

	sum = 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			sum += complex(sx[i], sy[j])
		}
	}

	return (sum / 4)
}

func mandelbrot(z complex128) color.RGBA {
	const iterations	= 200
	const contrast		= 15
	var v complex128
	c := color.RGBA{ R: 0xFF, G:0xFF, A:0xFF }

	for n := uint8(0); n < iterations; n++ {
		v = (v * v) + z
		if cmplx.Abs(v) > 2 {
			c.R = uint8(((math.Abs(real(v)) / 2) * 256))
			c.B = uint8(((math.Abs(imag(v)) / 2) * 256))
			c.G = 0
			c.A -= (contrast * n)
			return c
		}
	}
	return c
}

