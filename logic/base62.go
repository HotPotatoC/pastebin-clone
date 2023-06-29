package logic

import (
	"github.com/jxskiss/base62"
)

func EncodeBase62(src []byte) string {
	return base62.EncodeToString(src)
}

func DecodeBase62(src string) ([]byte, error) {
	return base62.DecodeString(src)
}
