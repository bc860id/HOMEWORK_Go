
package popcount

import (
	"testing"
)

var pc[256]byte

const (
	samplevalue		uint64 = 65531
)

func init() {
	for i := range pc {
		pc[i] = pc[i >> 1] + byte(i % 2)
	}
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(samplevalue)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(samplevalue)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(samplevalue)
	}
}

func PopCount1(x uint64) int {
	return int(pc[byte(x >> (0 * 8))] +
			   pc[byte(x >> (1 * 8))] +
			   pc[byte(x >> (2 * 8))] +
			   pc[byte(x >> (3 * 8))] +
			   pc[byte(x >> (4 * 8))] +
			   pc[byte(x >> (5 * 8))] +
			   pc[byte(x >> (6 * 8))] +
			   pc[byte(x >> (7 * 8))])
}

func PopCount2(x uint64) int {
	var count byte
	var i uint
	for i = 0; i < 8; i++ {
		count += pc[byte(x >> (i * 8))]
	}
	return int(count)
}

func PopCount3(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

