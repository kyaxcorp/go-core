package _int

import (
	"regexp"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func IsNumber(val string) bool {
	return digitCheck.MatchString(val)
}
