
// additional program for application of package tempconv
package main
import (
	"./tempconv"
	"fmt"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.FreezingC))
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
