package file

import "path/filepath"

func Name(filePath string) string {
	return filepath.Base(filePath)
}
