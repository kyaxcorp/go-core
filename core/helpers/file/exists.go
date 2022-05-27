package file

import (
	"errors"
	"os"
)

func Exists(filename string) bool {
	if _, _err := os.Stat(filename); _err == nil {
		return true
	} else {
		return false
	}
}

func ExistsErr(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
