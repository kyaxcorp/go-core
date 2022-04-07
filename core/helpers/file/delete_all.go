package file

import "os"

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func RemoveRecursive(path string) error {
	return RemoveAll(path)
}

func DeleteAll(path string) error {
	return RemoveAll(path)
}

func DeleteRecursive(path string) error {
	return RemoveAll(path)
}
