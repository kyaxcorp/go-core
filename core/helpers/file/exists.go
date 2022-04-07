package file

import "os"

func Exists(filename string) bool {
	if _, _err := os.Stat(filename); _err == nil {
		return true
	} else {
		return false
	}
}

func ExistsErr(filename string) (bool, error) {
	if _, _err := os.Stat(filename); _err == nil {
		return true, nil
	} else {
		return false, _err
	}
}
