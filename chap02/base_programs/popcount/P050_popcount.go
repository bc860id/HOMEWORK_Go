
// base program written on P50 gopl.io/ch2/popcount
package popcount

var pc[256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i >> 1] + byte(i % 2)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x >> (0 * 8))] +
			   pc[byte(x >> (1 * 8))] +
			   pc[byte(x >> (2 * 8))] +
			   pc[byte(x >> (3 * 8))] +
			   pc[byte(x >> (4 * 8))] +
			   pc[byte(x >> (5 * 8))] +
			   pc[byte(x >> (6 * 8))] +
			   pc[byte(x >> (7 * 8))])
}


