package str

func Includes(slice []string, item string) bool {
	if slice == nil || item == "" {
		return false
	}
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}
	return false
}
