
// 第2章 練習問題2.1
package main

import (
	"./tempconv"
	"fmt"
)

func main() {
	fmt.Println(tempconv.CToK(0))
	fmt.Println(tempconv.KToC(0))
	fmt.Println(tempconv.CToK(-1))
	fmt.Println(tempconv.KToC(1))
	fmt.Println(tempconv.CToK(1))
	fmt.Println(tempconv.KToC(-1))
}


