package conv

import "strconv"

func IntToStr(val int) string {
	return strconv.Itoa(val)
}

func IntSliceToStrSlice(val []int) []string {
	var r []string
	if len(val) == 0 {
		return r
	}
	for _, v := range val {
		r = append(r, IntToStr(v))
	}
	return r
}
