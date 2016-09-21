
// 第3章 練習問題3.13
package main
import (
	"fmt"
	//"math/big"
)

/*--------------------------------------------------------------*/
/* K	キロ	Kilo	1000^1									*/
/* M	メガ	Mega	1000^2									*/
/* G	ギガ	Giga	1000^3									*/
/* T	テラ	Tera	1000^4									*/
/* P	ペタ	Peta	1000^5									*/
/* E	エクサ	Exa		1000^6									*/
/* Z	ゼタ	Zetta	1000^7									*/
/* Y	ヨタ	Yotta	1000^8									*/
/*--------------------------------------------------------------*/
const (
	B	= 1
	KB	= B * K
	MB	= KB * K
	GB	= MB * K
	TB	= GB * K
	PB	= TB * K
	EB	= PB * K
	ZB	= EB * K
	YB	= ZB * K

	K	= 1000
)

func main() {
	fmt.Printf(" B=%d\n", B)
	fmt.Printf("KB=%d\n", KB)
	fmt.Printf("MB=%d\n", MB)
	fmt.Printf("GB=%d\n", GB)
	fmt.Printf("TB=GBx%d\n", (TB / GB))
	fmt.Printf("PB=TBx%d\n", (PB / TB))
	fmt.Printf("EB=PBx%d\n", (EB / PB))
	fmt.Printf("ZB=EBx%d\n", (ZB / EB))
	fmt.Printf("YB=ZBx%d\n", (YB / ZB))
}



