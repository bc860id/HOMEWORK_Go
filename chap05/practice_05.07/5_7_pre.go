
// 第5章 練習問題5.7
package main
import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

var depth	int

func main() {
	doc, err := html.Parse(os.Stdin)
	if ( err != nil ) {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if ( pre != nil ) {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if ( post != nil ) {
		post(n)
	}
}

func startElement(n *html.Node) {
	if ( n.Type == html.ElementNode ) {
		fmt.Printf("%*s<%s", depth * 2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s='%s'", a.Key, a.Val)
		}
		fmt.Printf(">\n")
		depth++
	}
}

func endElement(n *html.Node) {
	if ( n.Type == html.ElementNode ) {
		depth--
		fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
	}
}

