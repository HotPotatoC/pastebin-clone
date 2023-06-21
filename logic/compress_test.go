package logic_test

import (
	"bytes"
	"testing"

	"github.com/HotPotatoC/pastebin-clone/logic"
)

func TestCompressDecompress(t *testing.T) {
	original := []byte(`package logic

	import (
		"bytes"
		"compress/zlib"
		"io"
	)

	func Compress(b []byte) ([]byte, error) {
		var buf bytes.Buffer

		w := zlib.NewWriter(&buf)

		_, err := w.Write(b)
		if err != nil {
			return nil, err
		}

		err = w.Close()
		if err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	func Decompress(b []byte) ([]byte, error) {
		buf := bytes.NewReader(b)
		r, err := zlib.NewReader(buf)
		if err != nil {
			return []byte{}, err
		}

		res, err := io.ReadAll(r)
		if err != nil {
			return []byte{}, err
		}

		return res, nil
	}
	`)
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

	t.Log("compressed size:", len(compressed))
	t.Log("original size:", len(original))
}
