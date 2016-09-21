
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
	if ( len(os.Args) < 3 ) {
		return
	}

	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	/*
	c1 := [...]byte{0: 0xFF, 31: 0x0E}
	c2 := [...]byte{0: 0xFF, 31: 0x0C}
	*/

	for i := 0; i < len(c1); i++ {
		b := (c1[i] ^ c2[i])
		for ; b > 0; bit_num++ {
			b = (b & (b - 1))
		}
	}

	fmt.Printf("total diff bits:%d\n", bit_num)
}



