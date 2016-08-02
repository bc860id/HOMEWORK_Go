
// 第1章 練習問題1.8
package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var prefix = []string{ "http://", "https://" }

func main() {
	for _, url := range os.Args[1:] {
		if (strings.HasPrefix(url, prefix[0]) == false) &&
		   (strings.HasPrefix(url, prefix[1]) == false) {
				url = prefix[0] + url
				fmt.Fprintf(os.Stderr, "\n\x1b[41;33mappend prefix\x1b[0m: %v\n", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}


