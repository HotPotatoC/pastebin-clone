package logic_test

import (
	"testing"

	"github.com/HotPotatoC/pastebin-clone/logic"
)

func TestHash(t *testing.T) {
	x, err := logic.Hash([]byte("Hello, world!"))
	if err != nil {
		t.Error(err)
	}

	y, err := logic.Hash([]byte("Hello, world!"))
	if err != nil {
		t.Error(err)
	}

	if x != y {
		t.Error("x and y should be the same. got:", y)
	}
}
