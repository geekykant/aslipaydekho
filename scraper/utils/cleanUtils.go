package utils

import (
	"strings"
)

func BasicCleanCompensationPostContent(data *string) {
	*data = strings.ReplaceAll(*data, "\\n", "\n")
	*data = strings.ReplaceAll(*data, "\\", "")
	*data = strings.ReplaceAll(*data, "**", "")
	*data = strings.ReplaceAll(*data, "```", "")
}
