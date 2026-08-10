package lib

import (
	"fmt"
)

func PrettifySize(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%dB", bytes)
	}

	// Calculate the correct suffix index
	// 1 = K, 2 = M, 3 = G, 4 = T, etc.
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	suffixes := []string{"K", "M", "G", "T", "P", "E"}
	value := float64(bytes) / float64(div)

	// Formats to 1 decimal place. Use "%.0f" if you prefer no decimals for whole numbers.
	return fmt.Sprintf("%.1f%s", value, suffixes[exp])
}
