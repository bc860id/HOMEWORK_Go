
// 第4章 練習問題4.6
package main
import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

var sample_str1 = []byte("a b  c   あ 　i　　う")
var sample_str2 = []byte("   a b  c   あ 　i　　う")
var sample_str3 = []byte("　   　　a b  c   あ 　i　　う")
var sample_str4 = []byte("　   　　a b  c   あ 　i　　う ")
var sample_str5 = []byte("　   　　a b  c   あ 　i　　う  ")
var sample_str6 = []byte("　   　　a b  c   あ 　i　　う 　")

func main() {
	fmt.Println(string(compspace(sample_str1)))
	fmt.Println(string(compspace(sample_str2)))
	fmt.Println(string(compspace(sample_str3)))
	fmt.Println(string(compspace(sample_str4)))
	fmt.Println(string(compspace(sample_str5)))
	fmt.Println(string(compspace(sample_str6)))
}

func compspace(bs []byte) []byte {
	var i, j, l int
	_, size_space_ascii := utf8.DecodeRuneInString(" ")

	for i = 0; i < len(bs); {
		r1, size := utf8.DecodeRune(bs[i:])
		if ( unicode.IsSpace(r1) == false ) {
			i += size
			continue
		}

		/*------------------------------------------------------*/
		/* 全角でも半角でもASCIIスペースに置き換え。			*/
		/*------------------------------------------------------*/
		bs_temp := append(bs[:i], ' ')

		/*------------------------------------------------------*/
		/* スペースではないところを探す。						*/
		/*------------------------------------------------------*/
		for j, l = (i + size), len(bs); j < l; {
			r2, size_space := utf8.DecodeRune(bs[j:])
			if ( unicode.IsSpace(r2) == false ) {
				bs = append(bs_temp, bs[j:]...)
				break
			}
			j += size_space
		}

		/*------------------------------------------------------*/
		/* for文の前で付け足した' 'の分をiに加算。				*/
		/*------------------------------------------------------*/
		i += size_space_ascii

		/*------------------------------------------------------*/
		/* 後ろはスペースの連続。								*/
		/* len(bs)がループ開始時より小さくなっているので		*/
		/* この段階のlen(bs)とjを比較するのはダメ。				*/
		/*------------------------------------------------------*/
		if ( j >= l ) {
			break
		}
	}
	return bs[:i]
}

