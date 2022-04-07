package file

import (
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"strings"
)

func FilterPath(path string) string {
	p := strings.ReplaceAll(path, `\\`, filesystem.DirSeparator())
	p = strings.ReplaceAll(p, "//", filesystem.DirSeparator())
	p = strings.ReplaceAll(p, "/", filesystem.DirSeparator())
	p = strings.ReplaceAll(p, `\`, filesystem.DirSeparator())
	return p
}
