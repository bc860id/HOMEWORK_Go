
// additional progaram for application of package popcount
package main

import (
	"fmt"
	"os"
	"strconv"
	"./popcount"
)

func main() {
	for _, valstr := range os.Args[1:] {
		val, _ := strconv.ParseUint(valstr, 10, 64)
		fmt.Printf("%08Xh:%d\n", val, popcount.PopCount(val))
	}
}

