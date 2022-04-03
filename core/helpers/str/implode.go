package str

import (
	"strings"
)

func Implode(slice []string, sep string) string {
	return strings.Join(slice, sep)
}
