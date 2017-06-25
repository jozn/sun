package helper

import (
	"fmt"
	"testing"
)

func BenchmarkNextUUid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextRowsSeqId()
	}
}

func BenchmarkTimeNowMs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeNowMs()
	}
}

func BenchmarkParsing(b *testing.B) {
	gen := defGen
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d%02d%04d", gen.timeMs, gen.serverId, gen.counter)
	}
}

func BenchmarkRandomSeqUid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomSeqUid()
	}
}
