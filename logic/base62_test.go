package logic_test

import (
	"testing"

	"github.com/HotPotatoC/pastebin-clone/logic"
)

func TestBase62(t *testing.T) {
	src := []byte("Hello World!")
	encoded := logic.EncodeBase62(src)
	decoded := logic.DecodeBase62(encoded)

	if string(decoded) != string(src) {
		t.Errorf("Expected %s, got %s", src, decoded)
	}

	t.Logf("Encoded: %s", encoded)
	t.Logf("Decoded: %s", decoded)
}
