
// 第3章 練習問題4.1
package main
import (
	"os"
	"fmt"
	"crypto/sha256"
)

func main() {
	var bit_num int

	//fmt.Println(len(os.Args))
	if ( len(os.Args) < 2 ) {
		return
	}

	c := sha256.Sum256([]byte(os.Args[1]))
	//c := [...]byte{0: 0xFF, 31: 0x0E}

	for i := 0; i < len(c); i++ {
		b := c[i]
		for ; b > 0; bit_num++ {
			b = (b & (b - 1))
		}
	}

	fmt.Printf("total bits:%d\n", bit_num)
}



