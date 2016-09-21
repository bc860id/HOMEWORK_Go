
// 第4章 練習問題4.2
package main
import (
	//"os"
	"fmt"
	"flag"
	"crypto/sha256"
	//"crypto/sha384"
	"crypto/sha512"
)

var sha = flag.String("s", "256", "digest algo")

func main() {
	flag.Parse()

	//fmt.Printf("s:%s\n", *sha)
	switch *sha {
	 case "256"	:
		fmt.Printf("%X\n", sha256.Sum256([]byte((flag.Args())[0])))

	 case "384"	:
		fmt.Printf("%X\n", sha512.Sum384([]byte((flag.Args())[0])))

	 case "512"	:
		fmt.Printf("%X\n", sha512.Sum512([]byte((flag.Args())[0])))

	 default	:
		fmt.Printf("digest algo SHA%s is NOT supported!\n", *sha)
	}
}



