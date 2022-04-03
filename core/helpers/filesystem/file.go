package filesystem

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func FilePutContentsMode(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

func FilePutContents(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0751)
}

func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}

func FilterPath(path string) string {
	p := strings.ReplaceAll(path, `\\`, DirSeparator())
	p = strings.ReplaceAll(p, "//", DirSeparator())
	p = strings.ReplaceAll(p, "/", DirSeparator())
	p = strings.ReplaceAll(p, `\`, DirSeparator())
	return p
}

func Copy(src, dst string) (int64, error) {
	src = FilterPath(src)
	dst = FilterPath(dst)

	//log.Println(src)
	//log.Println(dst)
	sourceFileStat, _err := os.Stat(src)
	if _err != nil {
		return 0, _err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, _err := os.Open(src)
	if _err != nil {
		return 0, _err
	}
	defer source.Close()

	destination, _err := os.Create(dst)
	if _err != nil {
		return 0, _err
	}
	defer destination.Close()
	nBytes, _err := io.Copy(destination, source)
	return nBytes, _err
}

func Unlink(path string) bool {
	// TODO: we should also delete folders and recursive paths!

	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if !FileExists(path) {
		return true
	}
	return false
}

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func RemoveRecursive(path string) error {
	return RemoveAll(path)
}

func Remove(path string) bool {
	return Unlink(path)
}

func Delete(path string) bool {
	return Unlink(path)
}

func DeleteAll(path string) error {
	return RemoveAll(path)
}

func DeleteRecursive(path string) error {
	return RemoveAll(path)
}

func getPathType(path string) uint8 {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		//fmt.Println("directory")
		return 1
	case mode.IsRegular():
		// do file stuff
		//fmt.Println("file")
		return 2
	}
	return 0
}

func IsFile(path string) bool {
	if getPathType(path) == 2 {
		return true
	}
	return false
}

func IsDir(path string) bool {
	if getPathType(path) == 1 {
		return true
	}
	return false
}
