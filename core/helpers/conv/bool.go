package conv

func ParseBool(val interface{}) bool {
	switch val {
	case true:
		return true
	case false:
		return false
	case 1:
		return true
	case 0:
		return false
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
}
