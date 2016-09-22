
// 第4章 練習問題4.7
package main
import (
	"fmt"
	"unicode/utf8"
)

var sample_str1 = []byte("12")
var sample_str2 = []byte("A雲")
var sample_str3 = []byte("123７８９")
var sample_str4 = []byte("123④７８９")

func main() {
	fmt.Println(string(reverse(sample_str1)))
	fmt.Println(string(reverse(sample_str2)))
	fmt.Println(string(reverse(sample_str3)))
	fmt.Println(string(reverse(sample_str4)))
}

func reverse(bs []byte) []byte {
	var idx_bwd int

	/*
	fmt.Printf("\n\nlength:%d\n", len(bs))
	fmt.Printf("----------->\n")
	*/

	idx_bwd = (len(bs) - 1)
	idx_bwd_prev := idx_bwd
	for idx_fwd := 0; /* NOP */; {
		for /* NOP */; 0 <= idx_bwd; idx_bwd-- {
			if ( utf8.RuneStart(bs[idx_bwd]) == true ) { break }
		}
		if ( idx_fwd > idx_bwd ) { break }
		_, size_f := utf8.DecodeRune(bs[idx_fwd:])
		_, size_b := utf8.DecodeRune(bs[idx_bwd:])

		/*
		fmt.Printf("idx_f:%d %d idx_b:%d %d\n", idx_fwd, size_f, idx_bwd, size_b)
		fmt.Println(string(bs[idx_fwd:idx_fwd+size_f]),
			string(bs[idx_bwd:idx_bwd+size_b]))
		*/

		bs = exchange(bs, idx_fwd, idx_bwd)

		//fmt.Printf("length:%d\n", len(bs))

		idx_fwd += size_b	/* 後ろから持ってきた分シフト。	*/
		idx_bwd = (idx_bwd_prev - size_f)
		idx_bwd_prev = idx_bwd
	}

	//fmt.Printf("<-----------\n")
	return bs
}

func exchange(bs []byte, i, j int) []byte {
	var idx_h, idx_t int
	var bs_cp_h, bs_cp_t []byte

	if ( i == j ) { return bs }

	if ( i < j ) {
		idx_h = i
		idx_t = j
	} else {
		idx_h = j
		idx_t = i
	}

	_, size_h := utf8.DecodeRune(bs[idx_h:])
	_, size_t := utf8.DecodeRune(bs[idx_t:])

	bs_cp_h = make([]byte, len(bs[:idx_h]), len(bs))
	copy(bs_cp_h, bs[:idx_h])
	bs_temp_h := append(bs_cp_h, bs[idx_t:(idx_t + size_t)]...)

	bs_cp_t = make([]byte, len(bs[idx_h:(idx_h + size_h)]), len(bs))
	copy(bs_cp_t, bs[idx_h:(idx_h + size_h)])
	bs_temp_t := append(bs_cp_t, bs[(idx_t + size_t):]...)

	bs_temp := append(bs_temp_h, bs[(idx_h + size_h):idx_t]...)
	bs_temp = append(bs_temp, bs_temp_t...)

	/*
	bs_cp_h = make([]byte, len(bs[:idx_h]), len(bs))
	copy(bs_cp_h, bs[:idx_h])
	fmt.Println("1@", string(bs_cp_h))
	bs_temp_h := append(bs_cp_h, bs[idx_t:(idx_t + size_t)]...)
	fmt.Println("1@", string(bs_temp_h))

	fmt.Println("2@", string(bs[idx_h:(idx_h + size_h)]))
	bs_cp_t = make([]byte, len(bs[idx_h:(idx_h + size_h)]), len(bs))
	copy(bs_cp_t, bs[idx_h:(idx_h + size_h)])
	fmt.Println("2@", string(bs_cp_t))
	bs_temp_t := append(bs_cp_t, bs[(idx_t + size_t):]...)
	fmt.Println("2@", string(bs_temp_t))

	bs_temp := append(bs_temp_h, bs[(idx_h + size_h):idx_t]...)
	fmt.Println("3@", string(bs_temp))
	bs_temp = append(bs_temp, bs_temp_t...)
	fmt.Println("4@", string(bs_temp))
	*/

	return bs_temp
}

