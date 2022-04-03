package path

import (
	"github.com/KyaXTeam/go-core/v2/core/config"
	"github.com/KyaXTeam/go-core/v2/core/helpers/filesystem"
	"github.com/KyaXTeam/go-core/v2/core/helpers/hash"
	"github.com/KyaXTeam/go-core/v2/core/helpers/str"
	"log"
	"os"
	"path/filepath"
)

// Root -> It should return the current process dir path
func Root() (dir string) {
	//currentDirPath, err := os.Getwd()
	path, err := os.Executable()
	if err != nil || path == "" {
		log.Println("Error getting Root Path", err)
		return ""
	}
	currentDirPath := filesystem.Dir(path)
	currentDirPath += filesystem.DirSeparator()
	return currentDirPath
}

// GenRealPath -> It generates a real path... it's not realpath
func GenRealPath(path string, appDataPath bool) (string, error) {
	// Check if first character contains a slash /, if not then we should add root dir!, if yes then we should leave it
	// as it is

	p := ""
	var err error = nil
	switch char := str.GetChar(path, 0); char {
	case filesystem.DirSeparator():
		// It's being declared a full path! starting from root /
		p = path
	case ".":
		// Leave it as it is (it's the current dir)
		// But we should also take in consideration that . may be an invisible folder!

		switch char1 := str.GetChar(path, 1); char1 {
		case filesystem.DirSeparator():
			p = path
		case ".":
			//   ..
			tmp, _ := filepath.Abs(Root() + path)
			p = tmp + filesystem.DirSeparator()
		default:
			// it can be a folder or even empty string...
			if char1 != "" {
				// it's a name of an invisible folder
				tmp, _ := filepath.Abs(Root() + path)
				p = tmp + filesystem.DirSeparator()
			} else {
				// nothing else there...
			}
		}
	default:
		// Add the root path!
		p = Root()
		if p != "" {
			if appDataPath {
				// We take the name from the config + we generate additional suffix if other binaries exist in the same
				// folder... in this case, each binary will have its own appdata folder!
				appDataP, appErr := GenRealPath(config.GetConfig().Application.AppDataPath+"_"+hash.MD5(filepath.Base(os.Args[0])), false)
				if appErr == nil {
					p = appDataP + filesystem.DirSeparator()
				}
			}
			p = p + path + filesystem.DirSeparator()
		}
	}

	return p, err
}
