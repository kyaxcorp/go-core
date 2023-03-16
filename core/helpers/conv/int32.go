package conv

import "strconv"

func Int32ToStr(val int32) string {
	return strconv.FormatInt(int64(val), 10)
}

func Int32SliceToStrSlice(val []int32) []string {
	var r []string
	if len(val) == 0 {
		return r
	}
	for _, v := range val {
		r = append(r, Int32ToStr(v))
	}
	return r
}
