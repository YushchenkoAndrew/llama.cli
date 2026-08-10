package lib

import (
	"fmt"

	"github.com/YushchenkoAndrew/llama-cli/src/types"
)

func GetSize(m *types.ModelMeta) string {
	if m == nil {
		return "-"
	}

	if float64(m.Size) >= 1e9 {
		return fmt.Sprintf("%.1fG", float64(m.Size)/1e9)
	}

	if float64(m.Size) >= 1e6 {
		return fmt.Sprintf("%.1fM", float64(m.Size)/1e6)
	}

	if float64(m.Size) >= 1e3 {
		return fmt.Sprintf("%.1fk", float64(m.Size)/1e3)
	}

	return fmt.Sprintf("%d", m.Size)
}
