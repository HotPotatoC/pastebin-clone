package logic_test

import (
	"testing"

	"github.com/HotPotatoC/pastebin-clone/logic"
)

func TestBase62(t *testing.T) {
	tc := []struct {
		src []byte
	}{
		{[]byte("a")},
		{[]byte("Hello World!")},
		{[]byte("fffffffffffffffffffffffff")},
		{[]byte("6c30bdbe-82a1-4529-97f7-b306eff20c34127.179.52.136Hello, World!")},
	}

	for _, tt := range tc {
		t.Run(string(tt.src), func(t *testing.T) {
			encoded := logic.EncodeBase62(tt.src)
			decoded, err := logic.DecodeBase62(encoded)
			if err != nil {
				t.Errorf("Decode: Expected nil, got %s", err)
			}

			if string(decoded) != string(tt.src) {
				t.Errorf("Decode: Expected %s, got %s", tt.src, decoded)
			}

			t.Logf("Encoded: %s", encoded)
			t.Logf("Decoded: %s", decoded)
		})
	}
}
