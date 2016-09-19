
// 第3章 練習問題3.9
package main
import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

	"net/http"
	"net/url"
	"strconv"
	"io"
	"log"
	"fmt"
)

const (
	xmin, ymin, xmax, ymax		= -2, -2, +2, +2
	base_width, base_height		= 1024, 1024
	origin_x, origin_y float64	= 0, 0
)

func main() {
	http.HandleFunc("/", queryhandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func queryhandler(w http.ResponseWriter, r *http.Request) {
	//var m map[string][]string
	//var err os.Error
	//var zoom float64

	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "URL parse error:%v\n", err)
		return
	}

	x_org := origin_x
	y_org := origin_y
	width := float64(base_width)
	height := float64(base_height)

	if val, exist := m["zoom"]; exist == true {
		zoom, err_zoom := strconv.ParseFloat(val[0], 64)
		if err_zoom != nil {
			fmt.Fprintf(os.Stderr, "zoom error:%v\n", err_zoom)
			return
		}
		width = float64(width) * zoom
		height = float64(height) * zoom

		//fmt.Fprintf(os.Stderr, "zoom:%f w:%f h:%f\n", zoom, width, height)
	}

	if val, exist := m["x"]; exist == true {
		x_org_shift, err_x := strconv.ParseFloat(val[0], 64)
		if err_x != nil {
			fmt.Fprintf(os.Stderr, "orign x error:%v\n", err_x)
			return
		}
		x_org = x_org_shift
	}

	if val, exist := m["y"]; exist == true {
		y_org_shift, err_y := strconv.ParseFloat(val[0], 64)
		if err_y != nil {
			fmt.Fprintf(os.Stderr, "orign y error:%v\n", err_y)
			return
		}
		y_org = y_org_shift
	}

	if _, exist := m["terminate"]; exist == true {
		os.Exit(1)
	}

	draw(int(width), int(height), x_org, y_org, w)
}

func draw(width, height int, x_org, y_org float64, out io.Writer) {

	//fmt.Fprintf(os.Stderr, "w:%d h:%d x:%f y:%f\n", width, height, x_org, y_org)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := ((float64(py) / float64(height)) * (ymax - ymin)) + ymin
		y -= y_org;
		for px := 0; px < width; px++ {
			x := ((float64(px) / float64(width)) * (xmax - xmin)) + xmin
			x -= x_org
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations	= 200
	const contrast		= 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = (v * v) + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - (contrast * n)}
		}
	}
	return color.Black
}

