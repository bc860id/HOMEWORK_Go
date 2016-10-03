
// 第4章 練習問題4.8
package main
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	categories := make(map[string]int)
	var utflen[utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if ( err == io.EOF ) {
			break
		}
		if ( err != nil ) {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if ( (r == unicode.ReplacementChar) && (n == 1) ) {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if ( unicode.IsLetter(r) == true ) {
			categories["letter"]++
		}
		if ( unicode.IsDigit(r) == true ) {
			categories["number"]++
		}
		if ( unicode.IsSpace(r) == true ) {
			categories["space"]++
		}
		if ( unicode.IsLower(r) == true ) {
			categories["lower"]++
		}
		if ( unicode.IsUpper(r) == true ) {
			categories["upper"]++
		}
		if ( unicode.IsTitle(r) == true ) {
			categories["title"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	for t, n := range categories {
		fmt.Printf("%s\t%d\n", t, n)
	}
	if ( invalid > 0 ) {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

