
// 第3章 練習問題3.10
package main
import (
	"os"
	"fmt"
	"bytes"
	"strconv"
)

func main() {
	if (len(os.Args) < 2) {
		return
	}
	addcomma := comma(os.Args[1])
	fmt.Println(addcomma)
}

func comma(s string) string {
	var buf, buf_decpart bytes.Buffer
	val, err := strconv.ParseFloat(s, 64)
	if ( err != nil ) {
		fmt.Fprintf(os.Stderr, "input val error:%s %v\n", s, err)
		return ""
	}

	val_int := int(val)
	val_decpart := float64(val - float64(val_int))
	if ( val_decpart < 0 ) {
		val_decpart *= (-1)
	}

	str_val_int := strconv.Itoa(val_int)

	c := (len(str_val_int) % 3)
	for i := 0; i < len(str_val_int); i++ {
		if (c == 0) {
			if (i > 0) {
				buf.WriteByte(',')
			}
			c = 3
		}
		c--
		buf.WriteByte(str_val_int[i])
	}

	if ( val_decpart == 0 ) {
		return buf.String()
	}

	//str_val_decpart := strconv.Ftoa(val_decpart)
	fmt.Fprintf(&buf_decpart, "%f", val_decpart)
	str_val_decpart := buf_decpart.String()

	if ( (val_int == 0) && (val < 0) ) {
		return ("-" + str_val_decpart[0:])
	}
	return (buf.String() + str_val_decpart[1:])
}


