package json

import "github.com/tidwall/gjson"

func IsKeyExists(json string, keyPath string) bool {
	value := gjson.Get(json, keyPath)
	if !value.Exists() {
		return false
	} else {
		return true
	}
}
