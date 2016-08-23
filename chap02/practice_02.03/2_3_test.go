
// 第2章 練習問題2.3
package popcount

import (
	"./popcount"
	"testing"
)

const (
	samplevalue		uint64 = 65531
)

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount1(samplevalue)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount2(samplevalue)
	}
}
