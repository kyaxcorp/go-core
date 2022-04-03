package conv

import "strconv"

func UInt64ToStr(val uint64) string {
	// TODO: check if base is correct here!
	return strconv.FormatUint(val, 10)
}

func UInt64SliceToStrSlice(val []uint64) []string {
	var r []string
	if len(val) == 0 {
		return r
	}
	for _, v := range val {
		r = append(r, UInt64ToStr(v))
	}
	return r
}
