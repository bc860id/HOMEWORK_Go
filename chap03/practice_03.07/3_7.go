
// 第3章 練習問題3.7
package main
import (
	"image"
	"image/color"
	"image/png"
	"os"
	//"fmt"
)

var count_max, count_min uint32
const limit = 10000	// とりあえず収束回数上限を固定値で与える。

func main() {
	const (
		//zoomout				= 100
		zoomout					= 4
		xmin, ymin, xmax, ymax	=
			-(1 * zoomout), -(1 * zoomout), +(1 * zoomout), +(1 * zoomout)
		width, height			= 1024, 1024
	)

	var converge_count	[width][height]uint32
	var converge_point	[width][height]complex128

	//fmt.Fprintf(os.Stderr, "pass1\n")

	for px := 0; px < width; px++ {
		x := (((float64(px) / float64(width)) * (xmax - xmin)) + xmin)
		for py := 0; py < height; py++ {
			y := (((float64(py) / float64(height)) * (ymax - ymin)) + ymin)

			//fmt.Fprintf(os.Stderr, "(%d, %d) (%010f, %010f)\n", px, py, x, y)

			z, count := converge(complex(x, y))

			if ( count < count_min ) { count_min = count }
			if ( count_max < count ) { count_max = count }

			converge_count[px][py] = count
			converge_point[px][py] = z
		}
	}

	//fmt.Fprintf(os.Stderr, "pass2\n")

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			color := graduate(converge_point[px][py], converge_count[px][py])
			img.Set(px, py, color)
		}
	}

	//fmt.Fprintf(os.Stderr, "pass3\n")

	png.Encode(os.Stdout, img)
}

func graduate(z complex128, c uint32) color.RGBA {
	level := byte((float64(c) / float64(count_max - count_min)) * 255)
	color := color.RGBA{ A:0xFF }

	if ( z == 1i ) {
		color.B = level
		return color
	}

	if ( z == -1i ) {
		color.R = level
		return color
	}

	//color.G = level
	color.G = 0xFF
	return color
}

func converge(z complex128) (/* converge point */complex128,/* count */uint32) {
	var count uint32
	var z_prev complex128

	z_prev = z;
	for count = 0; count < limit; count++ {
		z = recursion(z)
		if z_prev == z { break; }
		z_prev = z
	}

	if ( count == limit ) {
		z = 0
		count = 0
	}

	//fmt.Fprintf(os.Stderr, "%010f %010fi %d\n", real(z), imag(z), count)

	return z, count
}

func recursion(/* z[n] */z complex128) /* z[n+1] */complex128 {
	// 「z^4 - 1 = 0」の虚数解部分なので簡略化して「z^2 + 1 = 0」を解く。
	return (z - (((z * z) + 1) / (2 * z)))
}


