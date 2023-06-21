package logic_test

import (
	"testing"

	"github.com/HotPotatoC/pastebin-clone/logic"
)

func TestByteCount(t *testing.T) {
	tc := []struct {
		name     string
		input    uint64
		expected string
	}{
		{
			name:     "1 byte",
			input:    1,
			expected: "1 B",
		},
		{
			name:     "1 kilobyte",
			input:    1 << (10 * 1),
			expected: "1.00 KB",
		},
		{
			name:     "1 megabyte",
			input:    1 << (10 * 2),
			expected: "1.00 MB",
		},
		{
			name:     "1 gigabyte",
			input:    1 << (10 * 3),
			expected: "1.00 GB",
		},
		{
			name:     "Hello world",
			input:    uint64(len(`Hello world`)),
			expected: "11 B",
		},
		{
			name:     "25 gigabytes",
			input:    25 << (10 * 3),
			expected: "25.00 GB",
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := logic.ByteCount(tt.input)
			if got != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
