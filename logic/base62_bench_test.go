package logic_test

import (
	"crypto/rand"
	"testing"

	"github.com/HotPotatoC/pastebin-clone/logic"
)

var randBytes = make([]byte, 16)

func init() {
	if _, err := rand.Read(randBytes); err != nil {
		panic(err)
	}
}

func BenchmarkBase62Encode(b *testing.B) {
	b.StopTimer()
	b.Logf("Benchmarking EncodeBase62 with %x", randBytes)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logic.EncodeBase62(randBytes)
	}
}
