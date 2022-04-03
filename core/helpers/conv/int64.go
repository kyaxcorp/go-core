package conv

import "strconv"

func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Int64SliceToStrSlice(val []int64) []string {
	var r []string
	if len(val) == 0 {
		return r
	}
	for _, v := range val {
		r = append(r, Int64ToStr(v))
	}
	return r
}
