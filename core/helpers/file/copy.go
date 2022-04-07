package file

import (
	"fmt"
	"io"
	"os"
)

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
