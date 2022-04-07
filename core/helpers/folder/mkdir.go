package folder

import "os"

func MkDir(path string) bool {
	/*err := os.Mkdir(path, 0751)
	if err != nil {
		log.Fatal(err)
	}*/

	err := os.MkdirAll(path, 0751)
	if err != nil {
		//log.Fatal(err)
		return false
	}
	return true
}
