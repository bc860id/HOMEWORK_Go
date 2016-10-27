
// 第5章 練習問題5.3
package main
import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

var skip	bool = false

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if ( n.Type == html.ElementNode ) {
		if ( (n.Data == "script") || (n.Data == "style") ) {
			skip = true
		} else {
			skip = false
		}
	}
	if ( (n.Type == html.TextNode) && (skip == false) ) {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}

