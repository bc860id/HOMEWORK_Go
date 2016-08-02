
// 第1章 練習問題1.2
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, arg := range os.Args[0:] {
		fmt.Println(strconv.Itoa(idx) + " " + arg)
	}
}
