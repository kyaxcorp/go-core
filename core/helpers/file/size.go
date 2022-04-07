package file

import "os"

func Size(filePath string) (int64, error) {
	fi, _err := os.Stat(filePath)
	if _err != nil {
		//log.Fatal(err)
		return 0, _err
	}
	return fi.Size(), nil
}
