
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
	forEachNode(doc, startElement1, startElement2, endElement1, endElement2)
}

func forEachNode(n *html.Node, pre1, pre2, post1, post2 func(n *html.Node)) {
	var pre, post func(n *html.Node)

	if ( n.FirstChild != nil ) {
		pre = pre1
		post = post1
	} else {
		pre = pre2
		post = post2
	}

	if ( pre != nil ) {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre1, pre2, post1, post2)
	}

	if ( post != nil ) {
		post(n)
	}
}

func startElement1(n *html.Node) {
	if ( n.Type == html.ElementNode ) {
		fmt.Printf("%*s<%s", depth * 2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s='%s'", a.Key, a.Val)
		}
		fmt.Printf(">\n")
		depth++
	}
}

func startElement2(n *html.Node) {
	if ( n.Type == html.ElementNode ) {
		fmt.Printf("%*s<%s", depth * 2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s='%s'", a.Key, a.Val)
		}
		depth++
	}
}

func endElement1(n *html.Node) {
	if ( n.Type == html.ElementNode ) {
		depth--
		fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
	}
}

func endElement2(n *html.Node) {
	if ( n.Type == html.ElementNode ) {
		depth--
		fmt.Printf("/>\n")
	}
}

