package util

import (
	"fmt"
)

// powers of 1024 and their suffixes
var suffixes = map[uint8]string{
	0: "B",
	1: "KB",
	2: "MB",
	3: "GB",
	4: "TB",
	5: "PB",
	6: "EB",
}

func formatRecursive(size float64, powerOfThousand uint8) string {
	if size > 1024 {
		return formatRecursive(size/1024, powerOfThousand+1)
	}
	return fmt.Sprintf("%.2f"+suffixes[powerOfThousand], size)
}

// FormatBytes formats a number of bytes into a string with the appropriate suffix.
func FormatBytes(bytes int64) string {
	return formatRecursive(float64(bytes), 0)
}
