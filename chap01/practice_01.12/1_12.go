
// 第1章 練習問題1.12
package main
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"

	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	Index_background = 0
	Index_curve = 1
)

type st_animattri struct {
	cycles		float64
	reso		float64
	size		int
	nframes		int
	delay		int
}

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}
var animattri = st_animattri{ cycles: 5, reso: 0.001, size: 100, nframes: 64, delay: 8}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		m, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			fmt.Fprintf(os.Stderr, "URL parse error:%v\n", err)
			return
		}

		if val, exist := m["cycles"]; exist == true {
			//animattri.cycles, err = strconv.Atof64(val[0])
			animattri.cycles, err = strconv.ParseFloat(val[0], 64)
		}

		if val, exist := m["reso"]; exist == true {
			//animattri.reso, err = strconv.Atof64(val[0])
			animattri.reso, err = strconv.ParseFloat(val[0], 64)
		}

		if val, exist := m["size"]; exist == true {
			animattri.size, err = strconv.Atoi(val[0])
		}

		if val, exist := m["nframes"]; exist == true {
			animattri.nframes, err = strconv.Atoi(val[0])
		}

		if val, exist := m["delay"]; exist == true {
			animattri.delay, err = strconv.Atoi(val[0])
		}

		lissajous(w)
	}
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
	//fmt.Printf("cycles:%f reso:%f size:%d nframes:%d delay:%d\n",
	//	animattri.cycles, animattri.reso, animattri.size,
	//	animattri.nframes, animattri.delay)

	freq := (rand.Float64() * 3.0)
	anim := gif.GIF{LoopCount: animattri.nframes}
	phase := 0.0

	for i := 0; i < animattri.nframes; i++ {
		rect := image.Rect(0, 0, (2 * animattri.size + 1), (2 * animattri.size + 1))
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < (animattri.cycles * 2 * math.Pi); t += animattri.reso {
			x := math.Sin(t)
			y := math.Sin((t * freq) + phase)
			img.SetColorIndex(animattri.size + int(x * float64(animattri.size) + 0.5),
				animattri.size + int(y * float64(animattri.size) + 0.5), Index_curve)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, animattri.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}


