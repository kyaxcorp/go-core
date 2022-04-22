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

func MkDirErr(path string, perm ...os.FileMode) error {
	var p os.FileMode
	p = 0751
	if len(perm) > 0 {
		p = perm[0]
	}
	return os.MkdirAll(path, p)
}
