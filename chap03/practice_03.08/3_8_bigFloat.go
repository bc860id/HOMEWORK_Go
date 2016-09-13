
// 第3章 練習問題3.8
package main
import (
	"image"
	"image/color"
	"image/png"
	//"math/cmplx"
	"math/big"
	"os"
	"strconv"
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
		y := toFloat(int64(py), int64(height), ymax, ymin)
		for px := 0; px < width; px++ {
			x := toFloat(int64(px), int64(width), xmax, xmin)
			img.Set(px, py, mandelbrot(&x, &y))
		}
	}
	png.Encode(os.Stdout, img)
}

func toFloat(grid int64, scope int64, max, min float64) big.Float {
	var f_grid, f_scope, f_max, f_min, f_range, ret big.Float
	f_grid.SetInt64(grid)
	f_scope.SetInt64(scope)
	f_max.SetFloat64(max)
	f_min.SetFloat64(min)
	f_range.Sub(&f_max, &f_min)

	ret.Quo(&f_grid, &f_scope)
	ret.Mul(&ret, &f_range)
	ret.Add(&ret, &f_min)

	return ret
}

func mandelbrot(z_r, z_i *big.Float) color.Color {
	const iterations	= 200
	const contrast		= 15
	var v_r, v_i, v_r_t, t1, t2, t3 big.Float	/* 0 initialized	*/

	for n := uint8(0); n < iterations; n++ {
		v_r_t.Copy(&v_r)
		v_r.Add(t3.Sub(t1.Mul(&v_r, &v_r), t2.Mul(&v_i, &v_i)), z_r)
		/*
		t1.Mul(&v_r, &v_r)
		t2.Mul(&v_i, &v_i)
		t3.Sub(&t1, &t2)
		v_r.Add(&t3, z_r)
		*/

		v_i.Add(t2.Mul(big.NewFloat(2), t1.Mul(&v_r_t, &v_i)), z_i)
		/*
		t1.Mul(&v_r_t, &v_i)
		t2.Mul(&t1, big.NewFloat(2))
		v_i.Add(&t2, z_i)
		*/

		t3.Add(t1.Mul(&v_r, &v_r), t2.Mul(&v_i, &v_i))
		/*
		t1.Mul(&v_r, &v_r)
		t2.Mul(&v_i, &v_i)
		t3.Add(&t1, &t2)
		*/
		if t3.Cmp(big.NewFloat(2 * 2)) > 0 {
			return color.Gray{255 - (contrast * n)}
		}
	}
	return color.Black
}

