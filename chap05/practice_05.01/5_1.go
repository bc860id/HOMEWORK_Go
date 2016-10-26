
// 第5章 練習問題5.1
package main
import (
	"fmt"
	"os"

	// $GOPATH C:/Program Files xUAC/Go/src
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	/*
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	*/
	links := visit(nil, doc)
	scanNode(links, 0, len(links))
}

func scanNode(links []string, index, total int) {
	if index >= total { return }
	fmt.Println(links[index])
	scanNode(links, index + 1, total)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
