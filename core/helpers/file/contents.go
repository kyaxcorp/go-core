package file

import (
	"os"
)

func PutContentsMode(filename string, data string, mode os.FileMode) error {
	//return ioutil.WriteFile(filename, []byte(data), mode)
	return os.WriteFile(filename, []byte(data), mode)
}

func PutContents(filename string, data string) error {
	//return ioutil.WriteFile(filename, []byte(data), 0751)
	return os.WriteFile(filename, []byte(data), 0751)
}

func GetContents(filename string) (string, error) {
	//data, err := ioutil.ReadFile(filename)
	data, err := os.ReadFile(filename)
	return string(data), err
}
