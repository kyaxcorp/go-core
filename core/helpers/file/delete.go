package file

import "os"

func Remove(filePath string) (bool, error) {
	return Delete(filePath)
}

func Delete(filePath string) (bool, error) {
	_err := os.Remove(filePath)
	if _err != nil {
		return false, _err
	}
	return true, nil
}
