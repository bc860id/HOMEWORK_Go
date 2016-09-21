
// 第3章 練習問題3.10
package main
import (
	"os"
	"fmt"
	"bytes"
)

func main() {
	if (len(os.Args) < 2) {
		return
	}
	addcomma := comma(os.Args[1])
	fmt.Println(addcomma)
}

func comma(s string) string {
	var buf bytes.Buffer
	c := (len(s) % 3)
	for i := 0; i < len(s); i++ {
		if (c == 0) {
			if (i > 0) {
				buf.WriteByte(',')
			}
			c = 3
		}
		c--
		buf.WriteByte(s[i])
	}
	return buf.String()
}


