package array

// array and values
func InArrayInt(array []int, value int) bool {
	found := false
	for _, val := range array {
		if val == value {
			found = true
			break
		}
	}
	if found {
		return true
	}
	return false
}

func ConvUint64ValuesToMapKey(array []uint64) map[uint64]int {
	m := make(map[uint64]int)
	for order, val := range array {
		// There may be repeating cases, but they'll be overwritten
		m[val] = order
	}
	return m
}

func ConvStringValuesToMapKey(array []string) map[string]int {
	m := make(map[string]int)
	for order, val := range array {
		// There may be repeating cases, but they'll be overwritten
		m[val] = order
	}
	return m
}

func InArrayUint64(array []uint64, value uint64) bool {
	found := false
	for _, val := range array {
		if val == value {
			found = true
			break
		}
	}
	if found {
		return true
	}
	return false
}

func InArrayString(array []string, value string) bool {
	found := false
	for _, val := range array {
		if val == value {
			found = true
			break
		}
	}
	if found {
		return true
	}
	return false
}

func InArray(array []interface{}, value interface{}) bool {
	found := false
	for _, val := range array {
		if val == value {
			found = true
			break
		}
	}
	if found {
		return true
	}
	return false
}

func IndexExists(index int, array []interface{}) bool {
	// There is no sparse slices in Go, so you could simply check the length:
	// https://stackoverflow.com/questions/27252152/how-to-check-if-a-slice-has-a-given-index-in-go
	if len(array) > index {
		return true
	}
	return false
}

func IndexExistsString(index int, array []string) bool {
	// There is no sparse slices in Go, so you could simply check the length:
	// https://stackoverflow.com/questions/27252152/how-to-check-if-a-slice-has-a-given-index-in-go
	if len(array) > index {
		return true
	}
	return false
}
