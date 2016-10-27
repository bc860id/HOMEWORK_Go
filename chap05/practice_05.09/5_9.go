
// 第5章 練習問題5.9
package main
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[1:])
	words := strings.Join(os.Args[1:], " ")
	fmt.Printf("%s\n", expand(words, sample1))
}

func expand(s string, f func(string) string) string {
	words := strings.Fields(s)

	for i, word := range words {
		if ( strings.HasPrefix(word, "$") == false ) { continue }
		words[i] = f(word[1:])
	}

	return strings.Join(words, " ")
}

func sample1(s string) string {
	return "add" + s
}

