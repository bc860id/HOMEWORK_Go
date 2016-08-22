
// 第2章 練習問題2.2
package main

import (
	"./distanceconv"
	"os"
	"fmt"
	"flag"
	"bufio"
	"strconv"
)

var inch2mm = flag.Bool("inch", false, "convert inch to mm")
var mm2inch = flag.Bool("mm", false, "convert mm to inch")

func main() {
	var val float64
	var err error

	flag.Parse()

	// inch2mmとmm2inchが共に真か共に偽の場合は変換方法不明.
	if ( *inch2mm == *mm2inch ) {
		fmt.Println("select one option only inch or mm")
		return
	}

	// フラグを除いた引数の総数を検査.
	if ( len(flag.Args()) < 1 ) {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		val, err = strconv.ParseFloat(scanner.Text(), 64)
	} else {
		val, err = strconv.ParseFloat((flag.Args())[0], 64)
	}

	if ( err != nil ) {
		fmt.Println("error:%v", err)
		return
	}

	if ( *inch2mm == true ) {
		fmt.Println(distanceconv.InchToMm(distanceconv.Inch(val)))
	} else {
		fmt.Println(distanceconv.MmToInch(distanceconv.Millimeter(val)))
	}
}
