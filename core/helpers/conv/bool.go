package conv

import "strings"

func ParseBool(val interface{}) bool {
	switch val.(type) {
	case string:
		sVal := strings.ToLower(val.(string))
		switch sVal {
		case "yes":
			return true
		case "no":
			return false
		case "true":
			return true
		case "false":
			return false
		case "0":
			return false
		case "1":
			return true
		default:
			return false
		}
	default:
		switch val {
		case true:
			return true
		case false:
			return false
		case 1:
			return true
		case 0:
			return false
		default:
			return false
		}
	}
}
