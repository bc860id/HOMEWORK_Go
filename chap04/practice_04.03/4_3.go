
// 第4章 練習問題4.3
package main
import (
	"fmt"
)

var a = [...]int{0, 1, 2, 3, 4, 5}

func main() {
	reverse(&a)
	fmt.Println(a)
}

func reverse(s *[len(a)]int) {
	for i, j := 0, (len(s) - 1); i < j; i, j = (i + 1), (j - 1) {
		s[i], s[j] = s[j], s[i]
	}
}

