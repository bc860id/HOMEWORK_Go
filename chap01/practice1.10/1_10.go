
// 第1章 練習問題1.10
package main
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"strings"
)

var prefix = []string{ "1st", "2nd" }

func main() {
	for i := 0; i < 2; i++ {
		start := time.Now()
		ch := make(chan string)
		for _, url := range os.Args[1:] {
			go fetch(i, url, ch)
		}
		for range os.Args[1:] {
			fmt.Println(<-ch)
		}
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
}

func fetch(cnt int, url string, ch chan<- string) {
	var fname string
	fname = prefix[cnt] + "_" + url + ".txt"
	fname = strings.Replace(fname, "/", "_", -1)
	fname = strings.Replace(fname, ":", "", -1)
	fmt.Printf("fname: %s\n", fname)
	f, err := os.Create(fname)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		f.Close()
		return
	}

	nbytes, err := io.Copy(f, resp.Body)
	f.Close()
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}


