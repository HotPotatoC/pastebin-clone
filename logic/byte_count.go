package logic

import "fmt"

// ByteCount convert sizes in bytes into a human-readable string
// src: https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func ByteCount(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
