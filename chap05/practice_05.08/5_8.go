
// 第5章 練習問題5.8
package main
import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

var depth		int
var id_compare	string

func main() {
	doc, err := html.Parse(os.Stdin)
	if ( err != nil ) {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	//fmt.Println("args:", len(os.Args))
	if ( len(os.Args) < 2 ) {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	id_compare = os.Args[1]

	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	var continue_scan	bool
	if ( pre != nil ) {
		continue_scan = pre(n)
	}

	if ( continue_scan == true ) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, pre, post)
		}
	}

	if ( post != nil ) {
		post(n)
	}
}

func startElement(n *html.Node) bool {
	result := true
	if ( n.Type == html.ElementNode ) {
		if ( ElementByID(n, id_compare) == nil ) {
			fmt.Printf("%*s<%s", depth * 2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s='%s'", a.Key, a.Val)
			}
			fmt.Printf(">\n")
			result = false
		}
		depth++
	}
	return result
}

func endElement(n *html.Node) bool {
	result := true
	if ( n.Type == html.ElementNode ) {
		depth--
		if ( ElementByID(n, id_compare) == nil ) {
			fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
			result = false
		}
	}
	return result
}

func ElementByID(doc *html.Node, id string) *html.Node {
	if ( doc.Type != html.ElementNode ) { return doc }
	for _, a := range doc.Attr {
		if ( (a.Key == "id") && (a.Val == id) ) { return nil }
	}
	return doc
}

