package filesystem

import (
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/realpath"
	"path/filepath"
)

// Returns the name from a path
// Example: /test1/test2/test -> test
// Example /test1/test2/test.txt -> test.txt
func Name(path string) string {
	return filepath.Base(path)
}

// It should be an existing!! It doesn't create a path based on non existing files or folders
func RealPath(path string) (string, error) {
	//return filepath.Clean(path)
	return realpath.Realpath(path)
}

func Dir(path string) string {
	return filepath.Dir(path)
}

func DirSeparator() string {
	//return filepath.FromSlash("/")
	return filepath.Separator
}
