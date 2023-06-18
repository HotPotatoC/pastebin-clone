package logic_test

import (
	"bytes"
	"testing"

	"github.com/HotPotatoC/pastebin-clone/logic"
)

func TestCompressDecompress(t *testing.T) {
	original := []byte("Hello, World!")
	compressed, err := logic.Compress(original)
	if err != nil {
		t.Error(err)
	}

	decompressed, err := logic.Decompress(compressed)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(decompressed, original) {
		t.Error("decompressed and original should be the same. got:", decompressed)
	}
}
