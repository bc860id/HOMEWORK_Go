
// 第4章 練習問題4.4
package main
import (
	"fmt"
)

var a = [...]int{0, 1, 2, 3, 4, 5}

func main() {
	//var shift int = -6
	var shift int = 2

	s := a[0:len(a)]

	//fmt.Println("<->", a)
	fmt.Println("<->", s)

	a_to_left, is_ok := rotate(s, shift)
	if ( is_ok == true ) {
		fmt.Printf("<%d] %v\n", shift, a_to_left)
	} else {
		fmt.Printf("error1\n")
	}

	//fmt.Println("<->", a)
	fmt.Println("<->", s)

	a_to_right, is_ok := rotate(s, -shift)
	if ( is_ok == true ) {
		fmt.Printf("[%d> %v\n", shift, a_to_right)
	} else {
		fmt.Printf("error2\n")
	}

	//fmt.Println("<->", a)
}

/*--------------------------------------------------------------*/
/* rotate()ってテキストにあったっけ？							*/
/*--------------------------------------------------------------*/
func rotate(s []int, n int) ([]int, bool) {
	if ( n == 0 ) { return s, true }

	if ( (n >= len(s)) || (-n >= len(s)) ) {
		return nil, false
	}

	if ( n < 0 ) {
		n = (len(s) + n)
	}

	return append(s[n:], s[:n]...), true
}

