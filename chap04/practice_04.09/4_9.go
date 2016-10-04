
// 第4章 練習問題4.9
package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var each = make(map[string]int)
var total uint64 = 0

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open error:%v\n", err)
		os.Exit(1)
	}
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanWords)
	wordfreq(s)
	defer file.Close()
}

func wordfreq(words *bufio.Scanner) {
	for words.Scan() {
		if err := words.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "error:%v %s\n", err, words.Text())
			return
		}
		key := strings.Trim(words.Text(), ".,:;")
		//each[words.Text()]++
		each[key]++
		total++
	}

	fmt.Printf("total:%d words\n\n", total)
	for w, c := range each {
		rate := ((float64(c) / float64(total)) * 100.0)
		fmt.Printf("%s\t\t%f%% (%d / %d)\n", w, rate, c, total)
	}
}


