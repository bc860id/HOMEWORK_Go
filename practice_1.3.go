
// ■第1章 練習問題1.3
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// 非効率バージョン(Ver.1)
	tm1 := time.Now()
	var s1, sep1 string
	for i := 1; i < len(os.Args); i++ {
		s1 += sep1 + os.Args[i]
		sep1 = " "
	}
	fmt.Println(s1)
	tm1 = time.Since(tm1).Seconds()

	// 効率バージョン(Ver.2)
	tm2 := time.Now()
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)
	tm2 = time.Since(tm2).Seconds()

	if ( tm1 > tm2 ) {
		fmt.Println("Ver.1 is " + strconv.Itoa(tm1 - tm2) + "second(s) later than Ver.2")
	}
	else {
		fmt.Println("Ver.2 is " + strconv.Itoa(tm2 - tm1) + "second(s) later than Ver.1")
	}
}
