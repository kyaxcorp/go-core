package conv

import "strconv"

func StrToBytes(val string) []byte {
	return []byte(val)
}

func StrToUint64(val string) uint64 {
	number, _ := strconv.ParseUint(val, 10, 64)
	return number
}

func StrToInt64(val string) int64 {
	number, _ := strconv.ParseInt(val, 10, 64)
	return number
}

func StrToInt32(val string) int32 {
	number, _ := strconv.ParseInt(val, 10, 32)
	return int32(number)
}

func StrToInt(val string) int {
	number, _ := strconv.Atoi(val)
	return number
}

func StrToUint64E(val string) (uint64, error) {
	number, _err := strconv.ParseUint(val, 10, 64)
	return number, _err
}

func StrToFloat64(val string) float64 {
	number, _ := strconv.ParseFloat(val, 64)
	return number
}

func StrToFloat64Err(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

func StrToFloat32(val string) float32 {
	number, _ := strconv.ParseFloat(val, 32)
	return float32(number)
}

func StrToFloat32Err(val string) (float32, error) {
	number, _err := strconv.ParseFloat(val, 32)
	return float32(number), _err
}
