
// 第5章 練習問題5.2
package main
import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

var elemlist	map[string]int

func main() {
	elemlist = make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	elems := visit(nil, doc)
	scanNode(elems, 0, len(elems))

	for elem, key := range elemlist {
		fmt.Println(elem, "\t\t", key)
	}
}

func scanNode(elems []string, index, total int) {
	if index >= total { return }
	scanNode(elems, index + 1, total)
}

func visit(elems []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		elems = append(elems, n.Data)
		elemlist[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elems = visit(elems, c)
	}
	return elems
}
