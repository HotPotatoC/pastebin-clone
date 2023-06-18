package logic

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(b []byte) (string, error) {
	h := sha256.New()
	_, err := h.Write(b)
	if err != nil {
		return "", err
	}

	src := h.Sum(nil)

	return hex.EncodeToString(src), nil
}
