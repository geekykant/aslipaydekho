package utils

import (
	"strings"
)

func CleanCompensationPostContent(data *string) {
	*data = strings.ReplaceAll(*data, "\\n", "\n")
	*data = strings.ReplaceAll(*data, "\\", "")
	*data = strings.ReplaceAll(*data, "**", "")
	*data = strings.ReplaceAll(*data, "```", "")
}
