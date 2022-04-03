package conv

import "strconv"

func UInt16ToStr(val uint16) string {
	return strconv.Itoa(int(val))
}

func UInt16SliceToStrSlice(val []uint16) []string {
	var r []string
	if len(val) == 0 {
		return r
	}
	for _, v := range val {
		r = append(r, UInt16ToStr(v))
	}
	return r
}
