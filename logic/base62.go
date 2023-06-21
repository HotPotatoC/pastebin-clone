package logic

import "math/big"

func EncodeBase62(src []byte) string {
	var i big.Int
	i.SetBytes(src)
	return i.Text(62)
}

func DecodeBase62(src string) []byte {
	var i big.Int
	i.SetString(src, 62)
	return i.Bytes()
}
