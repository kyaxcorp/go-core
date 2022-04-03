package json

import "github.com/tidwall/gjson"

func IsValid(json string) bool {
	if !gjson.Valid(json) {
		return false
	}
	return true
}
