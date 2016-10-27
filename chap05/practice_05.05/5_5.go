
// 第5章 練習問題5.5
package main
import (
	"fmt"
	"os"
	"net/http"
	"strings"
	"golang.org/x/net/html"

	//"io/ioutil"
)

func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if ( err != nil ) {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("words:", words, " images:", images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if ( err != nil ) {
		return
	}

	doc, err := html.Parse(resp.Body)
	fmt.Printf("%s\n", doc)
	resp.Body.Close()

	if ( err != nil ) {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if ( n.Data == "img" ) {
			images++
		}
		if ( n.Type == html.TextNode ) {
			fmt.Println(n.Data)
			words += len(strings.Fields(n.Data))
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImages(c)
	}
	return
}
