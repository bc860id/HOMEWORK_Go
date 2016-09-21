
// base program written on P81 gopl.io/ch3/comma
package main
import (
	"os"
	"fmt"
)

func main() {
	if (len(os.Args) < 2) {
		return
	}
	addcomma := comma(os.Args[1])
	fmt.Println(addcomma)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n - 3]) + "," + s[n - 3:]
}

