
// 第4章 練習問題4.5
package main
import (
	"os"
	"fmt"
)

var a = [...]int{0, 1, 2, 3, 4, 5}

func main() {
	if ( len(os.Args) < 2 ) {
		return
	}
	s := unrepeat(os.Args[1:])
	fmt.Println(s)
}

func unrepeat(s []string) []string {
	var i, j int
	if ( len(s) < 2 ) { return s }

	for i = 0; i < (len(s) - 1); i++ {
		for j = (i + 1); j < len(s); j++ {
			if ( s[i] == s[j] ) { continue }
			s = append(s[:i + 1], s[j:]...)
			break
		}

		/*------------------------------------------------------*/
		/* 残り全てがs[i]と同じ単語の場合はs[i]までで終わり。	*/
		/*------------------------------------------------------*/
		if ( j >= len(s) ) {
			//fmt.Printf("1: i:%d j:%d\n", i, j)
			break
		}
	}
	//fmt.Printf("2: i:%d j:%d\n", i, j)
	return s[:(i + 1)]
}

