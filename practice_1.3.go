
// 第1章 練習問題1.3
package main

import (
	"fmt"
	"os"
	//"strconv"
	"time"
	"strings"
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
	tm1_sec := time.Since(tm1).Seconds()

	// 効率バージョン(Ver.2)
	tm2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], "+"))
	tm2_sec := time.Since(tm2).Seconds()

	if ( tm1_sec > tm2_sec ) {
		fmt.Printf("Ver.1 is %f second(s) later than Ver.2\n", (tm1_sec - tm2_sec))
	} else {
		fmt.Printf("Ver.2 is %f second(s) later than Ver.1\n", (tm2_sec - tm1_sec))
	}
	fmt.Printf("tm1:%f tm2:%f\n", tm1_sec, tm2_sec)
}
