package str

import "strings"

func Explode(s string, sep string) []string {
	return strings.Split(s, sep)
}
