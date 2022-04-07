package file

import (
	"fmt"
	"os"
)

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
