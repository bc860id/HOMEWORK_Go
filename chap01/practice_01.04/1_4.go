
// 第1章 練習問題1.4
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2-modif(1): %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	searchFiles(os.Args[1:], counts, filenames)

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t:%s\n", n, line, filenames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func searchFiles(files []string, counts map[string]int, fnames map[string]string) {
	// 2つ以上出現した行について再度各ファイル中に存在するか調べる.
	for line, n := range counts {
		if n > 1 {
			for _, arg := range files {
				f, err := os.Open(arg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2-modif(2): %v\n", err)
					continue
				}
				input := bufio.NewScanner(f)
				// 1つのファイル中で1つでも見つかればファイルを記録して抜ける.
				for input.Scan() {
					if line == input.Text() {
						fnames[line] = fnames[line] + arg + " "
						break
					}
				}
				f.Close()
			}
		}
	}
}





